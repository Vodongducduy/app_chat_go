package db

import (
	"appchat/configs"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func Connect(configs *configs.Config) (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), configs.Database.Url)
	if err != nil {
		log.Println("Connect: Can not connect to database", err.Error())
		return nil, err
	}
	return conn, nil
}
