package services

import "appchat/internal/repository"

type (
	IService interface {
		IUserService
	}

	service struct {
		repository repository.IRepository
	}
)

func NewService(repo repository.IRepository) IService {
	return &service{repository: repo}
}
