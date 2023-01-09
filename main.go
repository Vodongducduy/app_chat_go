package main

import (
	"appchat/configs"
	"appchat/internal/router"
	"log"
)

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
