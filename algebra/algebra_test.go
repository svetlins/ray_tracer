package algebra

import (
	"testing"
)

func TestNewVector(t *testing.T) {
	vector := NewVector(1,2,3)

	if vector.x != 1 || vector.y != 2 || vector.z != 3 || vector.w != 0 {
		t.Error()
	}
}

func TestNewPoint(t *testing.T) {
	vector := NewPoint(1,2,3)

	if vector.x != 1 || vector.y != 2 || vector.z != 3 || vector.w != 1 {
		t.Errorf("new vector")
	}
}

func TestVectorDot(t *testing.T) {
	v1 := NewVector(1,2,3)
	v2 := NewVector(3,2,1)

	if v1.Dot(v2) != 10 {
		t.Error()
	}
}

func TestNewMatrix(t *testing.T) {
	matrix := NewMatrix(
		1,2,3,4,
		5,6,7,8,
		1,2,3,4,
		5,6,7,8,
	)

	if matrix.NotEqual(matrix) {
		t.Error()
	}
}

func TestMatrixTranspose(t *testing.T) {
	matrix1 := NewMatrix(
		1,2,3,4,
		5,6,7,8,
		1,2,3,4,
		5,6,7,8,
	)

	expected := NewMatrix(
		1,5,1,5,
		2,6,2,6,
		3,7,3,7,
		4,8,4,8,
	)

	if matrix1.Transpose().NotEqual(expected) {
		t.Error()
	}
}

func TestMatrixMultiplyIdentity(t *testing.T) {
	matrix1 := NewMatrix(
		1,2,3,4,
		5,6,7,8,
		1,2,3,4,
		5,6,7,8,
	)

	matrix2 := NewMatrix(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)

	if matrix1.Multiply(matrix2).NotEqual(matrix1) {
		t.Error()
	}
}

func TestMatrixMultiply(t *testing.T) {
	matrix1 := NewMatrix(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	)

	matrix2 := NewMatrix(
		-2, 1, 2, 3,
		3, 2, 1, -1,
		4, 3, 6, 5,
		1, 2, 7, 8,
	)

	expected := NewMatrix(
		20, 22, 50, 48,
		44, 54, 114, 108,
		40, 58, 110, 102,
		16, 26, 46, 42,
	)

	if matrix1.Multiply(matrix2).NotEqual(expected) {
		t.Error()
	}
}

func TestMatrixTransform(t *testing.T) {
	vector := Vector{1, 2, 3, 1}

	matrix := NewMatrix(
		1, 2, 3, 4,
		2, 4, 4, 2,
		8, 6, 4, 1,
		0, 0, 0, 1,
	)

	result := matrix.Transform(vector)

	if result.NotEqual(Vector{18, 24, 33, 1}) {
		t.Error()
	}
}

func TestMatrixDeterminant2(t *testing.T) {
	matrix := NewMatrix(
		1, 5,
		-3, 2,
	)

	if matrix.Determinant() != 17 {
		t.Error()
	}
}

func TestMatrixSubmatrix1(t *testing.T) {
	matrix := NewMatrix(
		1, 5, 0,
		-3, 2, 7,
		0, 6, 3,
	)

	expected := NewMatrix(
		-3, 2,
		0, 6,
	)

	if matrix.Submatrix(0, 2).NotEqual(expected) {
		t.Error()
	}

}

func TestMatrixSubmatrix2(t *testing.T) {
	matrix := NewMatrix(
		-6, 1, 1, 6,
		-8, 5, 8, 6,
		1, 0, 8, 2,
		-7, 1, -1, 1,
	)

	expected := NewMatrix(
		-6, 1, 6,
		-8, 8, 6,
		-7, -1, 1,
	)

	if matrix.Submatrix(2, 1).NotEqual(expected) {
		t.Error()
	}

}

func TestMatrixMinor(t *testing.T) {
	matrix := NewMatrix(
		3, 5, 0,
		2, -1, -7,
		6, -1, 5,
	)

	if matrix.Minor(1, 0) != 25 {
		t.Error()
	}
}

func TestMatrixCofactor(t *testing.T) {
	matrix := NewMatrix(
		3, 5, 0,
		2, -1, -7,
		6, -1, 5,
	)

	if matrix.Minor(1, 0) != 25 || matrix.Cofactor(1, 0) != -25 {
		t.Error()
	}

	if matrix.Minor(0, 0) != -12 || matrix.Cofactor(0, 0) != -12 {
		t.Error()
	}
}

func TestMatrixDeterminant3(t *testing.T) {
	matrix := NewMatrix(
		1, 2, 6,
		-5, 8, -4,
		2, 6, 4,
	)

	if matrix.Determinant() != -196 {
		t.Error()
	}
}

func TestMatrixDeterminant4(t *testing.T) {
	matrix := NewMatrix(
		-2, -8, 3, 5,
		-3, 1, 7, 3,
		1, 2, -9, 6,
		-6, 7, 7, -9,
	)

	if matrix.Determinant() != -4071 {
		t.Error()
	}
}

func TestMatrixInverse(t *testing.T) {
	matrix := NewMatrix(
		-5, 2, 6, -8,
		1, -5, 1, 8,
		7, 7, -6, -7,
		1, -3, 7, 4,
	)

	expected := NewMatrix(
		0.21805, 0.45113, 0.24060, -0.04511,
		-0.80827, -1.45677, -0.44361, 0.52068,
		-0.07895, -0.22368, -0.05263, 0.19737,
		-0.52256, -0.81391, -0.30075, 0.30639,
	)

	if matrix.Inverse().NotEqual(expected) {
		t.Error()
	}
}
