// https://codeforces.com/problemset/problem/263/A

package main

import (
	"fmt"
)

func main() {
	var i, j, nmb, outp int
	for i = 1; i < 6; i++ {
		for j = 1; j < 6; j++ {
			fmt.Scan(&nmb)
			if nmb == 1 {
				if i < 3 {
					outp += 3 - i 
				} else {
					outp += i - 3
				}
				if j < 3 {
					outp += 3 - j
				} else {
					outp += j - 3
				}
			}
		}
	}
	fmt.Println(outp)
}