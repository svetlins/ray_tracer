package common

import (
	"math"
	"testing"
	"time"
)

func TestNewWorld(t *testing.T) {
	world := NewWorld()

	if len(world.objects) != 0 {
		t.Errorf("New World not working")
	}

	if world.light != nil {
		t.Errorf("New World not working")
	}
}

func TestDefaultWorld(t *testing.T) {
	world := NewDefaultWorld()

	if world == nil {
		t.Errorf("New Default World not working")
	}

	if !world.light.Eq(NewPointLight(NewColor(1, 1, 1), NewPoint(-10, 10, -10))) {
		t.Errorf("New Default World not working")
	}

	sphere1 := NewSphere()
	material := NewMaterial(
		Color(NewColor(0.8, 1.0, 0.6)),
		Diffuse(0.7),
		Specular(0.2),
	)
	sphere1.SetMaterial(material)

	if !world.objects[0].(*Sphere).Eq(sphere1) {
		t.Errorf("New Default World not working")
	}

	sphere2 := NewSphere()
	sphere2.SetTransform(Scaling(0.5, 0.5, 0.5))

	if !world.objects[1].(*Sphere).Eq(sphere2) {
		t.Errorf("New Default World not working")
	}
}

func TestWorldIntersect(t *testing.T) {
	world := NewDefaultWorld()

	ray := NewRay(NewPoint(0, 0, -5), New3Vector(0, 0, 1))

	intersections := world.Intersect(ray)

	if intersections.Len() != 4 {
		t.Errorf("World intersect not working")
	}

	if intersections.At(0).t != 4 {
		t.Errorf("World intersect not working")
	}

	if intersections.At(1).t != 4.5 {
		t.Errorf("World intersect not working")
	}

	if intersections.At(2).t != 5.5 {
		t.Errorf("World intersect not working")
	}

	if intersections.At(3).t != 6 {
		t.Errorf("World intersect not working")
	}
}

func TestWorldShading(t *testing.T) {
	world := NewDefaultWorld()
	ray := NewRay(NewPoint(0, 0, -5), New3Vector(0, 0, 1))
	shape := world.objects[0]
	intersections := NewIntersections(&Intersection{shape, 4})
	comps := PrepareComputations(intersections.Hit(), ray, intersections)

	if world.ShadeHit(comps, 1).NotEq(NewColor(0.38066, 0.47583, 0.2855)) {
		t.Errorf("World shading not working")
	}
}

func TestWorldShadingInside(t *testing.T) {
	world := NewDefaultWorld()
	world.SetLight(NewPointLight(NewColor(1, 1, 1), NewPoint(0, 0.25, 0)))
	ray := NewRay(NewPoint(0, 0, 0), New3Vector(0, 0, 1))
	shape := world.objects[1]
	intersections := NewIntersections(&Intersection{shape, 0.5})
	comps := PrepareComputations(intersections.Hit(), ray, intersections)

	if world.ShadeHit(comps, 1).NotEq(NewColor(0.90498, 0.90498, 0.90498)) {
		t.Errorf("World shading not working")
	}
}

func TestWorldShadingShadow(t *testing.T) {
	world := NewWorld()
	world.SetLight(NewPointLight(NewColor(1, 1, 1), NewPoint(0, 0, -10)))
	world.AddObject(NewSphere())
	otherSphere := NewSphere()

	otherSphere.SetTransform(
		Translation(0, 0, 10),
	)

	world.AddObject(otherSphere)

	ray := NewRay(NewPoint(0, 0, 5), New3Vector(0, 0, 1))

	intersections := NewIntersections(&Intersection{otherSphere, 4})

	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if world.ShadeHit(precomputed, 1).NotEq(NewColor(0.1, 0.1, 0.1)) {
		t.Errorf("World ShadeHit Shadow not working")
	}

}

func TestWorldColorAt1(t *testing.T) {
	world := NewDefaultWorld()

	ray1 := NewRay(NewPoint(0, 0, -5), New3Vector(0, 1, 0))

	if world.ColorAt(ray1, 1).NotEq(NewColor(0, 0, 0)) {
		t.Errorf("World ColorAt not working")
	}

	ray2 := NewRay(NewPoint(0, 0, -5), New3Vector(0, 0, 1))
	if world.ColorAt(ray2, 1).NotEq(NewColor(0.38066, 0.47583, 0.2855)) {
		t.Errorf("World ColorAt not working")
	}
}

func TestWorldColorAtInside(t *testing.T) {
	world := NewDefaultWorld()

	world.ObjectAt(0).SetMaterial(NewMaterial(
		Color(NewColor(0.8, 1.0, 0.6)),
		Diffuse(0.7),
		Specular(0.2),
		Ambient(1),
	))

	world.ObjectAt(1).SetMaterial(NewMaterial(
		Ambient(1),
	))

	ray := NewRay(NewPoint(0, 0, 0.75), New3Vector(0, 0, -1))

	if world.ColorAt(ray, 1).NotEq(world.objects[1].Material().color) {
		t.Errorf("World ColorAt inside not working")
	}
}

func TestIsShadowed1(t *testing.T) {
	world := NewDefaultWorld()
	point := NewPoint(0, 10, 0)

	if world.IsShadowed(point) {
		t.Errorf("World Shadowed not working")
	}
}

func TestIsShadowed2(t *testing.T) {
	world := NewDefaultWorld()
	point := NewPoint(10, -10, 10)

	if !world.IsShadowed(point) {
		t.Errorf("World Shadowed not working")
	}
}

func TestIsShadowed3(t *testing.T) {
	world := NewDefaultWorld()
	point := NewPoint(-20, 20, -20)

	if world.IsShadowed(point) {
		t.Errorf("World Shadowed not working")
	}
}

func TestIsShadowed4(t *testing.T) {
	world := NewDefaultWorld()
	point := NewPoint(-2, 2, -2)

	if world.IsShadowed(point) {
		t.Errorf("World Shadowed not working")
	}
}

func TestReflectedColorOfNonReflective(t *testing.T) {
	world := NewDefaultWorld()
	ray := NewRay(NewPoint(0, 0, 0), New3Vector(0, 0, 1))
	shape := world.objects[1]
	material := NewMaterial(Ambient(1))
	shape.SetMaterial(material)
	intersections := NewIntersections(&Intersection{shape, 1})

	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if world.ReflectedColor(precomputed, 1).NotEq(Black()) {
		t.Errorf("World Reflected of non reflective not working")
	}
}

func TestReflectedColor(t *testing.T) {
	world := NewDefaultWorld()
	plane := NewPlane()
	material := NewMaterial(Reflective(0.5))
	plane.SetMaterial(material)
	plane.SetTransform(Translation(0, -1, 0))
	world.AddObject(plane)

	ray := NewRay(NewPoint(0, 0, -3), New3Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))

	intersections := NewIntersections(&Intersection{plane, math.Sqrt(2)})

	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if world.ReflectedColor(precomputed, 1).NotEq(NewColor(0.190332, 0.237915, 0.1427492)) {
		t.Errorf("World Reflected of non reflective not working")
	}
}

func TestShadeHitReflection(t *testing.T) {
	world := NewDefaultWorld()
	plane := NewPlane()
	material := NewMaterial(Reflective(0.5))
	plane.SetMaterial(material)
	plane.SetTransform(Translation(0, -1, 0))
	world.AddObject(plane)

	ray := NewRay(NewPoint(0, 0, -3), New3Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))

	intersections := NewIntersections(&Intersection{plane, math.Sqrt(2)})

	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if world.ShadeHit(precomputed, 1).NotEq(NewColor(0.8767577, 0.9243407, 0.829174)) {
		t.Errorf("World Reflected of non reflective not working")
	}
}

func TestShadeHitInfiniteReflection(t *testing.T) {
	world := NewWorld()
	world.SetLight(NewPointLight(White(), NewPoint(0, 0, 0)))

	material := NewMaterial(Reflective(1))

	lowerPlane := NewPlane()
	lowerPlane.SetMaterial(material)
	lowerPlane.SetTransform(Translation(0, -1, 0))
	world.AddObject(lowerPlane)

	upperPlane := NewPlane()
	upperPlane.SetMaterial(material)
	upperPlane.SetTransform(Translation(0, 1, 0))
	world.AddObject(upperPlane)

	ray := NewRay(NewPoint(0, 0, 0), New3Vector(0, 1, 0))

	selector := make(chan bool)

	go func() {
		world.ColorAt(ray, 1)
		selector <- true
	}()

	select {
	case <-selector:
	case <-time.After(1 * time.Millisecond):
		t.Errorf("Possible infinite recursion on reflective surfaces")
	}
}

func TestReflectedColorNoRemaining(t *testing.T) {
	world := NewDefaultWorld()
	plane := NewPlane()
	material := NewMaterial(Reflective(0.5))
	plane.SetMaterial(material)
	plane.SetTransform(Translation(0, -1, 0))
	world.AddObject(plane)

	ray := NewRay(NewPoint(0, 0, -3), New3Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))

	intersections := NewIntersections(&Intersection{plane, math.Sqrt(2)})

	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if world.ReflectedColor(precomputed, 0).NotEq(Black()) {
		t.Errorf("World Reflected of non reflective not working")
	}
}

func TestRefractedColorNoTransparency(t *testing.T) {
	world := NewDefaultWorld()
	shape := world.ObjectAt(0)

	ray := NewRay(NewPoint(0, 0, -5), New3Vector(0, 0, 1))

	intersections := NewIntersections(
		&Intersection{shape, 4},
		&Intersection{shape, 6},
	)

	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if world.RefractedColor(precomputed, ReflectionLimit).NotEq(Black()) {
		t.Errorf("World Reflected of non reflective not working")
	}
}

func TestRefractedColorAtLimit(t *testing.T) {
	world := NewDefaultWorld()
	shape := world.ObjectAt(0)
	material := NewMaterial(
		Transparency(1),
		Refractive(1.5),
	)
	shape.SetMaterial(material)

	ray := NewRay(NewPoint(0, 0, -5), New3Vector(0, 0, 1))

	intersections := NewIntersections(
		&Intersection{shape, 4},
		&Intersection{shape, 6},
	)

	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if world.RefractedColor(precomputed, 0).NotEq(Black()) {
		t.Errorf("World Reflected of non reflective not working")
	}
}

func TestRefractedTotalInternalReflection(t *testing.T) {
	world := NewDefaultWorld()
	shape := world.ObjectAt(0)
	material := NewMaterial(
		Transparency(1),
		Refractive(1.5),
	)
	shape.SetMaterial(material)

	ray := NewRay(NewPoint(0, 0, math.Sqrt(2)/2), New3Vector(0, 1, 0))

	intersections := NewIntersections(
		&Intersection{shape, -math.Sqrt(2)/2},
		&Intersection{shape, math.Sqrt(2)/2},
	)

	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if world.RefractedColor(precomputed, ReflectionLimit).NotEq(Black()) {
		t.Errorf("World Reflected of non reflective not working")
	}
}

func TestRefractedColor(t *testing.T) {
	world := NewDefaultWorld()
	shape1 := world.ObjectAt(0)
	pattern1 := NewPattern()
	pattern1.SetValueFunc(TestPattern())
	material1 := NewMaterial(
		Ambient(1),
		MaterialPattern(pattern1),
	)
	shape1.SetMaterial(material1)

	shape2 := world.ObjectAt(1)
	material2 := NewMaterial(
		Transparency(1),
		Refractive(1.5),
	)
	shape2.SetMaterial(material2)

	ray := NewRay(NewPoint(0, 0, 0.1), New3Vector(0, 1, 0))

	intersections := NewIntersections(
		&Intersection{shape1, -0.9899},
		&Intersection{shape2, -0.4899},
		&Intersection{shape2, 0.4899},
		&Intersection{shape1, 0.9899},
	)

	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if world.RefractedColor(precomputed, ReflectionLimit).NotEq(NewColor(0, 0.99888, 0.047219)) {
		t.Errorf("World Reflected of non reflective not working")
	}
}

func TestRefractedColorShadeHit(t *testing.T) {
	world := NewDefaultWorld()
	floor := NewPlane()
	floor.SetTransform(Translation(0, -1, 0))
	floor.SetMaterial(NewMaterial(
		Transparency(0.5),
		Refractive(1.5),
	))
	world.AddObject(floor)

	ball := NewSphere()
	ball.SetMaterial(NewMaterial(
		Color(PureRed),
		Ambient(0.5),
	))
	ball.SetTransform(Translation(0, -3.5, -0.5))
	world.AddObject(ball)

	ray := NewRay(NewPoint(0, 0, -3), New3Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))

	intersections := NewIntersections(
		&Intersection{floor, math.Sqrt(2)},
	)

	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if world.ShadeHit(precomputed, 5).NotEq(NewColor(0.93642, 0.68642, 0.68642)) {
		t.Errorf("World refracted shade hit not working")
	}
}

func TestReflectanceShadeHit(t *testing.T) {
	world := NewDefaultWorld()
	floor := NewPlane()
	floor.SetTransform(Translation(0, -1, 0))
	floor.SetMaterial(NewMaterial(
		Reflective(0.5),
		Transparency(0.5),
		Refractive(1.5),
	))
	world.AddObject(floor)

	ball := NewSphere()
	ball.SetMaterial(NewMaterial(
		Color(PureRed),
		Ambient(0.5),
	))

	ball.SetTransform(Translation(0, -3.5, -0.5))
	world.AddObject(ball)

	ray := NewRay(NewPoint(0, 0, -3), New3Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))

	intersections := NewIntersections(
		&Intersection{floor, math.Sqrt(2)},
	)

	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if world.ShadeHit(precomputed, 5).NotEq(NewColor(0.93391, 0.69643, 0.69243)) {
		t.Errorf("World reflectance shade hit not working")
	}
}
