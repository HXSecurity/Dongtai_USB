package model

type Request struct {
	Data struct {
		CreateTime int64 `json:"create_time"`
		Detail     struct {
			Addr  string `json:"addr"`
			Extra struct {
				AvgTime string `json:"avg_time"`
				NTime   string `json:"n_time"`
				PTime   string `json:"p_time"`
				Param   struct {
					Key      string `json:"key"`
					Position string `json:"position"`
					Value    string `json:"value"`
				} `json:"param"`
				SleepTime string `json:"sleep_time"`
				Stat      string `json:"stat"`
				StdDev    string `json:"std_dev"`
				Title     string `json:"title"`
				Type      string `json:"type"`
			} `json:"extra"`
			Payload  string     `json:"payload"`
			Snapshot [][]string `json:"snapshot"`
		} `json:"detail"`
		Plugin string `json:"plugin"`
		Target struct {
			Params []struct {
				Path     []string `json:"path"`
				Position string   `json:"position"`
			} `json:"params"`
			URL string `json:"url"`
		} `json:"target"`
	} `json:"data"`
	Type string `json:"type"`
}
