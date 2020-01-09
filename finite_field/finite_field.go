package finite_field

import (
	"math"

	. "github.com/riita10069/jc/number"
	. "github.com/riita10069/jc/numerical"
	"gonum.org/v1/gonum/mat"
)

type FiniteNumberField struct {
	prime   int
	numbers []Number
}

func NewFiniteNumberField(prime int, numbers []Number) *FiniteNumberField {
	return &FiniteNumberField{prime: prime, numbers: numbers}
}

type IFiniteField interface {
	Size() [][]int
	AddTable() [][]int
	SubTable() [][]int
	MulTable() [][]int
	DivTable() [][]interface{}
}

func (f FiniteNumberField) Size() int {
	return f.prime
}

func (f FiniteNumberField) LinearEquation(A *mat.Dense, B *mat.Dense) *mat.Dense {
	AInverse := mat.NewDense(A.RowView(0).Len(), A.ColView(0).Len(), nil)
	AInverse.Inverse(A)
	C := mat.NewDense(AInverse.RowView(0).Len(), B.ColView(0).Len(), nil)
	C.Product(AInverse, B)
	return C
}

func (f FiniteNumberField) AddTable() [][]int {
	matrix := make([][]int, f.Size(), f.Size())
	for k := range make([]int, f.Size()){
		matrix[k] = make([]int, f.Size())
	}

	for i := range make([]int, f.Size(), f.Size()) {
		for j := range make([]int, f.Size(), f.Size()) {
			matrix[i][j] = f.numbers[i].Add(&f.numbers[j])
		}
	}



	return matrix
}

func (f FiniteNumberField) SubTable() [][]int {
	matrix := make([][]int, f.Size(), f.Size())
	for k := range make([]int, f.Size()){
		matrix[k] = make([]int, f.Size())
	}
	for i := range make([]int, f.Size(), f.Size()) {
		for j := range make([]int, f.Size(), f.Size()) {
			matrix[i][j] = f.numbers[i].Sub(&f.numbers[j])
			if matrix[i][j] < 0 {
				matrix[i][j] += f.prime
			}
		}
	}
	return matrix
}

func (f FiniteNumberField) MulTable() [][]int {
	matrix := make([][]int, f.Size(), f.Size())
	for k := range make([]int, f.Size()){
		matrix[k] = make([]int, f.Size())
	}
	for i := range make([]int, f.Size(), f.Size()) {
		for j := range make([]int, f.Size(), f.Size()) {
			matrix[i][j] = f.numbers[i].Mul(&f.numbers[j])
		}
	}
	return matrix
}

func (f FiniteNumberField) DivTable() [][]interface{} {
	matrix := make([][]interface{}, f.Size(), f.Size())
	for k := range make([]interface{}, f.Size()){
		matrix[k] = make([]interface{}, f.Size())
	}
	for i := range make([]int, f.Size(), f.Size()) {
		for j := range make([]int, f.Size(), f.Size()) {
			matrix[i][j], _ = f.numbers[i].Div(&f.numbers[j])
			if matrix[i][j] == -1 {
				matrix[i][j] = "-"
			}
		}
	}
	return matrix
}

type FiniteNumericalField struct {
	prime Poly
	polys []Poly
}

func NewFiniteNumericalField(prime Poly, polys []Poly) *FiniteNumericalField {
	return &FiniteNumericalField{prime: prime, polys: polys}
}

func (f FiniteNumericalField) Size() int {
	return int(math.Pow(2.0, float64(f.prime.Deg())))
}

func (f FiniteNumericalField) AddTable() [][]Poly {
	matrix := make([][]Poly, f.Size(), f.Size())
	for i := range make([]int, f.Size(), f.Size()) {
		for j := range make([]int, f.Size(), f.Size()) {
			matrix[i][j] = PolyAdd(f.polys[i], f.polys[j])
		}
	}
	return matrix
}

func (f FiniteNumericalField) SubTable() [][]Poly {
	matrix := make([][]Poly, f.Size(), f.Size())
	for i := range make([]int, f.Size(), f.Size()) {
		for j := range make([]int, f.Size(), f.Size()) {
			matrix[i][j] = PolySub(f.polys[i], f.polys[j])
		}
	}
	return matrix
}

func (f FiniteNumericalField) MulTable() [][]Poly {
	matrix := make([][]Poly, f.Size(), f.Size())
	for i := range make([]int, f.Size(), f.Size()) {
		for j := range make([]int, f.Size(), f.Size()) {
			matrix[i][j] = PolyMul(f.polys[i], f.polys[j])
		}
	}
	return matrix
}

func (f FiniteNumericalField) DivTable() [][]interface{} {
	matrix := make([][]interface{}, f.Size(), f.Size())
	for i := range make([]int, f.Size(), f.Size()) {
		for j := range make([]int, f.Size(), f.Size()) {
			matrix[i][j], _ = PolyDiv(f.polys[i], f.polys[j])
		}
	}
	return matrix
}
