package main

import (
	"fmt"
	"github.com/riita10069/jc/numerical"
	"github.com/riita10069/jc/util"
	"math"
	"strings"
)

func make_row(m int) [][]float64 {
	var arr = make([][]float64, int(math.Pow(2, float64(m))))
	for i, _ := range arr {
		arr[i] = util.ChangedFormatInt(i, 2)
		for len(arr[i]) != m {
			arr[i] = append([]float64{0}, arr[i]...)
		}
	}
	return arr
}

func make_table_head(m int) {
	fmt.Print("　　　　")
	fmt.Println(make_row(m))
	fmt.Println("------------------------------------------------------------------------------------------------------------------------")
}

func make_add_table(m int)  {
	make_table_head(m)
	for _, v := range make_row(m) {
		fmt.Print(v)
		fmt.Print(" |")
		for _, w := range make_row(m) {
			fmt.Print(util.ModPrimeArray(2, numerical.PolyAdd(numerical.NewPoly(append([]float64{}, v...)), numerical.NewPoly(append([]float64{}, w...))).Coeffs))
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func make_sub_table(m int)  {
	make_table_head(m)
	for _, v := range make_row(m) {
		fmt.Print(v)
		fmt.Print(" |")
		for _, w := range make_row(m) {
			fmt.Print(util.ModPrimeArray(2, numerical.PolyAdd(numerical.NewPoly(append([]float64{}, v...)), numerical.NewPoly(append([]float64{}, w...))).Coeffs))
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func make_mul_table(m int, F numerical.Poly)  {
	make_table_head(m)
	for _, v := range make_row(m) {
		fmt.Print(v)
		fmt.Print(" |")
		for _, w := range make_row(m) {
			if -1 == numerical.NewPoly(append([]float64{}, w...)).Deg2() {
				fmt.Print(" [0 0 0] ")
			} else {
				q, _ := numerical.PolyDiv(numerical.PolyMul(numerical.NewPoly(append([]float64{}, v...)), numerical.NewPoly(append([]float64{}, w...))), F)
				r := make([]float64, 0)
				r = append([]float64{}, q.Coeffs...)
				for len(r) < m {
					r = append([]float64{0}, r...)
				}
				fmt.Print(util.ModPrimeArray(2, r))
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
}

func make_div_table(m int)  {
	make_table_head(m)
	for _, v := range make_row(m) {
		fmt.Print(v)
		fmt.Print(" |")
		for _, w := range make_row(m) {
			if -1 == numerical.NewPoly(append([]float64{}, w...)).Deg2() {
				fmt.Print(" -----  ")
			} else {
				q, _:= numerical.PolyDiv(numerical.NewPoly(append([]float64{}, v...)), numerical.NewPoly(append([]float64{}, w...)))
				r := make([]float64, 0)
				r = append([]float64{}, q.Coeffs...)
				for len(r) < m {
					r = append([]float64{0}, r...)
				}
				fmt.Print(util.ModPrimeArray(2, r))
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("INPUT:")
	var m int
	fmt.Print("m=")
	fmt.Scan(&m)

	f := make([]float64, m + 1)
	var c string
	fmt.Print("f=")
	fmt.Scan(&c)

	f = util.StringToFloat64(strings.Split(c, ""))
	F := numerical.NewPoly(append([]float64{}, f...))

	fmt.Println(F.Str())
	fmt.Println("OUTPUT:")
	fmt.Println("ADD")
	make_add_table(m)

	fmt.Println("SUB")
	make_sub_table(m)

	fmt.Println("MUL")
	make_mul_table(m, numerical.NewPoly(append([]float64{}, f...)))

	fmt.Println("DIV")
	make_div_table(m)
}
