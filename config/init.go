package config

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	Viper = Config()
	Log = USB_Log()
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}
