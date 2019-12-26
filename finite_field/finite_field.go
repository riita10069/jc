package finite_field

import (
	"errors"
	. "github.com/riita10069/jc/numerical"
	. "github.com/riita10069/jc/number"
)



type FiniteNumberField struct {
	prime    int
	numbers  []Number
}

type IFiniteField interface {
	AddTable() [][]int
	SubTable()
	MulTable()
	DivTable()
}


func (f FiniteNumberField) AddTable() [][]int {
	matrix := make([][]int, f.Size(), f.Size())
	for i := range make([]int, f.Size(), f.Size()) {
		for j := range make([]int, f.Size(), f.Size()) {
			matrix[i][j] = f.numbers[i].Add(&f.numbers[j])
		}
	}
	return matrix
}

func (f FiniteNumberField) SubTable()[][]int {
	matrix := make([][]int, f.Size(), f.Size())
	for i := range make([]int, f.Size(), f.Size()) {
		for j := range make([]int, f.Size(), f.Size()) {
			matrix[i][j] = f.numbers[i].Sub(&f.numbers[j])
		}
	}
	return matrix
}

func (f FiniteNumberField) MulTable() [][]int{
	matrix := make([][]int, f.Size(), f.Size())
	for i := range make([]int, f.Size(), f.Size()) {
		for j := range make([]int, f.Size(), f.Size()) {
			matrix[i][j] = f.numbers[i].Mul(&f.numbers[j])
		}
	}
	return matrix
}

func (f FiniteNumberField) DivTable()[][]interface{} {
	matrix := make([][]interface{}, f.Size(), f.Size())
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
	prime int
	polys []Poly
}


func main()  {
	a := FiniteNumberField{}
	a.MulTable()
}

