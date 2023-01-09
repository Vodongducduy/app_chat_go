package services

import (
	"appchat/internal/dtos"
	"context"

	"github.com/google/uuid"
)

type (
	IUserService interface {
		GetUserProfile(ctx context.Context, id uuid.UUID) (*dtos.UserProfile, error)
	}
)

func (s service) GetUserProfile(ctx context.Context, id uuid.UUID) (*dtos.UserProfile, error) {
	userProfile, err := s.repository.GetUsers(ctx, id)
	if err != nil {
		return userProfile, err
	}
	return userProfile, err
}
