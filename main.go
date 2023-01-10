package main

import (
	"appchat/configs"
	"appchat/internal/router"
	"log"
)

// @title Swagger Golang AppChat
// @version 1.0
// @description This is a document for Golang service
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
// @description Bearer token

func main() {
	config, err := configs.LoadConfig("config.dev")
	if err != nil {
		log.Fatalln("Cannot load config")
	}
	configs.AppConfig = config
	err = router.SetUpApplication(config)
	if err != nil {
		log.Println("SetUpApplication: setting router fail!!!", err.Error())
	}
}
