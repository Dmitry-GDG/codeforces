package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"time"
	"sort"
)

type Pair struct {
	d1, d2 time.Time
}

// // bubblesort
// func sort(ts []Pair) []Pair {
// 	for i := 0; i < len(ts) - 1; i++ {
// 		for j := 0; j < len(ts) - i - 1; j++ {
// 			if ts[j + 1].d1.Before(ts[j].d1) {
// 				ts[j + 1], ts[j] = ts[j], ts[j + 1]
// 			}
// 		}
// 	}
// 	return ts
// }

func checkAndPrint(dd []string) {
	var timeSlice []Pair
	for _, d := range dd {
		str := strings.Split(d, "-")
		// fmt.Println(str)
		
		// проверяем первое условие задания
		d1, err := time.Parse("15:04:05", str[0])
		if err != nil {
			fmt.Println("NO")
			return
		}
		d2, err := time.Parse("15:04:05", str[1])
		if err != nil {
			fmt.Println("NO")
			return
		}

		// проверяем второе условие задания
		if d2.Before(d1) {
			fmt.Println("NO")
			return
		}
		
		// добавляем даты в слайс
		timeSlice = append(timeSlice, Pair{d1, d2})
	}
	// fmt.Println(timeSlice)

	// проверяем третье условие задания
	// сортируем слайс по первой дате
	// timeSlice = sort(timeSlice)

	sort.SliceStable(timeSlice, func(i, j int) bool {
		return timeSlice[i].d1.Before(timeSlice[j].d1)
	})
	// fmt.Println("--\n", timeSlice)


	for i := 0; i < len(timeSlice) - 1; i++ {
		if timeSlice[i].d2.After(timeSlice[i + 1].d1) || timeSlice[i].d2.Equal(timeSlice[i + 1].d1) {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n int
	var data [][]string // срез наборов входных данных
	
	// распарсим входные данные с срез
	fmt.Fscan(in, &t)
	
	for ; t > 0; t-- {
		var tmp []string
		var s string
		fmt.Fscan(in, &n)
		for ; n > 0; n-- {
			fmt.Fscan(in, &s)
			tmp = append(tmp, s)
		}
		// добавим набор данных в срез
		data = append(data, tmp)
	}
	// fmt.Println(data)

	// Проверяем наборы данных
	for _, dd := range data {
		checkAndPrint(dd)
	}
}