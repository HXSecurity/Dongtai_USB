package model

import "time"

type Request_max1 struct {
	Limit       int         `json:"limit"`
	Offset      int         `json:"offset"`
	CreatedTime CreatedTime `json:"created_time"`
}
type CreatedTime struct {
	Before time.Time `json:"before"`
	After  time.Time `json:"after"`
}

type Request_max2 struct {
	Err  interface{} `json:"err"`
	Msg  string      `json:"msg"`
	Data struct {
		Total   int `json:"total"`
		Content []struct {
			ID           int       `json:"id"`
			CreatedTime  time.Time `json:"created_time"`
			UpdatedTime  time.Time `json:"updated_time"`
			Status       string    `json:"status"`
			RelatedTasks []int     `json:"related_tasks"`
			RelatedAsset struct {
				Type string `json:"type"`
				ID   int    `json:"id"`
			} `json:"related_asset"`
			Definiteness string      `json:"definiteness"`
			Title        string      `json:"title"`
			Severity     string      `json:"severity"`
			FixedTime    interface{} `json:"fixed_time"`
			LastScanAt   time.Time   `json:"last_scan_at"`
			Target       struct {
				URL    string        `json:"url"`
				Params []interface{} `json:"params"`
			} `json:"target"`
			Cwe               interface{}   `json:"cwe"`
			Poc               string        `json:"poc"`
			Exp               string        `json:"exp"`
			Summary           string        `json:"summary"`
			Impact            string        `json:"impact"`
			Detail            string        `json:"detail"`
			Solution          string        `json:"solution"`
			Cvss              string        `json:"cvss"`
			Category          string        `json:"category"`
			Exposures         []interface{} `json:"exposures"`
			PublishedDatetime time.Time     `json:"published_datetime"`
			XprocessIds       []int         `json:"xprocess_ids"`
			HostID            interface{}   `json:"host_id"`
		} `json:"content"`
	} `json:"data"`
}
