package common

import (
	_ "fmt"
	"math"
	_ "math/rand"
)

type PointLight struct {
	intensity Vector
	position  Vector
}

func NewPointLight(intensity, position Vector) *PointLight {
	return &PointLight{intensity, position}
}

func (l1 *PointLight) Eq(l2 *PointLight) bool {
	return l1.intensity.Eq(l2.intensity) && l1.position.Eq(l2.position)
}

func Lighting(m Material, l *PointLight, object Object, position, eyeVector, normalVector Vector, inShadow bool) Vector {
	var ambient, diffuse, specular Vector

	var materialColor Vector

	if m.pattern != nil {
		materialColor = m.pattern.AtObject(object, position)
	} else {
		materialColor = m.color
	}

	effectiveColor := materialColor.Multiply(l.intensity)
	lightVector := l.position.Subtract(position).Normalize()
	ambient = effectiveColor.Scale(m.ambient)

	lightDotNormal := lightVector.Dot(normalVector)

	if lightDotNormal < 0 || inShadow {
		diffuse = NewColor(0, 0, 0)
		specular = NewColor(0, 0, 0)
	} else {
		diffuse = effectiveColor.Scale(m.diffuse * lightDotNormal)
		reflectVector := lightVector.Scale(-1).Reflect(normalVector)
		reflectDotEye := reflectVector.Dot(eyeVector)

		if reflectDotEye <= 0 {
			specular = NewColor(0, 0, 0)
		} else {
			factor := math.Pow(reflectDotEye, m.shininess)
			specular = l.intensity.Scale(m.specular * factor)
		}

	}

	return ambient.Add(diffuse).Add(specular)
}
