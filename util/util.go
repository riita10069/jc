package util

import (
	"fmt"
	"math"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

func MatPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func Max(a []float64) (float64, int) {
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

func Min(a []float64) (float64, int) {
	min := math.MaxFloat64
	num := 0
	for i, v := range a {
		if v == 0 {
			continue
		} else if v < min {
			min = v
			num = i
		}
	}
	return min, num
}

func StringToFloat64(i []string) []float64 {
	f := make([]float64, len(i))
	f_prime := make([]int, len(i))
	for n := range i {
		f_prime[n], _ = strconv.Atoi(i[n])
		f[n] = float64(f_prime[n])
	}
	return f
}
