package common

type float float64

var simpleMode bool

func init() {
	simpleMode = false
}

func EnableSimpleMode() {
	simpleMode = true
}
