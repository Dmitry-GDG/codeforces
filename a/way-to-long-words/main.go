// https://codeforces.com/problemset/problem/71/A

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inp string
	var n int
	fmt.Scan(&n)
	for ; n > 0; n-- {
		fmt.Scan(&inp)
		if len(inp) > 10 {
			fmt.Println(string(inp[0]) + strconv.Itoa(len(inp) - 2) + string(inp[len(inp) - 1]))
		} else {
			fmt.Println(inp)
		}
	}
}