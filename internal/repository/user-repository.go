package repository

import (
	"appchat/internal/dtos"
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type (
	IUserRepository interface {
		GetUsers(ctx context.Context, id uuid.UUID) (*dtos.UserProfile, error)
	}
)

func (r repository) GetUsers(ctx context.Context, id uuid.UUID) (*dtos.UserProfile, error) {
	var output dtos.UserProfile
	err := r.Db.QueryRow(ctx, "select full_name, email, phone_number from user_profile where user_id = %1", id).
		Scan(&output.FullName, &output.Email, &output.PhoneNumber)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		log.Println("GetUsers: ", ErrorNoRecord)
		return nil, ErrorNoRecord
	}
	return &output, err
}
