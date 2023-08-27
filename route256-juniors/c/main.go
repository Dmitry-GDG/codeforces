package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	
	var t, n, ai, max, min int
	var sign string

	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		min, max = 15, 30
		fmt.Fscan(in, &n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &sign, &ai)
			if min == -1 {
				fmt.Println("-1")
				continue
			}
			if sign == "<=" && ai < max {
				max = ai
			} else if sign == ">=" && ai > min {
				min = ai
			}
			if max < min {
				min = -1
			}
			fmt.Println(min)
		}
		if i < t - 1 {
			fmt.Println()
		}
	}
}