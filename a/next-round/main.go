// https://codeforces.com/problemset/problem/158/A?locale=en

package main

import (
	"fmt"
)

func main() {
	var n, k, nmb, outp int
	fmt.Scan(&n, &k)
	for i := 0; i < k; i++ {
		fmt.Scan(&nmb)
		if nmb > 0 {
			outp++
		}
	}
	tmp := nmb
	for i := k; i < n; i++ {
		fmt.Scan(&nmb)
		if nmb == tmp && tmp > 0 {
			outp++
		}
	}
	fmt.Println(outp)
}