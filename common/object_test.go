package common

import (
	"math"
	"testing"
)

func TestNewSphere(t *testing.T) {
	sphere := NewSphere()

	if !sphere.transform.Eq(NewIdentity4()) {
		t.Errorf("NewSphere not working")
	}

	if !sphere.material.Equal(NewMaterial()) {
		t.Errorf("NewSphere not working")
	}
}

func TestNewGlassSphere(t *testing.T) {
	sphere := NewGlassSphere()

	if sphere.material.transparency != 1.0 {
		t.Errorf("New Glass Sphere not working")
	}

	if sphere.material.refractiveIndex != 1.5 {
		t.Errorf("New Glass Sphere not working")
	}
}

func TestSphereTransform(t *testing.T) {
	sphere := NewSphere()
	sphere.SetTransform(Scaling(0, 1, 0))

	if !sphere.transform.Eq(Scaling(0, 1, 0)) {
		t.Errorf("Sphere set transform not working")
	}

}

func TestSphereMaterial(t *testing.T) {
	sphere := NewSphere()
	material := NewMaterial()
	material.ambient = 0.5
	sphere.SetMaterial(material)

	if sphere.material.ambient != 0.5 {
		t.Errorf("Sphere set material not working")
	}

}

func TestSphereIntersect1(t *testing.T) {
	sphere := NewSphere()
	ray := NewRay(NewPoint(0, 0, -5), New3Vector(0, 0, 1))
	xs := sphere.Intersect(ray)

	if xs.Len() != 2 {
		t.Errorf("Sphere intersect not working")
	}

	if xs.At(0).t != 4.0 {
		t.Errorf("Sphere intersect not working")
	}

	if xs.At(1).t != 6.0 {
		t.Errorf("Sphere intersect not working")
	}

	if xs.At(0).object.Id() != sphere.Id() {
		t.Errorf("Sphere intersect not working")
	}
}

func TestSphereIntersect2(t *testing.T) {
	sphere := NewSphere()
	ray := NewRay(NewPoint(0, 1, -5), New3Vector(0, 0, 1))
	xs := sphere.Intersect(ray)

	if xs.Len() != 2 {
		t.Errorf("Sphere intersect not working")
	}

	if xs.At(0).t != 5.0 {
		t.Errorf("Sphere intersect not working")
	}

	if xs.At(1).t != 5.0 {
		t.Errorf("Sphere intersect not working")
	}
}

func TestSphereIntersect3(t *testing.T) {
	sphere := NewSphere()
	ray := NewRay(NewPoint(0, 2, -5), New3Vector(0, 0, 1))
	xs := sphere.Intersect(ray)

	if xs.Len() != 0 {
		t.Errorf("Sphere intersect not working")
	}
}

func TestSphereIntersect4(t *testing.T) {
	sphere := NewSphere()
	ray := NewRay(NewPoint(0, 0, 0), New3Vector(0, 0, 1))
	xs := sphere.Intersect(ray)

	if xs.Len() != 2 {
		t.Errorf("Sphere intersect not working")
	}

	if xs.At(0).t != -1.0 {
		t.Errorf("Sphere intersect not working")
	}

	if xs.At(1).t != 1.0 {
		t.Errorf("Sphere intersect not working")
	}
}

func TestSphereIntersect5(t *testing.T) {
	sphere := NewSphere()
	ray := NewRay(NewPoint(0, 0, 5), New3Vector(0, 0, 1))
	xs := sphere.Intersect(ray)

	if xs.Len() != 2 {
		t.Errorf("Sphere intersect not working")
	}

	if xs.At(0).t != -6.0 {
		t.Errorf("Sphere intersect not working")
	}

	if xs.At(1).t != -4.0 {
		t.Errorf("Sphere intersect not working")
	}
}

func TestHit1(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{s, 1}
	i2 := Intersection{s, 2}
	intersections := Intersections{
		[]*Intersection{&i1, &i2},
	}

	if intersections.Hit() != &i1 {
		t.Errorf("Intersection Hit not working")
	}
}

func TestNewIntersection(t *testing.T) {
	intersections := NewIntersections(
		&Intersection{nil, 2},
		&Intersection{nil, -1},
	)

	if intersections.Len() != 2 {
		t.Errorf("New Intersection not working")
	}

	if intersections.At(0).t != -1 {
		t.Errorf("New Intersection not working")
	}
}

func TestIntersectionMerge(t *testing.T) {
	intersections := NewIntersections(
		&Intersection{nil, 2},
		&Intersection{nil, -1},
	)

	intersections.Merge(
		NewIntersections(
			&Intersection{nil, -5},
			&Intersection{nil, 10},
		),
	)

	if intersections.Len() != 4 {
		t.Errorf("New Merge not working")
	}

	if intersections.At(0).t != -5 {
		t.Errorf("New Merge not working")
	}

	if intersections.At(3).t != 10 {
		t.Errorf("New Merge not working")
	}
}

func TestHit2(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{s, -1}
	i2 := Intersection{s, 2}
	intersections := Intersections{
		[]*Intersection{&i1, &i2},
	}

	if intersections.Hit() != &i2 {
		t.Errorf("Intersection Hit not working")
	}
}

func TestHit3(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{s, -2}
	i2 := Intersection{s, -1}
	intersections := Intersections{
		[]*Intersection{&i1, &i2},
	}

	if intersections.Hit() != nil {
		t.Errorf("Intersection Hit not working")
	}
}

func TestHit4(t *testing.T) {
	s := NewSphere()
	i1 := &Intersection{s, 5}
	i2 := &Intersection{s, 7}
	i3 := &Intersection{s, 2}
	i4 := &Intersection{s, -3}
	intersections := NewIntersections(i1, i2, i3, i4)

	if intersections.Hit() != i3 {
		t.Errorf("Intersection Hit not working")
	}
}

func TestSphereIntersectTransformed1(t *testing.T) {
	ray := NewRay(
		NewPoint(0, 0, -5),
		New3Vector(0, 0, 1),
	)

	sphere := NewSphere()
	sphere.SetTransform(Scaling(2, 2, 2))

	xs := sphere.Intersect(ray)

	if xs.Len() != 2 {
		t.Errorf("Sphere Intersect Transformed not working")
	}

	if xs.At(0).t != 3.0 {
		t.Errorf("Sphere Intersect Transformed not working (%f)", xs.At(0).t)
	}

	if xs.At(1).t != 7.0 {
		t.Errorf("Sphere Intersect Transformed not working")
	}
}

func TestSphereIntersectTransformed2(t *testing.T) {
	ray := NewRay(
		NewPoint(0, 0, -5),
		New3Vector(0, 0, 1),
	)

	sphere := NewSphere()
	sphere.SetTransform(Translation(5, 0, 0))

	xs := sphere.Intersect(ray)

	if xs.Len() != 0 {
		t.Errorf("Sphere Intersect Transformed not working")
	}
}

func TestSphereNormals(t *testing.T) {
	sphere := NewSphere()

	if !sphere.NormalAt(NewPoint(1, 0, 0)).Eq(New3Vector(1, 0, 0)) {
		t.Errorf("Sphere Normal not working")
	}

	if !sphere.NormalAt(NewPoint(0, 1, 0)).Eq(New3Vector(0, 1, 0)) {
		t.Errorf("Sphere Normal not working")
	}

	if !sphere.NormalAt(NewPoint(0, 0, 1)).Eq(New3Vector(0, 0, 1)) {
		t.Errorf("Sphere Normal not working")
	}

	v := math.Sqrt(3) / 3

	if !sphere.NormalAt(NewPoint(v, v, v)).Eq(New3Vector(v, v, v)) {
		t.Errorf("Sphere Normal not working")
	}

	if !sphere.NormalAt(NewPoint(v, v, v)).Eq(New3Vector(v, v, v).Normalize()) {
		t.Errorf("Sphere Normal not normalized")
	}
}

func TestTransformedSphereNormals(t *testing.T) {
	sphere := NewSphere()

	sphere.SetTransform(Translation(0, 1, 0))

	if !sphere.NormalAt(NewPoint(0, 1.70711, -0.70711)).Eq(New3Vector(0, 0.70711, -0.70711)) {
		t.Errorf("Sphere Normal not working 1")
	}

	sphere.SetTransform(Scaling(1, 0.5, 1).Multiply(RotationZ(math.Pi / 5)))

	if !sphere.NormalAt(NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)).Eq(New3Vector(0, 0.97014, -0.24254)) {
		t.Errorf("Sphere Normal not working 2")
	}

}
