package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	outp := make(map[string]int)
	for _, s := range str {
		outp[string(s)]++
	}
	if len(outp) % 2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
