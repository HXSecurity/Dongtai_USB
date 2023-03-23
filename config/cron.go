package config

import (
	"encoding/json"
	"time"

	"github.com/robfig/cron/v3"
)

func (usb *USB_config) Cron(timer func(time.Time, time.Time)) {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/5 * * * * *", func() {
		st, _ := time.ParseDuration("-700000s")
		timer(time.Now().UTC(), time.Now().Add(st).UTC())
	})
	c.Start()
}

func (usb *USB_config) Js(x interface{}) string {
	js, _ := json.Marshal(x)
	return string(js)
}
