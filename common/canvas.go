package common

import (
	"fmt"
	"math"
)

type Canvas struct {
	raster []Vector
	width  int
	height int
}

func NewCanvas(width, height int) *Canvas {
	raster := make([]Vector, width*height)

	return &Canvas{raster, width, height}
}

func (canvas *Canvas) GetPixel(x, y int) Vector {
	return canvas.raster[y*canvas.width+x]
}

func (canvas *Canvas) SetPixel(x, y int, color Vector) {
	canvas.raster[y*canvas.width+x] = color
}

func (canvas *Canvas) PPM() string {
	currentLine := ""
	output := fmt.Sprintf("P3\n%d %d\n255\n", canvas.width, canvas.height)

	for y := 0; y < canvas.height; y++ {
		for x := 0; x < canvas.width; x++ {
			c := canvas.GetPixel(x, y)
			currentPixel := fmt.Sprintf(
				"%d %d %d ",
				int(math.Min(c.x, 1.0)*255),
				int(math.Min(c.y, 1.0)*255),
				int(math.Min(c.z, 1.0)*255),
			)

			if len(currentLine)+len(currentPixel) > 70 {
				output += currentLine
				output += "\n"
				currentLine = ""
			}

			currentLine += currentPixel
		}
	}

	output += currentLine

	return output
}
