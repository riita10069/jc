package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	. "github.com/jc/kadai/vector"
	. "github.com/jc/kadai/vectors"
)

func main() {
	vec1 := Vector{}
	vec2 := Vector{}
	vec1.Value = []float64{1, 0, 0, 1, 0}
	vec2.Value = []float64{0, 0, 0, 0, 0}

	vectors := Vectors{}
	vectors.Values = []Vector{vec1, vec2}
	size := vectors.size()

	A := mat.NewDense(size, size, vectors.hammingMatrix())
	fmt.Print(size)
	matPrint(A)
}
