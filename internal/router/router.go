package router

import (
	"appchat/configs"
	"appchat/db"
	"appchat/internal/controller"
	"appchat/internal/repository"
	"appchat/internal/services"
	"context"
	"log"

	"github.com/labstack/echo/v4"
)

type (
	IRouter interface {
		setRouter()
		IUserRouter
	}

	router struct {
		controller    controller.IController
		routerSetting *echo.Echo
	}
)

func NewRouter(routerSetting *echo.Echo, controller controller.IController) IRouter {
	return &router{routerSetting: routerSetting, controller: controller}
}

func (r router) setRouter() {
	r.SetUserRouter()
	r.SetAccountUserRouter()
}

func SetUpApplication(configFile *configs.Config) error {
	conn, err := db.Connect(configFile)
	if err != nil {
		log.Fatal("SetUpApplication: ", err.Error())
	}
	defer conn.Close(context.Background())

	// Setup router echo
	e := echo.New()

	repo := repository.NewRepository(conn)
	revice := services.NewService(repo)
	ctrl := controller.NewController(revice)
	router := NewRouter(e, ctrl)

	// set up router
	router.setRouter()

	e.Logger.Fatal(e.Start(configFile.Server.Port))
	return err
}
