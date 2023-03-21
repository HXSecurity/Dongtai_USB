package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/HXSecurity/Dongtai_USB/config"
	"github.com/HXSecurity/Dongtai_USB/xray/engine"
	"github.com/HXSecurity/Dongtai_USB/xray/model"
	"github.com/gin-gonic/gin"
)

type USB_Xray struct {
}

var engine_Xray = new(engine.Engine_Xray)
var buffer bytes.Buffer
var request model.Request
var Json map[string]interface{}

func (s *USB_Xray) Xray(context *gin.Context) {
	err := context.ShouldBindJSON(&request)
	if err != nil {
		config.Log.Print(err.Error(), context)
		return
	}

	res, err := engine_Xray.ReadHTTP(request.Data.Detail.Snapshot, len(request.Data.Detail.Snapshot))
	if err != nil {
		config.Log.Print(err)
		context.Data(200, "application/json; charset=utf-8", []byte("ok"))
		return
	}
	if request.Type != "web_vuln" {
		config.Log.Printf("上报数据类型不是 web_vuln")
		context.Data(200, "application/json; charset=utf-8", []byte("ok"))
		return
	}
	red := res[0].Response.Header.Get("Dt-Request-Id")
	config.Log.Println(request.Data.Detail.Snapshot)
	if red == "" {
		config.Log.Printf("找不到 Dt-Request-Id 请求头")
		context.Data(200, "application/json; charset=utf-8", []byte("ok"))
		return
	}

	Response := &model.Response{
		VulName:         request.Data.Target.URL + " " + engine_Xray.VulType(request.Data.Plugin),
		Detail:          "在" + request.Data.Target.URL + "发现了" + engine_Xray.VulType(request.Data.Plugin),
		VulLevel:        "HIGH",
		Urls:            engine_Xray.EngineAdu(res, len(request.Data.Detail.Snapshot)).Urls,
		Payload:         request.Data.Detail.Payload,
		CreateTime:      time.Now().Unix(),
		VulType:         engine_Xray.VulType(request.Data.Plugin),
		RequestMessages: engine_Xray.RequestMessages(request.Data.Detail.Snapshot, len(request.Data.Detail.Snapshot)),
		Target:          fmt.Sprintf("%s", request.Data.Target),
		DtUUIDID:        engine_Xray.EngineAdu(res, len(request.Data.Detail.Snapshot)).DtuuidID,
		AgentID:         engine_Xray.EngineAdu(res, len(request.Data.Detail.Snapshot)).AgentID,
		DongtaiVulType:  engine_Xray.EngineAdu(res, len(request.Data.Detail.Snapshot)).Urls,
		DastTag:         "Xray",
	}

	config.Log.Println(Response)
	if err := json.NewEncoder(&buffer).Encode(Response); err != nil {
		config.Log.Printf("json 格式错误")
		context.Data(200, "application/json; charset=utf-8", []byte("ok"))
		return
	}
	body, err := s.Client(&buffer)
	if err != nil {
		config.Log.Println(err)
		return
	}
	context.Data(200, "application/json; charset=utf-8", []byte("ok"))
	json.Unmarshal(body, &Json)
	config.Log.Println(Json)
	// context.JSON(200, gin.H{
	// 	"msg": "success", "code": 200, "body": Json,
	// })
}
