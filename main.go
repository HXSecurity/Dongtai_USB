package main

import (
	"github.com/HXSecurity/Dongtai_USB/config"
	"github.com/HXSecurity/Dongtai_USB/xray/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.Viper = config.Config()

	var USB_Xray = new(service.USB_Xray)
	Usbrouter := router.Group("api").Use(config.JWTAuth())
	Usbrouter.POST("/v1/xray", USB_Xray.Xray)

	router.Run(":5005")
}
