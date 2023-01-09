package controller

import "appchat/internal/services"

type (
	IController interface {
		IUserController
	}

	controller struct {
		service services.IService
	}
)

func NewController(service services.IService) IController {
	return &controller{service: service}
}
