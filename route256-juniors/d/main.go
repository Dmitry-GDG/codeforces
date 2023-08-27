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
	
	var t, k, n, m int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var s string
		var outp [][]byte
		fmt.Fscan(in, &k, &n, &m)
		for j := 0; j < n; j ++ {
			fmt.Fscan(in, &s)
			outp = append(outp, []byte(s))
		}
		for i := 1; i < k; i++ {
			for j := 0; j < n; j ++ {
				fmt.Fscan(in, &s)
				for h := 0; h < m; h++ {
					if outp[j][h] == '.' && s[h] != '.' {
						outp[j][h] = s[h]
					}
				}
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Print(string(outp[i][j]))
			}
			fmt.Println()
		}
		if t > 1 {
			fmt.Println()
		}
	}
}