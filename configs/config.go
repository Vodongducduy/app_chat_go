package configs

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var AppConfig *Config

type (
	Config struct {
		Database database `yaml:"database"`
		Server   server   `yaml:"server"`
	}

	database struct {
		Url string `yaml:"url"`
	}

	server struct {
		Port string `yaml:"port"`
	}
)

func LoadConfig(configName string) (*Config, error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			log.Println("Config file not found", err)
			return nil, err
		} else {
			log.Println("Config file was found but another error was produced", err)
		}
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()

	config := &Config{}
	err := viper.Unmarshal(config)
	return config, err
}
