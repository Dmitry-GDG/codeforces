// https://codeforces.com/problemset/problem/339/A

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, outp string
	fmt.Scan(&str)
	strSlice := strings.Split(str, "+")
	for i := 0; i < len(strSlice) - 1; i++ {
		for j := 0; j < len(strSlice) - i - 1; j++ {
			if strSlice[j] > strSlice[j + 1] {
				strSlice[j], strSlice[j + 1] = strSlice[j + 1], strSlice[j]
			}
		}
	}
	//fmt.Print(strSlice)
	outp = string(strSlice[0])
	for i := 1; i < len(strSlice); i++ {
		outp += "+" + string(strSlice[i])
	}
	fmt.Println(outp)
}