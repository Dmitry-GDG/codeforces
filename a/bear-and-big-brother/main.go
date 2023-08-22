package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	if a <= b {
		n := 0
		for {
			a *= 3
			b *= 2
			n++
			if a > b {
				fmt.Println(n)
				return
			}
		}
	} else {
		fmt.Println(0)
	}
}
