package common

import (
	"math"
	"testing"
)

func TestReflectanceTotalInternalReflection(t *testing.T) {
	shape := NewGlassSphere()
	ray := NewRay(NewPoint(0, 0, math.Sqrt(2)/2), New3Vector(0, 1, 0))

	intersections := NewIntersections(
		&Intersection{shape, -math.Sqrt(2)/2},
		&Intersection{shape, math.Sqrt(2)/2},
	)

	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if Reflectance(precomputed) != 1.0 {
		t.Errorf("Reflectance total internal reflection not working")
	}
}

func TestReflectanceNormal(t *testing.T) {
	shape := NewGlassSphere()
	ray := NewRay(NewPoint(0, 0, 0), New3Vector(0, 1, 0))

	intersections := NewIntersections(
		&Intersection{shape, -1},
		&Intersection{shape, 1},
	)

	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if !floatEqual(Reflectance(precomputed), 0.04) {
		t.Errorf("Reflectance total internal reflection not working")
	}
}

func TestReflectanceGlance(t *testing.T) {
	shape := NewGlassSphere()
	ray := NewRay(NewPoint(0, 0.99, -2), New3Vector(0, 0, 1))

	intersections := NewIntersections(
		&Intersection{shape, 1.8589},
	)

	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if !floatEqual(Reflectance(precomputed), 0.48873) {
		t.Errorf("Reflectance total internal reflection not working")
	}
}
