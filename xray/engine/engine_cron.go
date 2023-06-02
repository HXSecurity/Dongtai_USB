package engine

import (
	"bufio"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/HXSecurity/Dongtai_USB/service"
	"github.com/HXSecurity/Dongtai_USB/xray/model"
)

func EngineAdu_max() {

}
func (engine *Engine_Xray) RequestMessages_max(Snapshot []model.Detail, p int) []service.RequestMessages {
	stream := make([]service.RequestMessages, 0)
	for i := 0; i < p; i++ {
		stream = append(stream, service.RequestMessages{Request: Snapshot[i].Request, Response: Snapshot[i].Response})
	}
	return stream
}
func (engine *Engine_Xray) ReadHTTP_max(xray_max string) ([]model.Detail, []model.Connection, error) {
	stream1 := make([]model.Detail, 0)
	stream2 := make([]model.Connection, 0)

	Detail_1 := strings.Split(xray_max, "漏洞探测过程的请求流为")[0:]
	Detail := strings.Split(Detail_1[1], "```")[1:]

	for i := 0; i < len(Detail); i += 4 {
		stream1 = append(stream1, model.Detail{Request: Detail[i], Response: Detail[i+2]})
	}

	for i := 0; i < len(stream1); i++ {
		req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(strings.TrimSpace(stream1[i].Request) + "\r\n\r\n")))
		if err != nil && errors.Is(err, io.EOF) {
			return nil, nil, err
		}
		if err != nil {
			return nil, nil, err
		}
		res, err := http.ReadResponse(bufio.NewReader(strings.NewReader(strings.TrimSpace(stream1[i].Response)+"\r\n\r\n")), req)
		if err != nil && errors.Is(err, io.EOF) {
			return nil, nil, err
		}
		if err != nil {
			return nil, nil, err
		}
		stream2 = append(stream2, model.Connection{Request: req, Response: res})
	}

	return stream1, stream2, nil
}
