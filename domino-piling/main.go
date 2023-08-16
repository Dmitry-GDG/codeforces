// https://codeforces.com/problemset/problem/50/A

package main

import (
	"fmt"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	if m < 2 && n < 2 {
		fmt.Println(0)
	} else {
		fmt.Println(m * n / 2)
	}
}