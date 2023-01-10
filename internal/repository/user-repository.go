package repository

import (
	"appchat/internal/dtos"
	"appchat/internal/models"
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
	var output models.UserProfile
	err := r.Db.QueryRow(ctx, "select full_name, email, phone_number from user_profile where user_id = $1", id).
		Scan(&output.FullName, &output.Email, &output.PhoneNumber)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("GetUsers: repository", ErrorNoRecord)
			return nil, ErrorNoRecord
		}
		log.Println("GetUsers: internal", err)
		return nil, err
	}

	UserProfile := &dtos.UserProfile{
		FullName:    output.FullName,
		Email:       output.Email,
		PhoneNumber: output.PhoneNumber,
	}

	return UserProfile, nil
}
