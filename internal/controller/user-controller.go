package controller

import (
	"appchat/internal/repository"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	IUserController interface {
		GetUserProfile(e echo.Context) error
	}
)

func (c controller) GetUserProfile(e echo.Context) error {
	id, err := uuid.Parse("3600f8d0-3e2d-65b2-4103-ae2fcacf2156")
	if err != nil {
		log.Println("Parse uuid: fail")
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	output, err := c.service.GetUserProfile(e.Request().Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrorNoRecord) {
			return e.JSON(http.StatusNotFound, err.Error())
		}
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, output)
}
