package main

import (
	"fmt"
	"github.com/riita10069/jc/finite_field"
	"github.com/riita10069/jc/number"
	"github.com/riita10069/jc/util"
)

func main()  {
	var p int
	fmt.Println("INPUT:")
	fmt.Print("p=")
	fmt.Scan(&p)

	arr := make([]number.Number, p)
	for i := range make([]int, p){
		arr[i] = number.NewNumber(p, i)
	}
	finiteNumberField := finite_field.NewFiniteNumberField(p, arr)

	fmt.Println("ADD")
	fmt.Println("  | 0 1 2 3 4")
	fmt.Println("--+-----------")
	addTable := finiteNumberField.AddTable()
	util.PrintArrayOfIntArray(addTable)

	fmt.Println("MUL")
	fmt.Println("  | 0 1 2 3 4")
	fmt.Println("--+-----------")
	mulTable := finiteNumberField.MulTable()
	util.PrintArrayOfIntArray(mulTable)

	fmt.Println("SUB")
	fmt.Println("  | 0 1 2 3 4")
	fmt.Println("--+-----------")
	subTable := finiteNumberField.SubTable()
	util.PrintArrayOfIntArray(subTable)

	fmt.Println("DIV")
	fmt.Println("  | 0 1 2 3 4")
	fmt.Println("--+-----------")
	divTable := finiteNumberField.DivTable()
	util.PrintArrayOfArray(divTable)

}