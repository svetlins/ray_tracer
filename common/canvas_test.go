package common

import (
	"strings"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	canvas := NewCanvas(320, 200)

	if canvas.width != 320 || canvas.height != 200 {
		t.Errorf("NewCanvas not working")
	}

	if canvas.raster == nil {
		t.Errorf("NewCanvas not working")
	}

	if canvas.GetPixel(20, 30) != NewColor(0, 0, 0) {
		t.Errorf("NewCanvas not working")
	}
}

func TestCanvasSetPixel(t *testing.T) {
	canvas := NewCanvas(320, 200)

	canvas.SetPixel(20, 30, NewColor(0.5, 0.5, 0.5))

	if canvas.GetPixel(20, 30) != NewColor(0.5, 0.5, 0.5) {
		t.Errorf("NewCanvas not working")
	}
}

func TestCanvasPPM(t *testing.T) {
	canvas := NewCanvas(5, 3)

	if !strings.HasPrefix(canvas.PPM(), "P3\n5 3\n255") {
		t.Errorf("Canvas PPM not working")
	}

	canvas.SetPixel(0, 0, NewColor(1, 0, 0))
	canvas.SetPixel(2, 1, NewColor(0, 0.5, 0))
	canvas.SetPixel(4, 2, NewColor(0, 0, 1))
	suffix := "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 255 "

	if !strings.HasSuffix(canvas.PPM(), suffix) {
		t.Errorf("Canvas PPM not working")
	}
}
