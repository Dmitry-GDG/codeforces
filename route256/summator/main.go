package main

import (
	"fmt"
)

func main() {
	var a, b, t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
