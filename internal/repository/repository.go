package repository

import (
	"github.com/jackc/pgx/v5"
)

type (
	IRepository interface {
		IUserRepository
		IAccountRepository
	}
	repository struct {
		Db *pgx.Conn
	}
)

func NewRepository(db *pgx.Conn) IRepository {
	return &repository{Db: db}
}
