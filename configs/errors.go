package configs

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Errors *ErrorConfig

func init() {
	Errors = LoadErrorsConfig()
}

type (
	ErrorDefine struct {
		ErrInternal struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrInternal"`
		ErrBadRequest struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrBadRequest"`
		ErrUnauthorized struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrUnauthorized"`
		ErrForbidden struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrForbidden"`
		ErrNotFound struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrNotFound"`
		ErrConflict struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrConflict"`

		//field validation
		ErrFieldRequired struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrFieldRequired"`
		ErrFieldInvalid struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrFieldInvalid"`
		ErrFieldFormat struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrFieldFormat"`
		ErrFieldLength struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrFieldLength"`
		ErrParamInvalid struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrParamInvalid"`
		ErrInvalidToken struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrInvalidToken"`
		ErrBankAccountExist struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrBankAccountExist"`
		ErrMissingValue struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrMissingValue"`
		ErrInvalidType struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrInvalidType"`
		ErrInvalidBusinessCode struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrInvalidBusinessCode"`
		ErrExistLicenseNumber struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrExistLicenseNumber"`
		ErrInvalidAddress struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrInvalidAddress"`
		ErrNationalIdUserExist struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrNationalIdUserExist"`
		ErrPhoneNumberExist struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrPhoneNumberExist"`
		ErrInComeTaxCodeExist struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrInComeTaxCodeExist"`
		ErrAddressInvalid struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrAddressInvalid"`
		ErrInvalidFormatTime struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrInvalidFormatTime"`
		ErrBankBranch struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrBankBranch"`
		ErrBankCode struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrBankCode"`
		ErrUnsupportedFile struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrUnsupportedFile"`
		ErrInsufficientBalance struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrInsufficientBalance"`
		ErrInvalidDocumentType struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrInvalidDocumentType"`
		ErrIsActivated struct {
			Code    string `yaml:"code"`
			Message string `yaml:"message"`
		} `yaml:"ErrIsActivated"`
	}

	ErrorConfig struct {
		ReturnError ErrorDefine `yaml:"returnError"`
	}
)

func LoadErrorsConfig() *ErrorConfig {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("errors")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Fatal("ErrorsConfig file not found", err)
		} else {
			logrus.Fatal("ErrorsConfig file was found but another error was produced", err)
		}
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("ErrorsConfigs file changed:", e.Name)
	})
	viper.WatchConfig()

	config := &ErrorConfig{}
	err := viper.Unmarshal(config)
	if err != nil {
		logrus.Fatal("ErrorsConfig: error unmarshal", err)
	}
	return config
}
