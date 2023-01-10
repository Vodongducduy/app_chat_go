package configs

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Common *CommonConfig

func init() {
	Common = LoadCommonConfig()
}

type (
	CommonConfig struct {
		FileStatus status `yaml:"status"`
	}

	status struct {
		Succsess string `yaml:"success"`
		Error    string `yaml:"error"`
	}
)

func LoadCommonConfig() *CommonConfig {
	viper.AddConfigPath("./configs")
	viper.SetConfigFile("common")
	viper.SetConfigFile("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			log.Fatalln("Config file not found", err)
			return nil
		} else {
			log.Fatalln("Config file was found but another error was produced", err)
		}
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("CommonConfig file changed:", e.Name)
	})
	viper.WatchConfig()

	config := &CommonConfig{}
	err := viper.Unmarshal(config)
	if err != nil {
		log.Fatalln("CommonConfig: error unmarshal", err)
	}
	return config
}
