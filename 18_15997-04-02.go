package main

import (
	"fmt"
	"github.com/riita10069/jc/util"
	"gonum.org/v1/gonum/mat"
	"strconv"
	"strings"
)

func main()  {
	var p int
	fmt.Println("INPUT:")
	fmt.Print("p=")
	fmt.Scan(&p)

	var n int
	fmt.Print("n=")
	fmt.Scan(&n)

	fmt.Println("半角スペースを入れずに一行づつ123のように入力してください。二桁の数字は受け取れません。ごめんなさい")
	fmt.Println("A=")
	C := make([][]string, n)
	for i := range make([]int, n, n) {
		var c string
		fmt.Scan(&c)
		C[i] = strings.Split(c, "")
	}

	A := make([]float64, n * n)
	for i, x := range C {
		for j, y := range x {
			if s, err := strconv.ParseFloat(y, 64); err == nil {
				A[n * i + j] = s
			}else if err != nil {
				err.Error()
			}
		}
	}

	ADense := mat.NewDense(n, n, A)

	var b = []float64{}
	fmt.Println("b=")
	for range make([]int, n) {
		var z float64
		fmt.Scan(&z)
		b = append(b, z)
	}

	fmt.Println(b)
	B := mat.NewDense(n, 1, b)

	x := mat.NewDense(n, 1, nil)

	AInverse := mat.NewDense(n, n, nil)
	AInverse.Inverse(ADense)
	determinant := mat.Det(ADense)
	x.Product(AInverse, B)

	x.Scale(determinant, x)
	for e := range make([]int, n) {
		x.Set(e, 0, util.ModPirme(p, x.At(e, 0)))
	}

	util.MatPrint(x)
}
