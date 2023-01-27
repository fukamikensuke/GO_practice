package main

import (
	"fmt"
	"time"
	//"./subcal"
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println(time.Now())
	var i int = 100
	fmt.Println(i)

	var j int

	_ = j
	j = i * 2

	//fmt.Println((subcal.Add(i, j)))

	input := "Hello world"
	fmt.Println(outer(input))
}
