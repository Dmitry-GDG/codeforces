package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

type Pair struct {
	np, ip int // np - порядковый номер друга, ip - его последняя карточка
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	
	var n, m, a int
	var friends []Pair

	fmt.Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		friends = append(friends, Pair{i + 1, a})
	}
	sort.Slice(friends, func(i, j int) bool {
		return(friends[i].ip < friends[j].ip)
	})

	if friends[n-1].ip >= m {
		fmt.Println("-1")
		return
	}

	friends[0].ip += 1
	cardNow := friends[0].ip + 1
	if cardNow > m + 1 {
		fmt.Println("-1")
		return
	}

	i := 1
	for ; i < n && cardNow <= m; i++ {
		if friends[i].ip < cardNow {
			friends[i].ip = cardNow
			cardNow++
		} else {
			friends[i].ip += 1
			cardNow = friends[i].ip + 1
		}
		if i < n-1 && cardNow > m {
			fmt.Println("-1")
			return
		} 
	}
	sort.Slice(friends, func(i, j int) bool {
		return(friends[i].np < friends[j].np)
	})
	for i := 0; i < n; i++ {
		fmt.Println(friends[i].ip)
	}
}