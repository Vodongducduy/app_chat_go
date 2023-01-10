package controller

import (
	"appchat/configs"
	"appchat/internal/services"
)

var Errors *configs.ErrorDefine

type (
	IController interface {
		IUserController
		IAccountController
	}

	controller struct {
		service services.IService
	}
)

func NewController(service services.IService) IController {
	return &controller{service: service}
}
