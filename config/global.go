package config

import (
	"log"

	"github.com/spf13/viper"
)

type USB_config struct {
}

var (
	Viper *viper.Viper
	Log   *log.Logger
)
