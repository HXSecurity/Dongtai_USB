package config

import (
	"log"
	"os"
)

func USB_Log() *log.Logger {
	USB_log := log.New(os.Stderr, "Dongtai_USB - ", 3)
	return USB_log
}
