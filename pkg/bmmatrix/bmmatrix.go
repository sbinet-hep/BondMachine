package bmmatrix

import "fmt"

// The BmMatrix struct holds a 2d matrix of floating point numbers

type BmMatrixSquareReal struct {
	N    int
	Data [][]float32
}

type Complex32 struct {
	Real float32
	Imag float32
}

func Complex32Add(a, b Complex32) Complex32 {
	return Complex32{a.Real + b.Real, a.Imag + b.Imag}
}

func Complex32Mul(a, b Complex32) Complex32 {
	return Complex32{a.Real*b.Real - a.Imag*b.Imag, a.Real*b.Imag + a.Imag*b.Real}
}

type BmMatrixSquareComplex struct {
	N    int
	Data [][]Complex32
}

func (m *BmMatrixSquareReal) String() string {
	str := ""
	for i := 0; i < m.N; i++ {
		substr := ""
		for j := 0; j < m.N; j++ {
			substr += fmt.Sprintf("| %f ", m.Data[i][j])
		}
		sep := ""
		for j := 0; j < len(substr)+1; j++ {
			sep += "-"
		}
		str += sep + "\n"
		str += substr + "|\n"
		if i == m.N-1 {
			str += sep + "\n"
		}
	}
	return str
}

func (m *BmMatrixSquareComplex) String() string {
	str := ""
	for i := 0; i < m.N; i++ {
		substr := ""
		for j := 0; j < m.N; j++ {
			substr += fmt.Sprintf("| %f + %fi ", m.Data[i][j].Real, m.Data[i][j].Imag)
		}
		sep := ""
		for j := 0; j < len(substr)+1; j++ {
			sep += "-"
		}
		str += sep + "\n"
		str += substr + "|\n"
		if i == m.N-1 {
			str += sep + "\n"
		}
	}

	return str
}

// NewBmMatrixSquare creates a new BmMatrixSquared with the given size
func NewBmMatrixSquareReal(n int) *BmMatrixSquareReal {
	data := make([][]float32, n)
	for i := range data {
		data[i] = make([]float32, n)
	}
	return &BmMatrixSquareReal{n, data}
}

// NewBmMatrixSquareComplex creates a new BmMatrixSquared with the given size
func NewBmMatrixSquareComplex(n int) *BmMatrixSquareComplex {
	data := make([][]Complex32, n)
	for i := range data {
		data[i] = make([]Complex32, n)
	}
	return &BmMatrixSquareComplex{n, data}
}

func TensorProductReal(a, b *BmMatrixSquareReal) *BmMatrixSquareReal {
	n := a.N * b.N
	c := NewBmMatrixSquareReal(n)
	for iA := 0; iA < a.N; iA++ {
		for jA := 0; jA < a.N; jA++ {
			for iB := 0; iB < b.N; iB++ {
				for jB := 0; jB < b.N; jB++ {
					c.Data[iA*b.N+iB][jA*b.N+jB] = a.Data[iA][jA] * b.Data[iB][jB]
				}
			}
		}
	}
	return c
}

func TensorProductComplex(a, b *BmMatrixSquareComplex) *BmMatrixSquareComplex {
	n := a.N * b.N
	c := NewBmMatrixSquareComplex(n)
	for iA := 0; iA < a.N; iA++ {
		for jA := 0; jA < a.N; jA++ {
			for iB := 0; iB < b.N; iB++ {
				for jB := 0; jB < b.N; jB++ {
					c.Data[iA*b.N+iB][jA*b.N+jB] = Complex32Mul(a.Data[iA][jA], b.Data[iB][jB])
				}
			}
		}
	}
	return c
}

func SwapRowsColsReal(a *BmMatrixSquareReal, x, y int) *BmMatrixSquareReal {
	b := NewBmMatrixSquareReal(a.N)
	for i := 0; i < a.N; i++ {
		for j := 0; j < a.N; j++ {
			if i == x {
				b.Data[i][j] = a.Data[y][j]
			} else if i == y {
				b.Data[i][j] = a.Data[x][j]
			} else {
				b.Data[i][j] = a.Data[i][j]
			}
		}
	}

	for i := 0; i < a.N; i++ {
		for j := 0; j < a.N; j++ {
			if j == x {
				b.Data[i][j] = a.Data[i][y]
			} else if j == y {
				b.Data[i][j] = a.Data[i][x]
			} else {
				b.Data[i][j] = a.Data[i][j]
			}
		}
	}

	return b
}

func SwapRowsColsComplex(a *BmMatrixSquareComplex, x, y int) *BmMatrixSquareComplex {
	b := NewBmMatrixSquareComplex(a.N)
	for i := 0; i < a.N; i++ {
		for j := 0; j < a.N; j++ {
			if i == x {
				b.Data[i][j] = a.Data[y][j]
			} else if i == y {
				b.Data[i][j] = a.Data[x][j]
			} else {
				b.Data[i][j] = a.Data[i][j]
			}
		}
	}

	for i := 0; i < a.N; i++ {
		for j := 0; j < a.N; j++ {
			if j == x {
				b.Data[i][j] = a.Data[i][y]
			} else if j == y {
				b.Data[i][j] = a.Data[i][x]
			} else {
				b.Data[i][j] = a.Data[i][j]
			}
		}
	}

	return b
}
