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

	var t, n, p, tmp int
	var outp []int
    fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		goods := make(map[int]int)
		tmp = 0
		fmt.Fscan(in, &n)
		for ; n > 0; n-- {
			fmt.Fscan(in, &p)
			goods[p]++
		}
		for key, value := range goods{
			tmp += key * (value - (value / 3))
		}
		outp = append(outp, tmp)
	}
	for _, res := range outp {
		fmt.Println(res)
	}
}