package main

import (
	"github.com/HXSecurity/Dongtai_USB/config"
	"github.com/HXSecurity/Dongtai_USB/xray/service"
)

var USB_Xray = new(service.USB_Xray)

func main() {
	USB := config.Init()
	router := USB.Group("api").Use(config.JWTAuth())
	router.POST("/v1/xray", USB_Xray.Xray)

	config.Log.Printf("The USB runs on port 5005")
	USB.Run(":5005")
}
