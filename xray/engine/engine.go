package engine

import (
	"bufio"
	"io"
	"net/http"
	"strings"

	"github.com/HXSecurity/Dongtai_USB/config"
	"github.com/HXSecurity/Dongtai_USB/xray/model"
)

type Engine_Xray struct {
}

func (engine *Engine_Xray) RequestMessages(Snapshot [][]string, p int) []model.RequestMessages {
	stream := make([]model.RequestMessages, 0)
	for i := 0; i < p; i++ {
		stream = append(stream, model.RequestMessages{Request: Snapshot[i][0], Response: Snapshot[i][1]})
	}
	return stream
}
func (engine *Engine_Xray) EngineXray(agent string, connection []model.Connection, p int) model.Engine {
	var xray model.Engine
	AgentID := make([]string, 0)
	DtuuidID := make([]string, 0)
	Dtmark := make([]string, 0)
	for i := 0; i < p; i++ {
		dtmark := connection[i].Request.Header.Get("dt-mark-header")
		url := connection[i].Request.URL.String()
		if strings.Contains(url, "?") {
			URL_arr := strings.Split(url, "?")
			xray.Urls = append(xray.Urls, URL_arr[0])
		} else {
			xray.Urls = append(xray.Urls, url)
		}
		if agent == "" {
			config.Log.Printf("找不到 Dt-Request-Id 请求头")
			xray.AgentID = AgentID
			xray.DtuuidID = DtuuidID
		} else {
			arr := strings.Split(agent, ".")
			xray.AgentID = append(xray.AgentID, arr[0])
			xray.DtuuidID = append(xray.DtuuidID, arr[1])
		}
		if dtmark == "" {
			config.Log.Printf("找不到 dt-mark-header 请求头")
			xray.Dtmark = Dtmark
		} else {
			xray.Dtmark = append(xray.Dtmark, dtmark)
		}

	}
	return xray
}

func (engine *Engine_Xray) ReadHTTP(Snapshot [][]string, p int) ([]model.Connection, error) {
	stream := make([]model.Connection, 0)
	for i := 0; i < p; i++ {
		req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(Snapshot[i][0])))
		if err == io.EOF {
			print("error")
		}
		res, err := http.ReadResponse(bufio.NewReader(strings.NewReader(Snapshot[i][1])), req)
		if err == io.EOF {
			print("error")
		}
		stream = append(stream, model.Connection{Request: req, Response: res})
	}
	return stream, nil
}

func (engine *Engine_Xray) VulType(m string) string {
	arr := strings.Split(m, "/")
	a := arr[0]
	return a
}
