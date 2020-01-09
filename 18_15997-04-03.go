package main

import (
	"fmt"
	"github.com/riita10069/jc/numerical"
	"github.com/riita10069/jc/util"
	"strings"
)

func main(){
	var n int
	fmt.Println("INPUT:")
	fmt.Print("n=")
	fmt.Scan(&n)

	var m int
	fmt.Print("m=")
	fmt.Scan(&m)

	f := make([]float64, n)
	var c string
	fmt.Print("f=")
	fmt.Scan(&c)

	f = util.StringToFloat64(strings.Split(c, ""))
	F := numerical.NewPoly(append([]float64{}, f...))

	g := make([]float64, n)
	var d string
	fmt.Print("g=")
	fmt.Scan(&d)
	g = util.StringToFloat64(strings.Split(d, ""))
	G := numerical.NewPoly(append([]float64{}, g...))

	fmt.Println("OUTPUT:")
	fmt.Print("f+g= ")
	add := numerical.PolyAdd(F, G)
	fmt.Println(util.ModPrimeArray(2, add.Coeffs))

	F = numerical.NewPoly(append([]float64{}, f...))
	G = numerical.NewPoly(append([]float64{}, g...))
	fmt.Print("f-g= ")
	sub := numerical.PolySub(F, G)
	fmt.Println(util.ModPrimeArray(2, sub.Coeffs))

	F = numerical.NewPoly(append([]float64{}, f...))
	G = numerical.NewPoly(append([]float64{}, g...))
	fmt.Print("f*g= ")
	mul := numerical.PolyMul(F, G)
	fmt.Println(util.ModPrimeArray(2, mul.Coeffs))

	F = numerical.NewPoly(append([]float64{}, f...))
	G = numerical.NewPoly(append([]float64{}, g...))
	fmt.Println("f/g= ")
	q, r := numerical.PolyDiv(F, G)
	fmt.Print("商")
	fmt.Println(util.ModPrimeArray(2, q.Coeffs))
	fmt.Println("剰余")
	fmt.Println(util.ModPrimeArray(2, r.Coeffs))
}