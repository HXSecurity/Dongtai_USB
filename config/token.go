package config

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

func (usb *USB_config) JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// token := context.Request.Header.Get("x-token")
		// if token == "" {
		// 	context.JSON(200, gin.H{
		// 		"msg": "权限不足", "code": 403,
		// 	})
		// 	context.Abort()
		// }

		if usb.IpContains(Viper.GetString("usb.ip"), context.ClientIP()) {
			Log.Printf("权限不足 403:  " + context.ClientIP())
			context.JSON(200, gin.H{
				"msg": "权限不足", "code": 403,
			})
			context.Abort()
		}
		context.Next()
	}
}

func (usb *USB_config) InMap(cidr string, ip string) bool {
	arr := strings.Split(cidr, ",")
	set := make(map[string]struct{})
	for _, value := range arr {
		set[value] = struct{}{}
	}
	_, ok := set[ip]
	return !ok
}

func (usb *USB_config) IpContains(cidr string, ip string) bool {
	arr := strings.Split(cidr, ",")
	set := make(map[bool]struct{})
	for i := 0; i < len(arr); i++ {
		if find := strings.Contains(arr[i], "/"); !find {
			if arr[i] == ip {
				set[false] = struct{}{}
			}
		}
		if find := strings.Contains(arr[i], "/"); find {
			_, ipnet, err := net.ParseCIDR(arr[i])
			if err != nil {
				Log.Printf("cidr写入格式不对: " + arr[i])
				return true
			}
			ipAddr := net.ParseIP(ip)
			set[!ipnet.Contains(ipAddr)] = struct{}{}
		}
	}
	_, ok := set[false]
	return !ok
}
