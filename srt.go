package main

import (
	"fmt"
	"math"
	"runtime"

	"github.com/svetlins/srt/common"
	"github.com/svetlins/srt/ui"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	renderInUi()
}

func renderDefaultWorld() {
	world := common.NewDefaultWorld()
	camera := common.NewCamera(640, 400, math.Pi/2)

	camera.SetTransform(
		common.ViewTransform(
			common.NewPoint(0, 0, -5),
			common.NewPoint(0, 0, 0),
			common.New3Vector(0, 1, 0),
		),
	)

	image := camera.Render(world)

	fmt.Println(image.PPM())
}

func renderInUi() {
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
		common.Translation(0, 0, 12).Multiply(common.RotationY(1).Multiply(common.RotationX(math.Pi / 2))),
	)

	stripe := common.NewPattern(
		common.PatternColors(common.Red, common.Yellow),
		common.PatternTransform(
			common.Scaling(1, 1, 0.5),
			common.RotationY(1),
			common.Translation(1, 0.01, 3),
		),
		common.PatternValue(common.Ring()),
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
		common.Refractive(1.4),
		common.Transparency(0.6),
		common.Reflective(0.9),
		common.Diffuse(0.2),
		common.Specular(1),
		common.Shininess(300),
	)

	middleSphere.SetMaterial(sphereMat)

	world.AddObject(rightSphere)
	world.AddObject(middleSphere)
	world.AddObject(leftSphere)
	world.AddObject(floor)
	world.AddObject(wall)

	camera := common.NewCamera(240, 150, math.Pi/3)

	camera.SetTransform(
		common.ViewTransform(
			common.NewPoint(-0.5, 1, -4),
			common.NewPoint(-0.5, 1, 0.5),
			common.New3Vector(0, 1, 0),
		),
	)

	camera.SetWorld(world)

	ui.StartUI(camera)
}
