package vectors

import (
	. "github.com/riita10069/jc/util"
	. "github.com/riita10069/jc/vector"
	"gonum.org/v1/gonum/mat"
	"math"
)

type Vectors struct {
	Values []Vector
	Dense mat.Dense
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

func (v *Vectors) SetMatrix(dense *mat.Dense) Vectors {
	return Vectors{Values: v.Values, Dense: *dense}
}

func (v *Vectors) GenerationStandardization(n, k int, dense mat.Dense) *mat.Dense {
	var S = dense.Slice(0, k, 0, k)
	T := mat.NewDense(k, k, nil)
	if err := T.Inverse(S); err != nil{
		err.Error()
	}
	G_L := mat.NewDense(k, k, nil)
	G_L.Product(S, T)

	G_R := mat.NewDense(k, n - k, nil)
	print("before gr")
	var R = dense.Slice(0, k, k, n)
	G_R.Product(R, T)

	print("after gr")

	G_ := mat.NewDense(k, n, nil)
	G_.Augment(G_L, G_R)

	return G_
}


