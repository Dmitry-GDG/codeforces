// https://codeforces.com/problemset/problem/281/A?locale=en

package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	if str[0] > 96 && str[0] < 123 {
		fmt.Println(string(str[0] - 32) + str[1:])
	} else {
		fmt.Println(str)
	}
}