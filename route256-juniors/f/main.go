package main

import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
	// "bytes"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	
	var t, n int
	
	
	fmt.Fscan(in, &t)
	var s, sCombo string
	for ; t > 0; t-- {
		fmt.Fscan(in, &n)
		for ; n > 0; n-- {
			fmt.Fscan(in, &s)
			sCombo += s
		}
		if t > 1 {
			sCombo += ", "
		}
	}
	data, _ := json.Marshal(sCombo)
	// data = bytes.Trim(data, "'\'")
	for i := 0; i < len(data); i++ {
		if i < len(data) - 1 && data[i] == 92 && data[i+1] == 34 {
			// i++
			continue
		}
		fmt.Printf("%s" , string(data[i]))
	}
	fmt.Println()
	// fmt.Printf("%s\n", data)
}