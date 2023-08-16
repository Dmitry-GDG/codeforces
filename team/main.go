// https://codeforces.com/problemset/problem/231/A?locale=en

package main

import (
	"fmt"
)

func main() {
	var n, a, b, c, outp int
	fmt.Scan(&n)
	for ; n > 0; n-- {
		fmt.Scan(&a, &b, &c)
		if a + b + c > 1 {
			outp++
		}
	}
	fmt.Println(outp)
}