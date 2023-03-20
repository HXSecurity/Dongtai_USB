package engine

import (
	"bufio"
	"io"
	"net/http"
	"strings"

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
func (engine *Engine_Xray) EngineAdu(connection []model.Connection, p int) model.Engine {
	var req model.Engine
	for i := 0; i < p; i++ {
		res := connection[i].Response.Header.Get("Dt-Request-Id")
		url := connection[i].Request.URL.String()
		arr := strings.Split(res, ".")
		req.AgentID = append(req.AgentID, arr[0])
		req.DtuuidID = append(req.DtuuidID, arr[1])
		req.Urls = append(req.Urls, url)
	}
	return req
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
