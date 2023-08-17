// https://codeforces.com/problemset/problem/112/A?locale=en

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	str1 = strings.ToUpper(str1)
	str2 = strings.ToUpper(str2)
	switch {
	case str1 == str2:
		fmt.Println(0)
	case str1 < str2:
		fmt.Println(-1)
	case str1 > str2:
		fmt.Println(1)
	}
}