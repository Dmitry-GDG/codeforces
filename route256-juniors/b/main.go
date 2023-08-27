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
	
	var s string
	var n, start, end int
	fmt.Fscan(in, &s, &n)
	s1 := []byte(s)
	for i := 0; i < n; i++ {
		var r string
		var tmp []byte
		fmt.Fscan(in, &start, &end, &r)
		if end < len(s1) {
			tmp = s1[end:]
		}
		r1 := []byte(r)
		s1 = append(s1[:start - 1], r1...)
		s1 = append(s1, tmp...)
	}
	fmt.Println(string(s1))
}