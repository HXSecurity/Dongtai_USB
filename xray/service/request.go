package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

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
		log.Print(err.Error(), context)
		return
	}

	res, err := engine_Xray.ReadHTTP(request.Data.Detail.Snapshot, len(request.Data.Detail.Snapshot))
	if err != nil {
		log.Print(err)
		context.Data(200, "application/json; charset=utf-8", []byte("ok"))
		return
	}
	if request.Type != "web_vuln" {
		log.Printf("上报数据类型不是 web_vuln")
		context.Data(200, "application/json; charset=utf-8", []byte("ok"))
		return
	}
	red := res[0].Response.Header.Get("Dt-Request-Id")
	log.Println(request.Data.Detail.Snapshot)
	if red == "" {
		log.Printf("找不到 Dt-Request-Id 请求头")
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

	log.Println(Response)
	if err := json.NewEncoder(&buffer).Encode(Response); err != nil {
		log.Printf("json 格式错误")
		context.Data(200, "application/json; charset=utf-8", []byte("ok"))
		return
	}
	body, err := s.Client(&buffer)
	if err != nil {
		log.Println(err)
		return
	}
	context.Data(200, "application/json; charset=utf-8", []byte("ok"))
	json.Unmarshal(body, &Json)
	log.Println(Json)
	// context.JSON(200, gin.H{
	// 	"msg": "success", "code": 200, "body": Json,
	// })
}
