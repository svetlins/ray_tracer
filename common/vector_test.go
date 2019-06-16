package common

import (
	"math"
	"testing"
)

func TestIsPoint(t *testing.T) {
	if (Vector{4.3, -4.2, 3.1, 1.0}).IsPoint() != true {
		t.Errorf("IsPoint does not recognize point")
	}

	if (Vector{4.3, -4.2, 3.1, 0.0}).IsPoint() != false {
		t.Errorf("IsPoint mistakenly recognizes vector")
	}
}

func TestCreatePoint(t *testing.T) {
	if NewPoint(1, 2, 3).IsPoint() != true {
		t.Errorf("CreatePoint did not create point")
	}
}

func TestCreateColor(t *testing.T) {
	if NewColor(0.1, 0.2, 0.3) != (Vector{0.1, 0.2, 0.3, 0}) {
		t.Errorf("CreatePoint did not create point")
	}
}

func TestNew3Vector(t *testing.T) {
	if New3Vector(1, 2, 3).IsPoint() == true {
		t.Errorf("NewVector did not create vector")
	}
}

func TestVectorEquality(t *testing.T) {
	if !(Vector{1, 2, 3, 0}).Eq(Vector{1, 2, 3, 0}) {
		t.Errorf("Vector equality not working")
	}
}

func TestVectorInequality(t *testing.T) {
	if !(Vector{1, 2, 3, 0}).NotEq(Vector{1, 2, 3, 4}) {
		t.Errorf("Vector equality not working")
	}
}

func TestVectorAdd(t *testing.T) {
	a := Vector{1, 2, 3, 0}
	b := Vector{10, 20, 30, 1}

	if !a.Add(b).Eq(Vector{11, 22, 33, 1}) {
		t.Errorf("VectorAdd not working")
	}
}

func TestVectorSubtract(t *testing.T) {
	a := Vector{10, 20, 30, 1}
	b := Vector{1, 2, 3, 0}

	if !a.Subtract(b).Eq(Vector{9, 18, 27, 1}) {
		t.Errorf("VectorSubtact not working")
	}
}

func TestVectorNegate(t *testing.T) {
	if !New3Vector(1, 2, 3).Negate().Eq(NewColor(-1, -2, -3)) {
		t.Errorf("VectorNegate not working")
	}
}

func TestVectorScale(t *testing.T) {
	if !New3Vector(1, 2, 3).Scale(10).Eq(New3Vector(10, 20, 30)) {
		t.Errorf("VectorScale not working")
	}
}

func TestVectorDivide(t *testing.T) {
	if !New3Vector(1, 2, 3).Divide(2).Eq(New3Vector(0.5, 1, 1.5)) {
		t.Errorf("VectorDivide not working")
	}
}

func TestVectorMultiply(t *testing.T) {
	if !New3Vector(1, 0.2, 0.4).Multiply(New3Vector(0.9, 1, 0.1)).Eq(New3Vector(0.9, 0.2, 0.04)) {
		t.Errorf("VectorDivide not working")
	}
}

func TestVectorMagnitude(t *testing.T) {
	if New3Vector(1, 0, 0).Magnitude() != 1 {
		t.Errorf("VectorMagnitude not working")
	}

	if New3Vector(1, 2, 3).Magnitude() != math.Sqrt(14) {
		t.Errorf("VectorMagnitude not working")
	}

	if New3Vector(-1, -2, -3).Magnitude() != math.Sqrt(14) {
		t.Errorf("VectorMagnitude not working")
	}
}

func TestVectorNormalize(t *testing.T) {
	if !New3Vector(4, 0, 0).Normalize().Eq(New3Vector(1, 0, 0)) {
		t.Errorf("VectorNormalize not working")
	}

	normalizedVector := New3Vector(1, 2, 3).Normalize()
	expectedVector := New3Vector(
		1.0/math.Sqrt(14),
		2.0/math.Sqrt(14),
		3.0/math.Sqrt(14),
	)

	if !normalizedVector.Eq(expectedVector) {
		t.Errorf("VectorNormalize not working")
	}

	if New3Vector(1, 2, 3).Normalize().Magnitude() != 1 {
		t.Errorf("VectorNormalize not working")
	}
}

func TestVectorDot(t *testing.T) {
	if New3Vector(1, 2, 3).Dot(New3Vector(2, 3, 4)) != 20 {
		t.Errorf("VectorDot not working")
	}
}

func TestVectorCross(t *testing.T) {
	if !New3Vector(1, 2, 3).Cross(New3Vector(2, 3, 4)).Eq(New3Vector(-1, 2, -1)) {
		t.Errorf("VectorCross not working")
	}

	if !New3Vector(2, 3, 4).Cross(New3Vector(1, 2, 3)).Eq(New3Vector(1, -2, 1)) {
		t.Errorf("VectorCross not working")
	}
}

func TestVectorReflect1(t *testing.T) {
	vector := New3Vector(1, -1, 0)
	normal := New3Vector(0, 1, 0)

	if !vector.Reflect(normal).Eq(New3Vector(1, 1, 0)) {
		t.Errorf("Vector Reflect not working")
	}
}

func TestVectorReflect2(t *testing.T) {
	vector := New3Vector(0, -1, 0)
	normal := New3Vector(math.Sqrt(2)/2, math.Sqrt(2)/2, 0)

	if !vector.Reflect(normal).Eq(New3Vector(1, 0, 0)) {
		t.Errorf("Vector Reflect not working %+v", vector.Reflect(normal))
	}
}
