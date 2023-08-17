package main

import (
	"fmt"
	"os"
	"bufio"
)

func duplRemove(t []int) []int {
	for i := 0; i < len(t) - 1; i++ {
		if t[i] == t[i+1] {
			t = append(t[:i], t[i + 1:]...)
			// t = t[:len(t)-1]
			i--
		}
	}
	return t
}

func checkAndPrint(elements []int) {
	res := make(map[int]int)
	for _, el := range elements {
		_, ok := res[el]
		if ok {
			fmt.Println("NO")
			return
		} else {
			res[el]++
		}
	}
	fmt.Println("YES")
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n, m int
	var a [][]int // срез отчетов сотрудников
	fmt.Fscan(in, &t)

	// распарсим входные данные в срезы отчетов
	for ; t > 0; t-- {
		fmt.Fscan(in, &n)
		var tmp []int
		for ; n > 0; n-- {
			fmt.Fscan(in, &m)
			tmp = append(tmp, m)
		}
		// Уберём повторы (подряд) из каждого среза
		tmp = duplRemove(tmp)
		a = append(a, tmp)
	}

	// Обработаем данные каждого сотрудника
	for _, r := range a {
		checkAndPrint(r)
	}
}