package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func (usb *USB_config) Config(path ...string) *viper.Viper {
	config := viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("config-tutorial")
	config.SetConfigType("ini")

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}
	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		log.Print("Config file updated.")
	})
	return config
}
