//␣プログラミング言語：golang
//␣コンパイル方法：env go build -o [ビルド後のファイル名] [ファイル名]
//␣実行方法：./[ビルド後のファイル名]

package main

import (
	"fmt"
	"github.com/riita10069/jc/util"
	"github.com/riita10069/jc/vector"
	"github.com/riita10069/jc/vectors"
	"strings"
)

func main() {
	var n int
	fmt.Println("INPUT:")
	fmt.Print("n=")
	fmt.Scan(&n)

	var M int
	fmt.Print("M=")
	fmt.Scan(&M)

	fmt.Println("C=")
	C := make([][]string, M)
	for i := range make([]int, M, M) {
		var c string
		fmt.Scan(&c)
		C[i] = strings.Split(c, "")
	}

	cs := make([]vector.Vector, M)
	for i, char := range C {
		cs[i] = vector.NewVector(util.StringToFloat64(char))
	}

	Vec := vectors.NewVectors(cs)

	var r string
	fmt.Print("r=")
	fmt.Scan(&r)
	R := strings.Split(r, "")
	recieve := vector.NewVector(util.StringToFloat64(R))

	decode, i := Vec.HammingMatrixDecode(recieve)

	fmt.Println("OUTPUT:")
	fmt.Print("hat_c=")
	fmt.Println(decode)
	fmt.Print("i_MD=")
	fmt.Println(i)

}