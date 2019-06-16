package common

import (
	"math"
	"testing"
)

func TestNewPointLight(t *testing.T) {
	intensity := NewColor(1, 1, 1)
	position := NewPoint(0, 0, 0)

	light := NewPointLight(intensity, position)

	if !light.intensity.Eq(intensity) || !light.position.Eq(position) {
		t.Errorf("New Point Light not working")
	}
}

func TestLighting1(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)

	eyeVector := New3Vector(0, 0, -1)
	normalVector := New3Vector(0, 0, -1)

	light := NewPointLight(NewColor(1, 1, 1), NewPoint(0, 0, -10))

	lighting := Lighting(
		material,
		light,
		NewSphere(),
		position,
		eyeVector,
		normalVector,
		false,
	)

	if !lighting.Eq(NewColor(1.9, 1.9, 1.9)) {
		t.Errorf("Lighting not working")
	}
}

func TestLighting2(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)

	eyeVector := New3Vector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalVector := New3Vector(0, 0, -1)

	light := NewPointLight(NewColor(1, 1, 1), NewPoint(0, 0, -10))

	lighting := Lighting(
		material,
		light,
		NewSphere(),
		position,
		eyeVector,
		normalVector,
		false,
	)

	if !lighting.Eq(NewColor(1, 1, 1)) {
		t.Errorf("Lighting not working")
	}
}

func TestLighting3(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)

	eyeVector := New3Vector(0, 0, -1)
	normalVector := New3Vector(0, 0, -1)

	light := NewPointLight(NewColor(1, 1, 1), NewPoint(0, 10, -10))

	lighting := Lighting(
		material,
		light,
		NewSphere(),
		position,
		eyeVector,
		normalVector,
		false,
	)

	if !lighting.Eq(NewColor(0.7364, 0.7364, 0.7364)) {
		t.Errorf("Lighting not working")
	}
}

func TestLighting4(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)

	eyeVector := New3Vector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalVector := New3Vector(0, 0, -1)

	light := NewPointLight(NewColor(1, 1, 1), NewPoint(0, 10, -10))

	lighting := Lighting(
		material,
		light,
		NewSphere(),
		position,
		eyeVector,
		normalVector,
		false,
	)

	if !lighting.Eq(NewColor(1.6364, 1.6364, 1.6364)) {
		t.Errorf("Lighting not working")
	}
}

func TestLighting5(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)

	eyeVector := New3Vector(0, 0, -1)
	normalVector := New3Vector(0, 0, -1)

	light := NewPointLight(NewColor(1, 1, 1), NewPoint(0, 0, 10))

	lighting := Lighting(
		material,
		light,
		NewSphere(),
		position,
		eyeVector,
		normalVector,
		false,
	)

	if !lighting.Eq(NewColor(0.1, 0.1, 0.1)) {
		t.Errorf("Lighting not working")
	}
}

func TestLightingShadow(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)

	eyeVector := New3Vector(0, 0, -1)
	normalVector := New3Vector(0, 0, -1)
	light := NewPointLight(NewColor(1, 1, 1), NewPoint(0, 0, -10))
	inShadow := true

	result := Lighting(
		material,
		light,
		NewSphere(),
		position,
		eyeVector,
		normalVector,
		inShadow,
	)

	if result.NotEq(NewColor(0.1, 0.1, 0.1)) {
		t.Errorf("Lighting Shadow not working")
	}
}

func TestLightPattern(t *testing.T) {
	material := NewMaterial(
		MaterialPattern(NewPattern()),
		Ambient(1),
		Diffuse(0),
		Specular(0),
	)

	eyeVector := New3Vector(0, 0, -1)
	normalVector := New3Vector(0, 0, -1)
	light := NewPointLight(NewColor(1, 1, 1), NewPoint(0, 0, -10))

	if Lighting(material, light, NewSphere(), NewPoint(0.9, 0, 0), eyeVector, normalVector, false).NotEq(Black()) {
		t.Errorf("Lighting Pattern not working")
	}

	if Lighting(material, light, NewSphere(), NewPoint(1.1, 0, 0), eyeVector, normalVector, false).NotEq(White()) {
		t.Errorf("Lighting Pattern not working")
	}
}
