package service

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/HXSecurity/Dongtai_USB/config"
)

func (s *USB_Xray) Client(content *bytes.Buffer) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", config.Viper.GetString("usb.url"), content)
	if err != nil {
		log.Printf("NewRequest: %v\n", err)
	}
	req.Header.Set("X-Dongtai-Dast-Vul-Api-Version", "v1")
	req.Header.Set("X-Dongtai-Dast-Vul-Api-Authorization", config.Viper.GetString("usb.token"))
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := client.Do(req)
	if err != nil {
		log.Printf("client: %v\n", err)
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}
