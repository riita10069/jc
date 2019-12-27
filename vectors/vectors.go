package vectors

import (
	. "github.com/riita10069/jc/util"
	. "github.com/riita10069/jc/vector"
	"math"
)

type Vectors struct {
	Values []Vector
}

func NewVectors(values []Vector) *Vectors {
	return &Vectors{Values: values}
}

func (v *Vectors) HammingMatrix() []float64 {
	var hamming_mat []float64
	for _, vector := range v.Values {
		for _, other := range v.Values {
			hamming_mat = append(hamming_mat, vector.HammingDistance(other))
		}
	}
	return hamming_mat
}

func (v *Vectors) Size() int {
	return len(v.Values)
}

func (v *Vectors) HammingMatrixDecode(recieve Vector) (Vector, int) {
	var tmp []float64
	for _, vector := range v.Values {
		j := vector.HammingDistance(recieve)
		tmp = append(tmp, j)
	}
	_, num := Min(tmp)
	return v.Values[num], num
}

func (v *Vectors) DMin() float64 {
	i, _ := Min(v.HammingMatrix())
	return i
}

func (v *Vectors) Rate() float64 {
	return  math.Log2(float64(v.Size())) / float64(v.Values[0].GetLength())
}


