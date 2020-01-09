package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"

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

func PrintArrayOfIntArray(arrays [][]int)  {
	for i, array := range arrays{
		fmt.Print(i)
		fmt.Print(" | ")
		for _, value := range array {
			fmt.Print(value)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func PrintArrayOfArray(arrays [][]interface{})  {
	for i, array := range arrays{
		fmt.Print(i)
		fmt.Print(" | ")
		for _, value := range array {
			fmt.Print(value)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func ModPirme(prime int, v float64) float64 {
	for {
		if v += float64(prime); v > 0 {
			break
		}
	}
	return float64(int(v) % prime)
}

func ModPrimeArray(prime int, values []float64) []float64 {
	var primed []float64
	for _, v := range values {
		primed = append(primed, ModPirme(prime, v))
	}
	return primed
}

func ChangedFormatInt(decimalNum int, baseNum int) []float64 {
	tmpNum := decimalNum
	conversionNum := ""

	for digitNum := 1;; {
		remainder := tmpNum%baseNum
		if remainder >= 10 {
			conversionNum = string('A' + (remainder-10)) + conversionNum
		} else {
			conversionNum = string('0' + remainder) + conversionNum
		}
		tmpNum = tmpNum/baseNum
		if tmpNum == 0 {
			break
		}
		digitNum *= 10
	}

	return StringToFloat64(strings.Split(conversionNum, ""))

}
