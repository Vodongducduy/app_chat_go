package controller

import (
	"appchat/internal/controller/responses"
	"appchat/internal/dtos"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	IAccountController interface {
		CreateAccountUser(e echo.Context) error
	}
)

func (c controller) CreateAccountUser(e echo.Context) error {
	var input dtos.UserRegister
	err := e.Bind(&input)
	if err != nil {
		respErr := responses.ErrorResponse{
			Message: Errors.ErrInternal.Message,
			Code:    Errors.ErrInternal.Code,
		}
		return e.JSON(http.StatusBadRequest, respErr)
	}
	// create user
	err = c.service.InsertAccount(e.Request().Context(), &input)
	if err != nil {
		respErr := responses.ErrorResponse{
			Message: Errors.ErrInternal.Message,
			Code:    Errors.ErrInternal.Code,
		}
		return e.JSON(http.StatusInternalServerError, respErr)
	}

	return e.JSON(http.StatusCreated, responses.JsonResponse{})
}
