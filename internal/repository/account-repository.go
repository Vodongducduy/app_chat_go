package repository

import (
	"appchat/internal/dtos"
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

type (
	IAccountRepository interface {
		InsertAccount(ctx context.Context, dtos *dtos.UserRegister) error
	}
)

func (r repository) InsertAccount(ctx context.Context, dtos *dtos.UserRegister) error {

	tx, err := r.Db.Begin(ctx)
	if err != nil {
		logrus.Error("InsertAccount: ", err.Error())
		return err
	}
	defer tx.Rollback(ctx)
	// setting uuid
	dtos.AccountRegister.ID = uuid.New()
	// input insert account user
	today := time.Now()
	_, err = r.Db.Exec(ctx, "insert into account (id,user_name,password,created_by,updated_by,created_at,updated_at) values ($1,$2,$3,$4,$5,$6,$7)",
		dtos.AccountRegister.ID,
		dtos.AccountRegister.UserName,
		dtos.AccountRegister.Password,
		dtos.AccountRegister.ID,
		dtos.AccountRegister.ID,
		today,
		today)
	if err != nil {
		log.Println("Insert account user", err.Error())
		return err
	}
	// user profile
	var InsertUser = "insert into user_profile (user_id,full_name,email,phone_number,created_by,updated_by,created_at,updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8)"
	_, err = r.Db.Exec(ctx, InsertUser,
		dtos.AccountRegister.ID,
		dtos.UserProfileRegister.FullName,
		dtos.UserProfileRegister.Email,
		dtos.UserProfileRegister.PhoneNumber,
		dtos.AccountRegister.ID,
		dtos.AccountRegister.ID,
		today,
		today)
	if err != nil {
		log.Println("Insert user profile", err.Error())
		return err
	}
	return tx.Commit(ctx)
}
