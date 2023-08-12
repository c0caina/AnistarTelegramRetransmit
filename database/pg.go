package database

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type postgreSQL struct {
	conn *pgx.Conn
}

func NewPostgreSQL(dbURL string) (*postgreSQL, error) {
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}
	return &postgreSQL{conn: conn}, nil
}

func (p postgreSQL) ListAnime(telegramId int) ([]string, error) {
	rows, _ := p.conn.Query(context.Background(), `SELECT anime_title FROM "user" WHERE telegram_id=$1`, telegramId)
	var anime_title []string
	for rows.Next() {
		err := rows.Scan(&anime_title)
		if err != nil {
			return nil, err
		}
	}
	return anime_title, rows.Err()
}

func (p postgreSQL) InsertNewUser(telegramId int) error {
	_, err := p.conn.Exec(context.Background(), `INSERT INTO "user"(telegram_id) VALUES($1)`, telegramId)
	return err
}

func (p postgreSQL) AddInAnimeList(telegramId int, anime string) error {
	_, err := p.conn.Exec(context.Background(), `UPDATE "user" SET anime_title = array_append(anime_title, $1) WHERE telegram_id=$2 `, anime, telegramId)
	return err
}

func (p postgreSQL) RemoveFromAnimeList(telegramId int, anime string) error {
	_, err := p.conn.Exec(context.Background(), `UPDATE "user" SET anime_title = array_remove(anime_title, $1) WHERE telegram_id=$2 `, anime, telegramId)
	return err
}
