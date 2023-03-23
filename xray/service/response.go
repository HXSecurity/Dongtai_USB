package service

import (
	"bytes"
	"encoding/json"
	"io"

	"net/http"

	"github.com/HXSecurity/Dongtai_USB/config"
	"github.com/HXSecurity/Dongtai_USB/xray/model"
)

func (s *USB_Xray) Client(content *model.Response, c ...interface{}) string {
	var Json map[string]interface{}
	var buffer bytes.Buffer
	if err := json.NewEncoder(&buffer).Encode(content); err != nil {
		config.Log.Printf("json 格式错误")
		return "ok"
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", config.Viper.GetString("usb.iast_url"), &buffer)
	if err != nil {
		config.Log.Printf("NewRequest: %v\n", err)
	}
	req.Header.Set("X-Dongtai-Dast-Vul-Api-Version", "v1")
	req.Header.Set("X-Dongtai-Dast-Vul-Api-Authorization", config.Viper.GetString("usb.dast_token"))
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		config.Log.Printf("client: %v\n", err)
		return "ok"
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		config.Log.Println(err)
		return "ok"
	}
	json.Unmarshal(body, &Json)
	config.Log.Println(Json)
	return "ok"
}
