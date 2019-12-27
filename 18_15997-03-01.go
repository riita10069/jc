//␣プログラミング言語：golang
//␣コンパイル方法：env go build -o [ビルド後のファイル名] [ファイル名]
//␣実行方法：./[ビルド後のファイル名]

package main

import (
	"fmt"
	"strings"

	. "github.com/riita10069/jc/util"
	. "github.com/riita10069/jc/vector"
)

func main() {
	var a int
	fmt.Print("INPUT LENGTH:")
	fmt.Scan(&a)

	var c1 string
	fmt.Print("c1=")
	fmt.Scan(&c1)

	var c2 string
	fmt.Print("c2=")
	fmt.Scan(&c2)

	C1 := strings.Split(c1, "")
	C2 := strings.Split(c2, "")

	c_1 := NewVector(StringToFloat64(C1))
	c_2 := NewVector(StringToFloat64(C2))

	distance := c_1.HammingDistance(c_2)

	fmt.Println("OUTPUT:")
	fmt.Print(distance)
}
