package main

import (
	"github.com/HXSecurity/Dongtai_USB/config"
	"github.com/HXSecurity/Dongtai_USB/xray/request"
)

var usb = new(config.USB_config)
var USB_Xray = new(request.USB_Xray)

func main() {
	USB := usb.Init()
	router := USB.Group("api").Use(usb.JWTAuth())

	//推流模式(webhook)：
	router.POST("/v1/xray", USB_Xray.Xray)
	//拉流模式(cron):
	usb.Cron("xray", USB_Xray.Xray_cron)

	config.Log.Printf("The USB runs on port 5005")
	USB.Run(":5005")
}
