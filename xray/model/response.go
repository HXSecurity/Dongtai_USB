package model

type Response struct {
	VulName         string            `json:"vul_name"`
	Detail          string            `json:"detail"`
	VulLevel        string            `json:"vul_level"`
	Urls            []string          `json:"urls"`
	Payload         string            `json:"payload"`
	CreateTime      int64             `json:"create_time"`
	VulType         string            `json:"vul_type"`
	RequestMessages []RequestMessages `json:"request_messages"`
	Target          string            `json:"target"`
	DtUUIDID        []string          `json:"dt_uuid_id"`
	AgentID         []string          `json:"agent_id"`
	DongtaiVulType  []string          `json:"dongtai_vul_type"`
	DastTag         string            `json:"dast_tag"`
}
type RequestMessages struct {
	Request  string `json:"request"`
	Response string `json:"response"`
}

type Target struct {
	Params []struct {
		Path     []string `json:"path"`
		Position string   `json:"position"`
	} `json:"params"`
	URL string `json:"url"`
}
