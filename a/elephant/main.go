package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Scan(&x)

	outp := x / 5

	if x % 5 != 0 {
		outp++
	}

	fmt.Println(outp)
}
