package config

import (
	"encoding/json"
	"time"

	"github.com/robfig/cron/v3"
)

func (usb *USB_config) Cron(name string, timer func(time.Time, time.Time)) {
	if Viper.GetString("usb.type") != name {
		return
	}
	Log.Printf("开始自动拉取黑盒扫描器数据: (" + name + ")")
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/5 * * * * *", func() {
		// st, _ := time.ParseDuration("-700000s")
		st, _ := time.ParseDuration("-10s")
		timer(time.Now().UTC(), time.Now().Add(st).UTC())
	})
	c.Start()
}

func (usb *USB_config) Js(x interface{}) string {
	js, _ := json.Marshal(x)
	return string(js)
}
