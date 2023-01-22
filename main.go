package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println(time.Now())
	var i int = 100
	fmt.Println(i)

	var j int

	_ = j
	j = i * 2
}
