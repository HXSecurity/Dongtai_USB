package config

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// token := context.Request.Header.Get("x-token")
		// if token == "" {
		// 	context.JSON(200, gin.H{
		// 		"msg": "权限不足", "code": 403,
		// 	})
		// 	context.Abort()
		// }
		if InMap(Viper.GetString("usb.ip"), context.ClientIP()) {
			context.JSON(200, gin.H{
				"msg": "权限不足", "code": 403,
			})
			context.Abort()
		}
		context.Next()
	}
}

func InMap(m string, i string) bool {
	arr := strings.Split(m, ",")
	set := make(map[string]struct{})
	for _, value := range arr {
		set[value] = struct{}{}
	}
	_, ok := set[i]
	return !ok
}
