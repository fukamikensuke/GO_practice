package main

import (
	"fmt"
)

func outer(input string) {
	var s4 string = input
	fmt.Println(s4)
}
func main() {
	var i int = 100
	fmt.Println(i)

	var j, k int
	j = i * 2
	k = i * 300

	fmt.Println(i, j, k)

	a := 100
	fmt.Println(a)

	input := "Hello world"
	outer(input)
	// := は 最初は宣言で使っていいがその後上書き不可能

	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Println("%T,%T\n", fl64, fl)

	zero := 0.0
	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

	var arr2 [3]string = [3]string{"A", "B", "c"}
	fmt.Println(arr2[0])

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"A", "B", "C", "D"}
	fmt.Println(arr4)

}
