package util

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
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
