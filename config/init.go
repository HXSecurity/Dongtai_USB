package config

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (usb *USB_config) Init(f func(time.Time, time.Time)) *gin.Engine {
	Viper = usb.Config()
	Log = usb.USB_Log()
	usb.Cron(f)
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}
