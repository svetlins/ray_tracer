package main

import (
	"image"
	"github.com/svetlins/srt/common"
	"math"
	"testing"
)

var result interface{}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		world := common.NewWorld()

		world.SetLight(
			common.NewPointLight(
				common.NewColor(1, 1, 1), common.NewPoint(-10, 10, -10),
			),
		)

		middleSphere := common.NewSphere()
		middleSphere.SetTransform(
			common.Translation(-0.5, 1, 0.5),
		)

		rightSphere := common.NewSphere()
		rightSphere.SetTransform(
			common.Translation(1, 1, 4).Multiply(common.Scaling(0.8, 0.8, 0.8)),
		)

		leftSphere := common.NewSphere()
		leftSphere.SetTransform(
			common.Translation(-2.0, 0.33, 3).Multiply(common.Scaling(0.33, 0.33, 0.33)),
		)
		leftMaterial := common.NewMaterial(
			common.Color(common.Blue),
		)
		leftSphere.SetMaterial(leftMaterial)

		floor := common.NewPlane()
		wall := common.NewPlane()
		wall.SetTransform(
			common.Translation(0, 0, 12).Multiply(common.RotationY(1).Multiply(common.RotationX(math.Pi/2))),
		)

		stripe := common.NewPattern(
			common.PatternColors(common.Yellow, common.Red),
			common.PatternTransform(common.Translation(0, 0.01, 0).Multiply(common.RotationY(0)).Multiply(common.Scaling(2,2,2))),
			common.PatternValue(common.Gradient()),
		)

		mat := common.NewMaterial(
			common.MaterialPattern(stripe),
			common.Specular(0.3),
			common.Ambient(0.2),
			common.Diffuse(0.9),
		)

		floor.SetMaterial(mat)
		wall.SetMaterial(mat)

		sphereMat := common.NewMaterial(
			common.Color(common.Green),
			common.Ambient(0.3),
			common.Refractive(1.1),
			common.Transparency(0.6),
			common.Reflective(0.9),
			common.Diffuse(0.2),
			common.Specular(1),
			common.Shininess(300),
		)

		middleSphere.SetMaterial(sphereMat)

		world.AddObject(middleSphere)
		world.AddObject(leftSphere)
		// world.AddObject(rightSphere)
		world.AddObject(floor)
		world.AddObject(wall)

		// camera := common.NewCamera(80, 40, math.Pi/3)
		camera := common.NewCamera(900, 900, math.Pi/3)
		common.EnableSimpleMode()

		camera.SetTransform(
			common.ViewTransform(
				common.NewPoint(0, 1.5, -5),
				common.NewPoint(0, 1, 0),
				common.New3Vector(0, 1, 0),
			),
		)

		camera.SetWorld(world)

		im := image.NewRGBA(image.Rect(0, 0, 900, 900))

		camera.RenderToSequential(im)

		result = im
	}
}
