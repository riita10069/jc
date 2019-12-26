package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	. "github.com/riita10069/jc/vector"
	. "github.com/riita10069/jc/vectors"
	. "github.com/riita10069/jc/util"
)

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func main() {
	vec1 := Vector{}
	vec2 := Vector{}
	vec1.Value = []float64{1, 0, 0, 1, 0}
	vec2.Value = []float64{0, 0, 0, 0, 0}

	vectors := Vectors{}
	vectors.Values = []Vector{vec1, vec2}
	size := vectors.Size()

	recieve := Vector{}
	recieve.Value = []float64{1, 1, 1, 0, 1}
	min_vector := vectors.HammingMatrixDecode(recieve)
	fmt.Print(min_vector.Value)
	A := mat.NewDense(size, size, vectors.HammingMatrix())
	fmt.Print(size)
	matPrint(A)
}
