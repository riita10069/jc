package vectors

import (
	. "github.com/riita10069/jc/vector"
)

type Vectors struct {
	Values []Vector
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

func (v *Vectors) HammingMatrixDecode(recieve Vector) Vector {
	tmp := []float64{}
	for _, vector := range v.Values {
		j := vector.HammingDistance(recieve)
		tmp = append(tmp, j)
	}
	_, num := min(tmp)
	return v.Values[num]
}

func max(a []float64) (float64, int) {
	max := a[0]
	num := 0
	for i, v := range a {
		if v < max {
			max = v
			num = i
		}
	}
	return max, num
}

func min(a []float64) (float64, int) {
	min := a[0]
	num := 0
	for i, v := range a {
		if v < min {
			min = v
			num = i
		}
	}
	return min, num
}

