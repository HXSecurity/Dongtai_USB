package model

import (
	"net/http"
)

type Connection struct {
	Request  *http.Request
	Response *http.Response
}
type Engine struct {
	AgentID  []string
	DtuuidID []string
	Urls     []string
	Dtmark   []string
}
