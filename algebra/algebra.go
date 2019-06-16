package algebra

import (
	"math"
)

const Epsilon = 0.1

type Vector struct {
	x,y,z,w float64
}

type Matrix struct {
	es []float64
	rows, cols int
}

func NewVector(x,y,z float64) Vector {
	return Vector{x,y,z,0}
}

func NewPoint(x,y,z float64) Vector {
	return Vector {x,y,z,1}
}

func NewMatrix(data ...float64) Matrix {
	size := int(math.Sqrt(float64(len(data))))

	return Matrix{
		es: data,
		rows: size,
		cols: size,
	}
}

func floatEqual(a float64, b float64) bool {
	return math.Abs(a-b) < Epsilon
}

func (v Vector)Equal(other Vector) bool {
	return floatEqual(v.x, other.x) &&
		floatEqual(v.y, other.y) &&
		floatEqual(v.z, other.z) &&
		floatEqual(v.w, other.w)
}

func (v Vector)NotEqual(other Vector) bool {
	return !(v.Equal(other))
}

func (v Vector)Dot(other Vector) float64 {
	return v.x*other.x + v.y*other.y + v.z*other.z + v.w*other.w
}

func (m Matrix)Equal(other Matrix) bool {
	if m.rows != other.rows || m.cols != other.cols {
		return false
	}

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			if !floatEqual(m.es[i * m.cols + j], other.es[i * m.cols + j]) {
				return false
			}
		}
	}

	return true
}

func (m Matrix)NotEqual(other Matrix) bool {
	return !(m.Equal(other))
}

func (m Matrix)Transpose() Matrix {
	buffer := make([]float64, m.rows * m.cols)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			buffer[j * m.cols + i] = m.es[i * m.cols + j]
		}
	}

	return NewMatrix(buffer...)
}

func (m Matrix)Multiply(other Matrix) Matrix {
	buffer := make([]float64, m.rows * other.cols)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < other.cols; j++ {
			elem := 0.0

			for k := 0; k < m.cols; k++ {
				elem = elem + m.es[i * m.cols + k] * other.es[k * other.cols + j]
			}

			buffer[i * other.cols + j] = elem
		}
	}

	return NewMatrix(buffer...)
}

func (m Matrix)Transform(v Vector) Vector {
	x := m.es[0] * v.x +
		m.es[1] * v.y +
		m.es[2] * v.z +
		m.es[3] * v.w

	y := m.es[4] * v.x +
		m.es[5] * v.y +
		m.es[6] * v.z +
		m.es[7] * v.w

	z := m.es[8] * v.x +
		m.es[9] * v.y +
		m.es[10] * v.z +
		m.es[11] * v.w

	w := m.es[12] * v.x +
		m.es[13] * v.y +
		m.es[14] * v.z +
		m.es[15] * v.w

	return Vector{x,y,z,w}
}

func (m Matrix)Determinant() float64 {
	if (m.rows == 2) && (m.cols == 2) {
		return m.es[0] * m.es[3] - m.es[1] * m.es[2]
	} else if m.rows == m.cols {
		determinant := 0.0

		for i:= 0; i < m.cols; i++ {
			determinant = determinant +
				m.es[i] * m.Cofactor(0, i)
		}

		return determinant
	}

	panic("Can't compute determinant of non-square matrix")
}

func (m Matrix)Submatrix(row, column int) Matrix {
	buffer := make([]float64, (m.rows - 1) * (m.cols - 1))

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			var r, c int

			if i != row && j != column {
				if i < row {
					r = i
				} else if i > row {
					r = i -1
				}

				if j < column {
					c = j
				} else if j > column {
					c = j - 1
				}

				buffer[r * (m.cols - 1) + c] = m.es[i * m.cols + j]
			}
		}
	}

	return NewMatrix(buffer...)
}

func (m Matrix)Minor(row, column int) float64 {
	return m.Submatrix(row, column).Determinant()
}

func (m Matrix)Cofactor(row, column int) float64 {
	if (row + column) % 2 == 0 {
		return m.Minor(row, column)
	} else {
		return -m.Minor(row, column)
	}
}

func (m Matrix)Inverse() Matrix {
	buffer := make([]float64, m.rows * m.cols)

	determinant := m.Determinant()

	if floatEqual(determinant, 0) {
		panic("Can't invert non-invertible matrix")
	}

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			buffer[i * m.cols + j] = m.Cofactor(j, i) / determinant
		}
	}

	return NewMatrix(buffer...)
}
