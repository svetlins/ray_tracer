package algebra

import (
	"testing"
	"gonum.org/v1/gonum/mat"
)

var result Matrix
var resultGonum mat.Matrix

func BenchmarkTranspose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matrix := NewMatrix(
			1,2,3,4,
			5,6,7,8,
			1,2,3,4,
			5,6,7,8,
		)

		result = matrix.Transpose()
	}
}

func BenchmarkTransposeGonum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matrix := mat.NewDense(4, 4, []float64{
			1,2,3,4,
			5,6,7,8,
			1,2,3,4,
			5,6,7,8,
		})

		resultGonum = matrix.T()
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
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

		result = matrix1.Multiply(matrix2)
	}
}

func BenchmarkMultiplyGonum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := mat.NewDense(4, 4, nil)

		matrix1 := mat.NewDense(4, 4, []float64{
			1, 2, 3, 4,
			5, 6, 7, 8,
			9, 8, 7, 6,
			5, 4, 3, 2,
		})

		matrix2 := mat.NewDense(4, 4, []float64{
			-2, 1, 2, 3,
			3, 2, 1, -1,
			4, 3, 6, 5,
			1, 2, 7, 8,
		})

		result.Mul(matrix1, matrix2)

		resultGonum = result
	}
}

func BenchmarkInverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matrix := NewMatrix(
			-5, 2, 6, -8,
			1, -5, 1, 8,
			7, 7, -6, -7,
			1, -3, 7, 4,
		)

		result = matrix.Inverse()

	}
}

func BenchmarkInverseGonum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := mat.NewDense(4, 4, nil)

		matrix := mat.NewDense(4, 4,[]float64{
			-5, 2, 6, -8,
			1, -5, 1, 8,
			7, 7, -6, -7,
			1, -3, 7, 4,
		})

		result.Inverse(matrix)

		resultGonum = result
	}
}
