package config

import (
	"github.com/gin-gonic/gin"
)

func (usb *USB_config) Init() *gin.Engine {
	Viper = usb.Config()
	Log = usb.USB_Log()
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}
