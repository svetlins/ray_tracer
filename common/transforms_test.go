package common

import (
	"math"
	"testing"
)

func TestTranslation(t *testing.T) {
	translate := Translation(5, -3, 2)
	point := NewPoint(-3, 4, 5)
	vector := New3Vector(-3, 4, 5)

	if !translate.Transform(point).Eq(NewPoint(2, 1, 7)) {
		t.Errorf("Translate not working")
	}

	if !translate.Inverse().Transform(point).Eq(NewPoint(-8, 7, 3)) {
		t.Errorf("Translate not working")
	}

	if translate.Transform(vector).NotEq(vector) {
		t.Errorf("Translate not working")
	}
}

func TestScaling(t *testing.T) {
	scale := Scaling(2, 3, 4)
	point := NewPoint(-4, 6, 8)
	vector := New3Vector(-4, 6, 8)

	if scale.Transform(point).NotEq(NewPoint(-8, 18, 32)) {
		t.Errorf("Translate not working")
	}

	if scale.Inverse().Transform(vector).NotEq(New3Vector(-2, 2, 2)) {
		t.Errorf("Translate not working")
	}

	if scale.Transform(vector).NotEq(New3Vector(-8, 18, 32)) {
		t.Errorf("Translate not working")
	}

	scale = Scaling(-1, 1, 1)
	point = NewPoint(2, 3, 4)

	if scale.Transform(point).NotEq(NewPoint(-2, 3, 4)) {
		t.Errorf("Translate not working")
	}
}

func TestRotationX(t *testing.T) {
	halfQuarter := RotationX(math.Pi / 4)
	fullQuarter := RotationX(math.Pi / 2)

	point := NewPoint(0, 1, 0)

	if !halfQuarter.Transform(point).Eq(NewPoint(0, math.Sqrt(2)/2, math.Sqrt(2)/2)) {
		t.Errorf("RotateX not working")
	}

	if !halfQuarter.Inverse().Transform(point).Eq(NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)) {
		t.Errorf("RotateX not working")
	}

	if !fullQuarter.Transform(point).Eq(NewPoint(0, 0, 1)) {
		t.Errorf("RotateX not working")
	}
}

func TestRotationY(t *testing.T) {
	halfQuarter := RotationY(math.Pi / 4)
	fullQuarter := RotationY(math.Pi / 2)

	point := NewPoint(0, 0, 1)

	if !halfQuarter.Transform(point).Eq(NewPoint(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)) {
		t.Errorf("RotateY not working")
	}

	if !halfQuarter.Inverse().Transform(point).Eq(NewPoint(-math.Sqrt(2)/2, 0, math.Sqrt(2)/2)) {
		t.Errorf("RotateY not working")
	}

	if !fullQuarter.Transform(point).Eq(NewPoint(1, 0, 0)) {
		t.Errorf("RotateY not working")
	}
}

func TestRotationZ(t *testing.T) {
	halfQuarter := RotationZ(math.Pi / 4)
	fullQuarter := RotationZ(math.Pi / 2)

	point := NewPoint(0, 1, 0)

	if !halfQuarter.Transform(point).Eq(NewPoint(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)) {
		t.Errorf("RotateZ not working")
	}

	if !halfQuarter.Inverse().Transform(point).Eq(NewPoint(math.Sqrt(2)/2, math.Sqrt(2)/2, 0)) {
		t.Errorf("RotateZ not working")
	}

	if !fullQuarter.Transform(point).Eq(NewPoint(-1, 0, 0)) {
		t.Errorf("RotateZ not working")
	}
}

// func TestShearing(t *testing.T) {
// 	point := NewPoint(2, 3, 4)

// 	if !Shearing(1, 0, 0, 0, 0, 0).Multiply(point.ToMatrix()).ToVector().Eq(NewPoint(5, 3, 4)) {
// 		t.Errorf("Shearing not working")
// 	}

// 	if !Shearing(0, 1, 0, 0, 0, 0).Multiply(point.ToMatrix()).ToVector().Eq(NewPoint(6, 3, 4)) {
// 		t.Errorf("Shearing not working")
// 	}

// 	if !Shearing(0, 0, 1, 0, 0, 0).Multiply(point.ToMatrix()).ToVector().Eq(NewPoint(2, 5, 4)) {
// 		t.Errorf("Shearing not working")
// 	}

// 	if !Shearing(0, 0, 0, 1, 0, 0).Multiply(point.ToMatrix()).ToVector().Eq(NewPoint(2, 7, 4)) {
// 		t.Errorf("Shearing not working")
// 	}

// 	if !Shearing(0, 0, 0, 0, 1, 0).Multiply(point.ToMatrix()).ToVector().Eq(NewPoint(2, 3, 6)) {
// 		t.Errorf("Shearing not working")
// 	}

// 	if !Shearing(0, 0, 0, 0, 0, 1).Multiply(point.ToMatrix()).ToVector().Eq(NewPoint(2, 3, 7)) {
// 		t.Errorf("Shearing not working")
// 	}
// }

// func TestChaining(t *testing.T) {
// 	point := NewPoint(1, 0, 1)

// 	rotation := RotationX(math.Pi / 2)
// 	scaling := Scaling(5, 5, 5)
// 	translation := Translation(10, 5, 7)

// 	chained := translation.Multiply(scaling.Multiply(rotation))

// 	if !chained.Multiply(point.ToMatrix()).ToVector().Eq(NewPoint(15, 0, 7)) {
// 		t.Errorf("Chained transform not working")
// 	}
// }

// func TestViewTransform1(t *testing.T) {
// 	viewTranform := ViewTransform(
// 		NewPoint(0, 0, 0),
// 		NewPoint(0, 0, -1),
// 		New3Vector(0, 1, 0),
// 	)

// 	if viewTranform.NotEq(NewIdentity4()) {
// 		t.Errorf("ViewTransform not working")
// 	}
// }

// func TestViewTransform2(t *testing.T) {
// 	viewTranform := ViewTransform(
// 		NewPoint(0, 0, 0),
// 		NewPoint(0, 0, 1),
// 		New3Vector(0, 1, 0),
// 	)

// 	if viewTranform.NotEq(Scaling(-1, 1, -1)) {
// 		t.Errorf("ViewTransform not working")
// 	}
// }

// func TestViewTransform3(t *testing.T) {
// 	viewTranform := ViewTransform(
// 		NewPoint(0, 0, 8),
// 		NewPoint(0, 0, 0),
// 		New3Vector(0, 1, 0),
// 	)

// 	if viewTranform.NotEq(Translation(0, 0, -8)) {
// 		t.Errorf("ViewTransform not working")
// 	}
// }

// func TestViewTransform4(t *testing.T) {
// 	viewTransform := ViewTransform(
// 		NewPoint(1, 3, 2),
// 		NewPoint(4, -2, 8),
// 		New3Vector(1, 1, 0),
// 	)

// 	expected := NewMatrix4([]float64{
// 		-0.50709, 0.50709, 0.67612, -2.36643,
// 		0.76772, 0.60609, 0.12122, -2.82843,
// 		-0.35857, 0.59761, -0.71714, 0.00000,
// 		0.00000, 0.00000, 0.00000, 1.00000,
// 	})

// 	if viewTransform.NotEq(expected) {
// 		t.Errorf("ViewTransform not working")
// 	}
// }

// func TestMultipleTransforms(t *testing.T) {
// 	transforms1 := Translation(1,0,0).Multiply(RotationY(1)).Multiply(RotationZ(2))
// 	transforms2 := Transform(
// 		RotationZ(2),
// 		RotationY(1),
// 		Translation(1,0,0),
// 	)

// 	if transforms1.NotEq(transforms2) {
// 		t.Errorf("Multiple transforms not working")
// 	}
// }
