package request

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/HXSecurity/Dongtai_USB/config"
	"github.com/HXSecurity/Dongtai_USB/service"
	"github.com/HXSecurity/Dongtai_USB/xray/model"
)

func (s *USB_Xray) Xray_cron(before time.Time, after time.Time) {
	if config.Viper.GetString("usb.xray_url") == "" {
		return
	}
	config.Log.Printf("正在自动从 xray 拉取数据 !!!")
	var buffer bytes.Buffer
	var Request_max2 model.Request_max2

	Response_max := &model.Request_max1{
		Limit:  10,
		Offset: 0,
		CreatedTime: model.CreatedTime{
			Before: before,
			After:  after,
		},
	}
	config.Log.Println(Response_max)
	if err := json.NewEncoder(&buffer).Encode(Response_max); err != nil {
		config.Log.Printf("Response_max: json 格式错误")
	}
	req, err := http.NewRequest("POST", config.Viper.GetString("usb.xray_url"), &buffer)
	if err != nil {
		config.Log.Printf("NewRequest: %v\n", err)
	}

	req.Header.Set("token", config.Viper.GetString("usb.xray_token"))
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		config.Log.Printf("client: %v\n", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		config.Log.Println(err)
		return
	}
	json.Unmarshal(body, &Request_max2)

	for i := 0; i < len(Request_max2.Data.Content); i++ {
		xray_max := Request_max2.Data.Content[i]

		Detail, Connection, err := engine_Xray.ReadHTTP_max(xray_max.Detail)
		if err != nil {
			config.Log.Print(err)
			return
		}
		config.Log.Println(xray_max.Detail)
		agent := Connection[0].Response.Header.Get("Dt-Request-Id")
		if agent == "" {
			config.Log.Printf("找不到 Dt-Request-Id 请求头")
			return
		}

		Response := &service.Response{
			VulName:         xray_max.Target.URL + " " + xray_max.Category,
			Detail:          "在" + xray_max.Target.URL + "发现了" + xray_max.Title,
			VulLevel:        (model.VulLevel()[engine_Xray.VulType(xray_max.Category)]),
			Urls:            engine_Xray.EngineXray(Connection[0].Response.Header.Get("Dt-Request-Id"), Connection, len(Connection)).Urls,
			Payload:         fmt.Sprintf("%s", xray_max.Target.Params...),
			CreateTime:      time.Now().Unix(),
			VulType:         engine_Xray.VulType(xray_max.Category),
			RequestMessages: engine_Xray.RequestMessages_max(Detail, len(Detail)),
			Target:          xray_max.Target.URL,
			DtUUIDID:        engine_Xray.EngineXray(Connection[0].Response.Header.Get("Dt-Request-Id"), Connection, len(Connection)).DtuuidID,
			AgentID:         engine_Xray.EngineXray(Connection[0].Response.Header.Get("Dt-Request-Id"), Connection, len(Connection)).AgentID,
			DongtaiVulType:  []string{(model.Vultype()[engine_Xray.VulType(xray_max.Category)])},
			DastTag:         "Xray",
			Dtmark:          engine_Xray.EngineXray(Connection[0].Response.Header.Get("Dt-Request-Id"), Connection, len(Connection)).Dtmark,
		}
		resResponse, err := json.Marshal(Response)
		if err != nil {
			config.Log.Printf("无法解析json")
		} else {
			config.Log.Print(string(resResponse))
		}
		config.Log.Print(service.Client(Response))
	}
}
