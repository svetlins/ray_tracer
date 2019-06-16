package common

import (
	"testing"
)

func TestMatrixMultiply(t *testing.T) {
	matrix1 := NewMatrix4([]float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	})

	matrix2 := NewMatrix4([]float64{
		-2, 1, 2, 3,
		3, 2, 1, -1,
		4, 3, 6, 5,
		1, 2, 7, 8,
	})

	expected := NewMatrix4([]float64{
		20, 22, 50, 48,
		44, 54, 114, 108,
		40, 58, 110, 102,
		16, 26, 46, 42,
	})

	if !matrix1.Multiply(matrix2).Eq(expected) {
		t.Errorf("Matrix multiply not working")
	}

}

func TestMatrixVectorMultiply(t *testing.T) {
	vector := Vector{1, 2, 3, 1}
	matrix := NewMatrix4([]float64{
		1, 2, 3, 4,
		2, 4, 4, 2,
		8, 6, 4, 1,
		0, 0, 0, 1,
	})

	result := matrix.Transform(vector)

	if !result.Eq(Vector{18, 24, 33, 1}) {
		t.Errorf("Matrix vector multiply not working")
	}

	if !matrix.Multiply(NewIdentity4()).Eq(matrix) {
		t.Errorf("Matrix identity multiply not working")
	}
}

func TestMatrixTranspose(t *testing.T) {
	matrix := NewMatrix4([]float64{
		1, 2, 3, 4,
		2, 4, 4, 2,
		8, 6, 4, 1,
		0, 0, 0, 1,
	})

	expected := NewMatrix4([]float64{
		1, 2, 8, 0,
		2, 4, 6, 0,
		3, 4, 4, 0,
		4, 2, 1, 1,
	})

	if !matrix.Transpose().Eq(expected) {
		t.Errorf("Matrix transpose not working")
	}

	if matrix.es[2] != 3 {
		t.Errorf("Matrix transpose not working")
	}
}

func TestMatrixInverse(t *testing.T) {
	matrix := NewMatrix4([]float64{
		-5, 2, 6, -8,
		1, -5, 1, 8,
		7, 7, -6, -7,
		1, -3, 7, 4,
	})

	expected := NewMatrix4([]float64{
		0.21805, 0.45113, 0.24060, -0.04511,
		-0.80827, -1.45677, -0.44361, 0.52068,
		-0.07895, -0.22368, -0.05263, 0.19737,
		-0.52256, -0.81391, -0.30075, 0.30639,
	})

	if !matrix.Inverse().Eq(expected) {
		t.Errorf("Matrix inverse not working")
	}
}
