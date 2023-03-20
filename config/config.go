package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Config(path ...string) *viper.Viper {
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
	return config
}
