package vectors

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	. "github.com/jc/kadai/vector"
)

type Vectors struct {
	Values []Vector
}

func (v *Vectors) hammingMatrix() []float64 {
	var hamming_mat []float64
	for _, vector := range v.Values {
		for _, other := range v.Values {
			hamming_mat = append(hamming_mat, vector.hammingDistance(other))
		}
	}
	return hamming_mat
}

func (v *Vectors) size() int {
	return len(v.Values)
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}