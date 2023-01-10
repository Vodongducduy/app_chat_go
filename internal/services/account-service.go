package services

import (
	"appchat/internal/dtos"
	"context"
)

type (
	IAccountService interface {
		InsertAccount(ctx context.Context, dto *dtos.UserRegister) error
	}
)

func (s service) InsertAccount(ctx context.Context, dto *dtos.UserRegister) error {
	if err := s.repository.InsertAccount(ctx, dto); err != nil {
		return err
	}
	return nil
}
