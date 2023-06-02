package main

import (
	"github.com/HXSecurity/Dongtai_USB/config"
	"github.com/HXSecurity/Dongtai_USB/xray/request"
)

var usb = new(config.USB_config)
var USB_Xray = new(request.USB_Xray)

func main() {
	USB := usb.Init(USB_Xray.Xray_cron)
	router := USB.Group("api").Use(usb.JWTAuth())
	router.POST("/v1/xray", USB_Xray.Xray)

	config.Log.Printf("The USB runs on port 5005")
	USB.Run(":5005")
}
