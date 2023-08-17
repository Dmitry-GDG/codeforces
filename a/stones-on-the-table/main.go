package main

import (
	"fmt"
)

func main() {
	var str string
	var n, outp int
	fmt.Scan(&n, &str)
	r := str[0]
	for i := 1; i < len(str); i++ {
		if str[i] == r {
			outp++
		} else {
			r = str[i]
		}
	}
	fmt.Println(outp)
}