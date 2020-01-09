//␣プログラミング言語：golang
//␣コンパイル方法：env go build -o [ビルド後のファイル名] [ファイル名]
//␣実行方法：./[ビルド後のファイル名]

package main

import (
	"fmt"
	"github.com/riita10069/jc/util"
	"github.com/riita10069/jc/vector"
	"github.com/riita10069/jc/vectors"
	"gonum.org/v1/gonum/mat"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Println("INPUT:")
	fmt.Print("n=")
	fmt.Scan(&n)

	var k int
	fmt.Print("k=")
	fmt.Scan(&k)

	fmt.Println("G=")
	C := make([][]string, k)
	for i := range make([]int, k, k) {
		var c string
		fmt.Scan(&c)
		C[i] = strings.Split(c, "")
	}

	G := make([]float64, n * k)
	for i, x := range C {
		for j, y := range x {
			if s, err := strconv.ParseFloat(y, 64); err == nil {
				G[n * i + j] = s
			}else if err != nil {
				err.Error()
			}
		}
	}

	GDense := mat.NewDense(k, n, G)

	cs := make([]vector.Vector, k)
	for i, char := range C {
		cs[i] = vector.NewVector(util.StringToFloat64(char))
	}

	Vec := vectors.NewVectors(cs)
	Vec.SetMatrix(GDense)

	fmt.Println("OUTPUT:")
	fmt.Print("M=")
	fmt.Println(n * k)

	fmt.Print("R=")
	rate := Vec.Rate()
	fmt.Println(rate)

	fmt.Print("d_min=")
	dMin := Vec.DMin()
	fmt.Println(dMin)

	fmt.Print("G'=")
	generationStandardization := Vec.GenerationStandardization(n, k, *GDense)
	util.MatPrint(generationStandardization)

}