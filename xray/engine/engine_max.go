package engine

import (
	"bufio"
	"io"
	"net/http"
	"strings"

	"github.com/HXSecurity/Dongtai_USB/xray/model"
)

func EngineAdu_max() {

}
func (engine *Engine_Xray) RequestMessages_max(Snapshot []model.Detail, p int) []model.RequestMessages {
	stream := make([]model.RequestMessages, 0)
	for i := 0; i < p; i++ {
		stream = append(stream, model.RequestMessages{Request: Snapshot[i].Request, Response: Snapshot[i].Response})
	}
	return stream
}
func (engine *Engine_Xray) ReadHTTP_max(xray_max string) ([]model.Detail, []model.Connection, error) {
	stream1 := make([]model.Detail, 0)
	stream2 := make([]model.Connection, 0)

	Detail := strings.Split(xray_max, "```")[9:]
	for i := 0; i < len(Detail); i += 4 {
		stream1 = append(stream1, model.Detail{Request: Detail[i], Response: Detail[i+2]})
	}

	for i := 0; i < len(stream1); i++ {
		req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(stream1[i].Response)))
		if err == io.EOF {
			print("error")
		}
		res, err := http.ReadResponse(bufio.NewReader(strings.NewReader(stream1[i].Request)), req)
		if err == io.EOF {
			print("error")
		}
		print(stream1[i].Request)
		print(stream1[i].Response)
		stream2 = append(stream2, model.Connection{Request: req, Response: res})
	}
	return stream1, stream2, nil
}
