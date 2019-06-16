package common

import (
	"testing"
)

func TestPlaneNormal(t *testing.T) {
	plane := NewPlane()

	if plane.NormalAt(NewPoint(0, 0, 0)).NotEq(New3Vector(0, 1, 0)) {
		t.Errorf("Plane Normal not working")
	}

	if plane.NormalAt(NewPoint(10, 0, 30)).NotEq(New3Vector(0, 1, 0)) {
		t.Errorf("Plane Normal not working")
	}

	if plane.NormalAt(NewPoint(200, 0, -5)).NotEq(New3Vector(0, 1, 0)) {
		t.Errorf("Plane Normal not working")
	}
}

func TestPlaneIntersectParallel(t *testing.T) {
	plane := NewPlane()
	ray := NewRay(NewPoint(0, 10, 0), New3Vector(0, 0, 1))

	if plane.Intersect(ray).Len() != 0 {
		t.Errorf("Plane Intersect with parallel not working")
	}
}

func TestPlaneIntersectCoplanar(t *testing.T) {
	plane := NewPlane()
	ray := NewRay(NewPoint(0, 0, 0), New3Vector(0, 0, 1))

	if plane.Intersect(ray).Len() != 0 {
		t.Errorf("Plane Intersect coplanar not working")
	}
}

func TestPlaneIntersectAbove(t *testing.T) {
	plane := NewPlane()
	ray := NewRay(NewPoint(0, 1, 0), New3Vector(0, -1, 0))

	if plane.Intersect(ray).Len() != 1 {
		t.Errorf("Plane Intersect coplanar not working")
	}

	if plane.Intersect(ray).At(0).object != plane {
		t.Errorf("Plane Intersect coplanar not working")
	}
}

func TestPlaneIntersectBelow(t *testing.T) {
	plane := NewPlane()
	ray := NewRay(NewPoint(0, -1, 0), New3Vector(0, 1, 0))

	if plane.Intersect(ray).Len() != 1 {
		t.Errorf("Plane Intersect coplanar not working")
	}

	if plane.Intersect(ray).At(0).object != plane {
		t.Errorf("Plane Intersect coplanar not working")
	}
}
