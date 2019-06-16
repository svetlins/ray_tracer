package common

import (
	_ "fmt"
	"math"
)

type World struct {
	objects []Object
	light   *PointLight
}

func NewWorld() *World {
	return &World{}
}

func NewDefaultWorld() *World {
	sphere1 := NewSphere()

	material := NewMaterial(
		Color(NewColor(0.8, 1.0, 0.6)),
		Diffuse(0.7),
		Specular(0.2),
	)

	sphere1.SetMaterial(material)

	sphere2 := NewSphere()
	sphere2.SetTransform(Scaling(0.5, 0.5, 0.5))

	light := NewPointLight(NewColor(1, 1, 1), NewPoint(-10, 10, -10))

	return &World{
		[]Object{sphere1, sphere2},
		light,
	}
}

func (w *World) AddObject(o Object) {
	w.objects = append(w.objects, o)
}

func (w *World)ObjectAt(i int) Object {
	return w.objects[i]
}

func (w *World) SetLight(l *PointLight) {
	w.light = l
}

func (w *World) Intersect(r Ray) *Intersections {
	is := &Intersections{}

	for _, object := range w.objects {
		is.Merge(object.Intersect(r))
	}

	return is
}

func (w *World) ShadeHit(comps *Precomputed, remaining int) Vector {
	material := comps.Object.Material()

	lighting := Lighting(
		material,
		w.light,
		comps.Object,
		comps.Point,
		comps.EyeVector,
		comps.NormalVector,
		w.IsShadowed(comps.OverPoint),
	)

	if material.transparency > 0 && material.reflective > 0 {
		var reflected Vector

		if simpleMode {
			reflected = Black()
		} else {
			reflected = w.ReflectedColor(comps, remaining)
		}

		refracted := w.RefractedColor(comps, remaining)

		reflectance := Reflectance(comps)

		return lighting.Add(
			reflected.Scale(reflectance),
		).Add(
			refracted.Scale(1 - reflectance),
		)
	} else if material.transparency > 0 {
		refracted := w.RefractedColor(comps, remaining)
		return lighting.Add(refracted)
	} else if material.reflective > 0 {
		reflected := w.ReflectedColor(comps, remaining)
		return lighting.Add(reflected)
	} else {
		return lighting
	}
}

func (w *World) ColorAt(ray Ray, remaining int) Vector {
	intersections := w.Intersect(ray)
	hit := intersections.Hit()

	if hit != nil {
		precomputed := PrepareComputations(hit, ray, intersections)
		return w.ShadeHit(precomputed, remaining)
	} else {
		return Black()
	}
}

func (w *World) IsShadowed(point Vector) bool {
	pointToLight := w.light.position.Subtract(point)

	hit := w.Intersect(NewRay(point, pointToLight.Normalize())).Hit()

	if hit != nil && hit.t < pointToLight.Magnitude() {
		return true
	}

	return false
}

func (w *World) ReflectedColor(precomputed *Precomputed, remaining int) Vector {
	if remaining < 1 {
		return Black()
	}

	if precomputed.Object.Material().reflective > 0 {
		reflectedRay := NewRay(
			precomputed.OverPoint,
			precomputed.ReflectVector,
		)

		return w.ColorAt(reflectedRay, remaining - 1).Scale(
			precomputed.Object.Material().reflective,
		)
	} else {
		return Black()
	}

}

func (w *World)RefractedColor(precomputed *Precomputed, remaining int) Vector {
	if precomputed.Object.Material().transparency < 0.001 ||
		remaining < 1 {
		return Black()
	}

	nRatio := precomputed.RefractiveIndexEnter / precomputed.RefractiveIndexExit
	cosI := precomputed.EyeVector.Dot(precomputed.NormalVector)
	sin2T := nRatio*nRatio * (1 - cosI*cosI)

	cosT := math.Sqrt(1 - sin2T)

	direction := precomputed.NormalVector.Scale(nRatio * cosI - cosT).Subtract(
		precomputed.EyeVector.Scale(nRatio),
	)

	refractedRay := NewRay(precomputed.UnderPoint, direction)

	return w.ColorAt(refractedRay, remaining - 1).Scale(
			precomputed.Object.Material().transparency,
	)
}
