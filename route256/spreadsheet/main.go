package main

import (
	"fmt"
	"bufio"
	"os"
)

// пузырьковая сортировка
func sort(t [][]int, n int) [][]int {
	for i := 0; i < len(t)-1; i++ {
		for j := 0; j < len(t)-i-1; j++ {
		   if (t[j][n] > t[j+1][n]) {
			  t[j], t[j+1] = t[j+1], t[j]
		   }
		}
	 }
	 return t
}

func print(t [][]int) {
	for _, i := range t {
		for _, j := range i {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var pivotTable [][][]int // сводная таблица всех наборов
	var pivotClicks [][]int // сводная таблица кликов по каждому набору
	// i-ый срез сводной таблицы кликов хранит инфо по кликам на i-м наборе

	var t, n, m, k int
	fmt.Fscan(in, &t)

	// Заполняем сводные таблицы всех наборов и кликов
	for ; t > 0; t-- {
		var table [][]int

		fmt.Fscan(in, &n, &m)

		for ; n > 0; n-- {
			var str []int
			for j := 0; j < m; j++ {
				fmt.Fscan(in, &k)
				str = append(str, k)
			}
			table = append(table, str)
		}
		pivotTable = append(pivotTable, table)
		fmt.Fscan(in, &k)
		var str []int
		for ; k > 0; k-- {
			fmt.Fscan(in, &n)
			str = append(str, n - 1) // нумерация в условии начинается с 1
		}
		pivotClicks = append(pivotClicks, str)
	}

	// обрабатываем каждый набор
	for i := 0; i < len(pivotTable); i++ {
		// fmt.Println(pivotTable[i])
		// fmt.Println(pivotClicks[i])
		// fmt.Println()
		for j := 0; j < len(pivotClicks[i]); j++ {
			pivotTable[i] = sort(pivotTable[i], pivotClicks[i][j])
		}
		// Печать изменённого набора
		print(pivotTable[i])
		if i < len(pivotTable) - 1 {
			fmt.Println()
		}
	}
}