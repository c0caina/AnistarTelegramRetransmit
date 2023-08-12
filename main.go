package main

import (
	"github.com/c0caina/AnistarTelegramRetransmit/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

//5280716552:AAH-0CPol0dptgSPHHI6cViS1LQDN7OU-aQ

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Panic(err)
	}

	bot, err := tgbotapi.NewBotAPI("5280716552:AAH-0CPol0dptgSPHHI6cViS1LQDN7OU-aQ")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		command := cfg.CheckCommand(update.Message.Command())
		if command == "#list" {
			//Todo: список всех т айтлов в расписании.
			msg.Text = command
		} else if command == "#subscribe" {
			//Todo: Подписка
			msg.Text = command
		} else if command == "#unsubscribe" {
			//Todo: Отписка
			msg.Text = command
		} else if command != "" {
			msg.Text = command
		} else {
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
