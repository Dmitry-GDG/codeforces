// https://codeforces.com/problemset/problem/282/A?locale=en

package main

import "fmt"

func main() {
	var n, outp int
	var str string
	fmt.Scan(&n)
	for ; n > 0; n-- {
		fmt.Scan(&str)
		switch str[1] {
		case '+':
			outp++
		case '-':
			outp--
		}
	}
	fmt.Println(outp)
}