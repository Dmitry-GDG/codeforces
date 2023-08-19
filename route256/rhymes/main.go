package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	var s string

	var ss []string

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s)
		ss = append(ss, s)
	}

	fmt.Fscan(in, &q)
	for ; q > 0; q-- {
		fmt.Fscan(in, &s)
		done := false
		ls := s
		for i := 0; i < len(ls); i++ {
			for _, w := range ss {
				if strings.HasSuffix(w, s) && w != ls {
					fmt.Fprintln(out, w)
					done = true
					break
				}
			}
			if done {
				break
			} else if len(s) == 1 {
				if ss[0] != ls {
					fmt.Fprintln(out, ss[0])
				} else {
					fmt.Fprintln(out, ss[1])
				}
				break
			} else {
				s = s[1:]
			}
		}
	}
}
