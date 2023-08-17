// https://codeforces.com/problemset/problem/4/A

package main

import (
	"fmt"
)

func main() {
	var n int
	// for {
	// 	fmt.Print("Введите вес арбуза: ")
	// 	fmt.Scan(&n)
	// 	if n < 1 || n > 100 {
	// 		return
	// 	}
	// 	if n % 2 == 0 && n > 2{
	// 		fmt.Println("YES")
	// 	} else {
	// 		fmt.Println("NO")
	// 	}
	// }
	fmt.Scan(&n)
	if n < 1 || n > 100 {
		return
	}
	if n % 2 == 0 && n > 2{
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}