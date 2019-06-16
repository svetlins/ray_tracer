package common

import (
	"testing"
)

func TestNewRay(t *testing.T) {
	origin := NewPoint(1, 2, 3)
	direction := New3Vector(4, 5, 6)
	ray := NewRay(origin, direction)

	if !ray.origin.Eq(origin) {
		t.Errorf("New ray not working")
	}

	if !ray.direction.Eq(direction) {
		t.Errorf("New ray not working")
	}
}

func TestRayPosition(t *testing.T) {
	origin := NewPoint(2, 3, 4)
	direction := New3Vector(1, 0, 0)
	ray := NewRay(origin, direction)

	if !ray.Position(0).Eq(NewPoint(2, 3, 4)) {
		t.Errorf("Ray position not working")
	}

	if !ray.Position(1).Eq(NewPoint(3, 3, 4)) {
		t.Errorf("Ray position not working")
	}

	if !ray.Position(-1).Eq(NewPoint(1, 3, 4)) {
		t.Errorf("Ray position not working")
	}

	if !ray.Position(2.5).Eq(NewPoint(4.5, 3, 4)) {
		t.Errorf("Ray position not working")
	}
}

func TestRayTransform1(t *testing.T) {
	ray := NewRay(NewPoint(1, 2, 3), New3Vector(0, 1, 0))
	translation := Translation(3, 4, 5)

	tranformedRay := ray.Transform(translation)

	if !tranformedRay.origin.Eq(NewPoint(4, 6, 8)) {
		t.Errorf("Ray transform not working")
	}

	if !tranformedRay.direction.Eq(New3Vector(0, 1, 0)) {
		t.Errorf("Ray transform not working")
	}
}

func TestRayTransform2(t *testing.T) {
	ray := NewRay(NewPoint(1, 2, 3), New3Vector(0, 1, 0))
	translation := Scaling(2, 3, 4)

	tranformedRay := ray.Transform(translation)

	if !tranformedRay.origin.Eq(NewPoint(2, 6, 12)) {
		t.Errorf("Ray transform not working")
	}

	if !tranformedRay.direction.Eq(New3Vector(0, 3, 0)) {
		t.Errorf("Ray transform not working")
	}
}
