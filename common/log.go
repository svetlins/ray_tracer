package common

import (
	"log"
)

var logEnabled bool

func init() {
	logEnabled = true
	log.SetPrefix("LOG: ")
	log.SetFlags(0)
}

func logger(args ...interface{}) {
	if !logEnabled {
		return
	}

	switch args[0].(type) {
	case string:
		log.Printf(args[0].(string), args[1:]...)
	default:
		log.Printf("%+v", args...)
	}
}

func EnableLogging() {
	logEnabled = true
}
