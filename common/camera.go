package common

import (
	"image/color"
	"image"
	"github.com/svetlins/srt/ui"
	"math"
	"runtime"
	_ "time"
	"sync"
)

var ReflectionLimit int
var numChunks int
var numWorkers int

func init() {
	ReflectionLimit = 5
	numChunks       = 20
	numWorkers      = runtime.NumCPU() / 1
}

type Camera struct {
	hSize, vSize          int
	fov                   float64
	transform             Matrix
	inverseTransform             Matrix
	halfWidth, halfHeight float64
	world                 *World
	currentPosition       Vector
}

func NewCamera(hSize, vSize int, fov float64) *Camera {
	var halfWidth, halfHeight float64

	halfView := math.Tan(fov / 2)
	aspect := float64(hSize) / float64(vSize)

	if aspect > 1 {
		halfWidth = halfView
		halfHeight = halfWidth / aspect
	} else {
		halfHeight = halfView
		halfWidth = halfHeight * aspect
	}

	return &Camera{
		hSize:      hSize,
		vSize:      vSize,
		fov:        fov,
		transform:  NewIdentity4(),
		halfWidth:  halfWidth,
		halfHeight: halfHeight,
		currentPosition: NewPoint(0, 1.5, -5),
	}
}

func (c *Camera) SetTransform(t Matrix) {
	c.transform = t
	c.inverseTransform = t.Inverse()
}

func (c *Camera) pixelSize() float64 {
	return (c.halfWidth * 2) / float64(c.hSize)
}

func (c *Camera) RayFor(px, py int) Ray {
	pixelSize := c.pixelSize()

	xOffset := (float64(px) + 0.5) * pixelSize
	yOffset := (float64(py) + 0.5) * pixelSize

	canvasX := c.halfWidth - xOffset
	canvasY := c.halfHeight - yOffset

	worldPixelCenter := c.inverseTransform.Transform(
		NewPoint(canvasX, canvasY, -1),
	)

	origin := c.inverseTransform.Transform(
		NewPoint(0, 0, 0),
	)

	direction := worldPixelCenter.Subtract(origin).Normalize()

	ray := NewRay(origin, direction)

	return ray
}

func (c *Camera) Render(w *World) *Canvas {
	canvas := NewCanvas(c.hSize, c.vSize)

	var wg sync.WaitGroup
	cores := 100
 	xc := float64(c.hSize) / math.Sqrt(float64(cores))
 	yc := float64(c.vSize) / math.Sqrt(float64(cores))

	for xi := 0; xi < int(math.Sqrt(float64(cores))); xi++ {
		for yi := 0; yi < int(math.Sqrt(float64(cores))); yi++ {
			wg.Add(1)
			go func (xi, yi int) {
				for y := 0; y < int(yc); y++ {
					for x := 0; x < int(xc); x++ {
						xx := x + xi * int(xc)
						yy := y + yi * int(yc)

						canvas.SetPixel(xx, yy,
							w.ColorAt(
								c.RayFor(xx, yy),
								ReflectionLimit,
							),
						)
					}
				}
				wg.Done()
			}(xi, yi)
		}
	}

	wg.Wait()

	// for ci := 0; ci < cores; ci++ {
	// 	for y := 0; y < int(yc); y++ {
	// 		for x := 0; x < int(xc); x++ {
	// 			xx := int(math.Min(float64(x + ci * int(xc)), float64(c.hSize - 1)))
	// 			yy := int(math.Min(float64(y + ci * int(yc)), float64(c.vSize - 1)))

	// 			canvas.SetPixel(xx, yy,
	// 				w.ColorAt(
	// 					c.RayFor(xx, yy),
	// 					ReflectionLimit,
	// 				),
	// 			)
	// 		}
	// 	}

	// }
	// for y := 0; y < c.vSize; y++ {
	// 	for x := 0; x < c.hSize; x++ {
	// 		canvas.SetPixel(x, y,
	// 			w.ColorAt(
	// 				c.RayFor(x, y),
	// 				ReflectionLimit,
	// 			),
	// 		)
	// 	}
	// }

	return canvas
}

func (c *Camera) DebugPixel(w *World, x, y int) {
	EnableLogging()

	w.ColorAt(
		c.RayFor(x, y),
		ReflectionLimit,
	)
}

func (c *Camera)SetWorld(w *World) {
	c.world = w
}

func (c *Camera)RenderToSequential(im *image.RGBA) {
	for y := 0; y < c.vSize; y++ {
		for x := 0; x < c.hSize; x++ {

			rayColor := c.world.ColorAt(
				c.RayFor(x, y),
				ReflectionLimit,
			)

			c := color.RGBA{
				uint8(math.Min(rayColor.x, 1) * 255),
				uint8(math.Min(rayColor.y, 1) * 255),
				uint8(math.Min(rayColor.z, 1) * 255),
				1,
			}
			im.SetRGBA(x, y, c)
		}
	}
}

func (c *Camera)RenderTo(im *image.RGBA) {
	cores := 100
 	xc := math.Ceil(float64(c.hSize) / math.Sqrt(float64(cores)))
 	yc := math.Ceil(float64(c.vSize) / math.Sqrt(float64(cores)))

	var wg sync.WaitGroup
	for xi := 0; xi < int(math.Ceil(math.Sqrt(float64(cores)))); xi++ {
		for yi := 0; yi < int(math.Ceil(math.Sqrt(float64(cores)))); yi++ {
			wg.Add(1)
			go func (xi, yi int) {
				for y := 0; y < int(yc); y++ {
					for x := 0; x < int(xc); x++ {
						xx := int(math.Min(float64(x + xi * int(xc)), float64(c.hSize)))
						yy := int(math.Min(float64(y + yi * int(yc)), float64(c.vSize)))

						rayColor := c.world.ColorAt(
							c.RayFor(xx, yy),
							ReflectionLimit,
						)

						c := color.RGBA{
							uint8(math.Min(rayColor.x, 1) * 255),
							uint8(math.Min(rayColor.y, 1) * 255),
							uint8(math.Min(rayColor.z, 1) * 255),
							1,
						}
						im.SetRGBA(xx, yy, c)
					}
				}
				wg.Done()
			}(xi, yi)
		}
	}

	wg.Wait()

	// var wg1 sync.WaitGroup

 	// xc := math.Ceil(float64(c.hSize) / math.Sqrt(float64(numChunks)))
 	// yc := math.Ceil(float64(c.vSize) / math.Sqrt(float64(numChunks)))

	// ch := make(chan Work, 2 * numChunks)


	// cc := 0
	// for xi := 0; xi < int(math.Ceil(math.Sqrt(float64(numChunks)))); xi++ {
	// 	for yi := 0; yi < int(math.Ceil(math.Sqrt(float64(numChunks)))); yi++ {
	// 		cc += 1
	// 		ch <- Work{xi, yi}
	// 	}
	// }

	// for n := 0; n < numWorkers; n++ {
	// 	wg1.Add(1)
	// 	go func() {
	// 		for {
	// 			select {
	// 			case work := <-ch:
	// 				xi := work.xc
	// 				yi := work.yc

	// 				for y := 0; y < int(yc); y++ {
	// 					for x := 0; x < int(xc); x++ {
	// 						xx := int(math.Min(float64(x + xi * int(xc)), float64(c.hSize)))
	// 						yy := int(math.Min(float64(y + yi * int(yc)), float64(c.vSize)))

	// 						rayColor := c.world.ColorAt(
	// 							c.RayFor(xx, yy),
	// 							ReflectionLimit,
	// 						)

	// 						c := color.RGBA{
	// 							uint8(math.Min(rayColor.x, 1) * 255),
	// 							uint8(math.Min(rayColor.y, 1) * 255),
	// 							uint8(math.Min(rayColor.z, 1) * 255),
	// 							1,
	// 						}
	// 						im.SetRGBA(xx, yy, c)
	// 					}
	// 				}
	// 			default:
	// 				wg1.Done()
	// 				return
	// 			}
	// 		}
	// 	}()
	// }

	// wg1.Wait()
}

func (c *Camera)Input(keys [4]bool) {
	if keys[ui.Up] {
		c.currentPosition.y += 0.15
	}

	if keys[ui.Down] {
		c.currentPosition.y -= 0.15
	}

	if keys[ui.Left] {
		c.currentPosition = RotationY(0.04).Transform(c.currentPosition)
	}

	if keys[ui.Right] {
		c.currentPosition = RotationY(-0.04).Transform(c.currentPosition)
	}

	c.SetTransform(
		ViewTransform(
			c.currentPosition,
			NewPoint(0, 1, 0),
			New3Vector(0, 1, 0),
		),
	)
}
