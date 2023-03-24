package service

import (
	"time"

	"github.com/HXSecurity/Dongtai_USB/config"
	"github.com/HXSecurity/Dongtai_USB/xray/engine"
	"github.com/HXSecurity/Dongtai_USB/xray/model"
	"github.com/gin-gonic/gin"
)

type USB_Xray struct {
}

var engine_Xray = new(engine.Engine_Xray)

func (s *USB_Xray) Xray(context *gin.Context) {

	var request model.Request
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

	agent := res[0].Response.Header.Get("Dt-Request-Id")
	config.Log.Println(request.Data.Detail.Snapshot)
	if agent == "" {
		config.Log.Printf("找不到 Dt-Request-Id 请求头")
		context.Data(200, "application/json; charset=utf-8", []byte("ok"))
		return
	}
	config.Log.Println("我的漏洞格式是 ============" + engine_Xray.VulType(request.Data.Plugin) + "=========================")
	Response := &model.Response{
		VulName:         request.Data.Target.URL + " " + engine_Xray.VulType(request.Data.Plugin),
		Detail:          "在" + request.Data.Target.URL + "发现了" + engine_Xray.VulType(request.Data.Plugin),
		VulLevel:        "HIGH",
		Urls:            engine_Xray.EngineXray(res[0].Response.Header.Get("Dt-Request-Id"), res, len(request.Data.Detail.Snapshot)).Urls,
		Payload:         request.Data.Detail.Payload,
		CreateTime:      time.Now().Unix(),
		VulType:         engine_Xray.VulType(request.Data.Plugin),
		RequestMessages: engine_Xray.RequestMessages(request.Data.Detail.Snapshot, len(request.Data.Detail.Snapshot)),
		Target:          request.Data.Target.URL,
		DtUUIDID:        engine_Xray.EngineXray(res[0].Response.Header.Get("Dt-Request-Id"), res, len(request.Data.Detail.Snapshot)).DtuuidID,
		AgentID:         engine_Xray.EngineXray(res[0].Response.Header.Get("Dt-Request-Id"), res, len(request.Data.Detail.Snapshot)).AgentID,
		DongtaiVulType:  []string{(model.Vultype()[engine_Xray.VulType(request.Data.Plugin)])},
		DastTag:         "Xray",
		Dtmark:          engine_Xray.EngineXray(res[0].Response.Header.Get("Dt-Request-Id"), res, len(request.Data.Detail.Snapshot)).Dtmark,
	}
	context.Data(200, "application/json; charset=utf-8", []byte(s.Client(Response, context)))
}
