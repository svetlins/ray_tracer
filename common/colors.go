package common

var PureRed Vector
var Red, Blue, Yellow, Green Vector

func init() {
	PureRed = NewColor(1,0,0)

	Red	= NewColor(0xee/255.0, 0x40/255.0, 0x35/255.0)
	Blue	= NewColor(0x03/255.0, 0x92/255.0, 0xcf/255.0)
	Yellow	= NewColor(0xfd/255.0, 0xf4/255.0, 0x98/255.0)
	Green	= NewColor(0x7b/255.0, 0xc0/255.0, 0x43/255.0)
}

func Black() Vector { return NewColor(0, 0, 0) }
func White() Vector { return NewColor(1, 1, 1) }
