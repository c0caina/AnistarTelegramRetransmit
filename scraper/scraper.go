package scraper

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"regexp"
	"strconv"
	"strings"
)

type animevost struct {
	colly   *colly.Collector
	rootUrl string
}

func NewAnimevost(url string) *animevost {
	return &animevost{colly: colly.NewCollector(), rootUrl: url}
}

func (a animevost) ПоследниеОбновления() []string {
	var animeList []string
	a.colly.OnHTML(`ul[class="raspis raspis_fixed"] > li > a[href]`, func(e *colly.HTMLElement) {

		animeList = append(animeList, e.Attr("href"))

	})
	a.colly.Visit(a.rootUrl)

	return animeList
}

func (a animevost) Расписание() []string {
	var animeList []string
	a.colly.OnHTML(`div[class="interDubBgTwo"] > div > a[href]`, func(e *colly.HTMLElement) {

		animeList = append(animeList, a.rootUrl+e.Attr("href"))

	})
	a.colly.Visit(a.rootUrl)

	return animeList
}

func (a animevost) НомераСерий(serialUrl string) []int {
	var serialNumbers []int
	a.colly.OnHTML(`div[class="shortstoryContent"]`, func(e *colly.HTMLElement) {
		s := e.DOM.Find("script").Text()
		r := regexp.MustCompile(`\d+ серия`)
		for {
			res := r.FindString(s)
			if res == "" {
				return
			}
			serialNumber, _ := strconv.Atoi(strings.ReplaceAll(res, ` серия`, ``))
			serialNumbers = append(serialNumbers, serialNumber)
			s = strings.ReplaceAll(s, `"`+res, ``)
		}
	})
	a.colly.Visit(serialUrl)
	return serialNumbers
}

//Может вернуть пустой плеер
func (a animevost) ПлеерСерии(serialUrl string, serialNumber int) string {
	var playerUrl string
	a.colly.OnHTML(`div[class="shortstoryContent"]`, func(e *colly.HTMLElement) {
		s := e.DOM.Find("script").Text()

		sn := fmt.Sprintf(`%v серия":"`, serialNumber)
		r := regexp.MustCompile(sn + `\d+`)
		res := strings.ReplaceAll(r.FindString(s), sn, "")

		playerUrl = "https://v2.vost.pw/frame2.php?play=" + res
	})
	a.colly.Visit(serialUrl)
	return playerUrl
}

func (a animevost) СсылкаНаСкачиваниеСерии(playerUrl string) string {
	var downloadUrl string
	a.colly.OnHTML(`a[href]`, func(e *colly.HTMLElement) {
		if e.Text == "720p (HD)" {
			downloadUrl = e.Attr("href")
		}
	})
	a.colly.Visit(playerUrl)
	return downloadUrl
}
