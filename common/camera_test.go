package common

import (
	"math"
	"testing"
)

func TestNewCamera(t *testing.T) {
	camera := NewCamera(
		160,
		120,
		math.Pi/2,
	)

	if (camera.hSize != 160) || (camera.vSize != 120) || (camera.fov != math.Pi/2) {
		t.Errorf("New Camera not working")
	}
}

func TestCameraPixelSize(t *testing.T) {
	cameraH := NewCamera(200, 125, math.Pi/2)

	if cameraH.pixelSize() != 0.01 {
		t.Errorf("Camera pixel size not working")
	}

	cameraV := NewCamera(125, 200, math.Pi/2)

	if cameraV.pixelSize() != 0.01 {
		t.Errorf("Camera pixel size not working")
	}
}

func TestCameraRay1(t *testing.T) {
	camera := NewCamera(201, 101, math.Pi/2)

	ray := camera.RayFor(100, 50)

	if ray.origin.NotEq(NewPoint(0, 0, 0)) {
		t.Errorf("Camera RayFor not working")
	}

	if ray.direction.NotEq(New3Vector(0, 0, -1)) {
		t.Errorf("Camera RayFor not working")
	}
}

func TestCameraRay2(t *testing.T) {
	camera := NewCamera(201, 101, math.Pi/2)

	ray := camera.RayFor(0, 0)

	if ray.origin.NotEq(NewPoint(0, 0, 0)) {
		t.Errorf("Camera RayFor not working")
	}

	if ray.direction.NotEq(New3Vector(0.66519, 0.33259, -0.66851)) {
		t.Errorf("Camera RayFor not working")
	}
}

func TestCameraRay3(t *testing.T) {
	camera := NewCamera(201, 101, math.Pi/2)
	camera.SetTransform(
		RotationY(math.Pi / 4).Multiply(Translation(0, -2, 5)),
	)

	ray := camera.RayFor(100, 50)

	if ray.origin.NotEq(NewPoint(0, 2, -5)) {
		t.Errorf("Camera RayFor not working")
	}

	if ray.direction.NotEq(New3Vector(math.Sqrt(2)/2, 0, -math.Sqrt(2)/2)) {
		t.Errorf("Camera RayFor not working")
	}
}

func TestCameraRender(t *testing.T) {
	world := NewDefaultWorld()
	camera := NewCamera(11, 11, math.Pi/2)

	camera.SetTransform(
		ViewTransform(
			NewPoint(0, 0, -5),
			NewPoint(0, 0, 0),
			New3Vector(0, 1, 0),
		),
	)

	image := camera.Render(world)

	if image.GetPixel(5, 5).NotEq(NewColor(0.38066, 0.47583, 0.2855)) {
		t.Errorf("Camera Render not working")
	}
}
