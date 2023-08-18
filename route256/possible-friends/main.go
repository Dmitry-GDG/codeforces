package main

import (
	"fmt"
	"os"
	"bufio"
	"sort"
)
 
type Pair struct {
	a1, a2 int
}
 
// func sort(s []int) []int {
// 	for i := 0; i < len(s) - 1; i ++ {
// 		for j := 0; j < len(s) - i - 1; j++ {
// 			if s[j] > s[j+1] {
// 				s[j+1], s[j] = s[j], s[j+1]
// 			}
// 		}
// 	}
// 	return s
// }
 
func notInSlice(n int, ss []int) bool {
	for _, s := range ss {
		if n == s {
			return false
		}
	}
	return true
}
 
func choseAndPrint(friends [][]int, id int, outp *[]int) {
	res := make(map[int]int)
	// var outp []int
	var max int
	
	// соберём для конкретного пользователя карту пар потенциальных друзей: номер потенциального друга - количество общих друзей
	for _, ff := range friends[id] {
		for _, f := range friends[ff] {
			// учитываем только людей, которые пока ещё не друзья конкретного пользователя
			if f != id && notInSlice(f, friends[id]) {
				res[f]++
				// выберем кандидатов с максимальным кол-вом общих друзей
				switch {
				case res[f] == max:
					*outp = append(*outp, f)
				case res[f] > max:
					max = res[f]
					*outp = append((*outp)[:0], f)
				}
			}
		}
	}
	// fmt.Printf("-- карта пар потенциальных друзей пользователя %d --\n%v\n-------------\n", id, res)
	
	// выведем на печать максимально рекомендованных друзей
	// при этом отсортируем список и восстановим нумерацию с 1
	// outp = sort(outp)
	sort.SliceStable(*outp, func(i, j int) bool {
		return (*outp)[i] < (*outp)[j]
	})
	// return outp
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
 
	var n, m, k1, k2 int // n - количество пользователей, m - количество пар друзей
	var inputData []Pair // срез входных данных - пар друзей
	var friends [][]int // срез друзей каждого пользователя (0-первый пользователь, 1-2й пользователь и тд)
	
	// распарсим  в срез
	// при этом исправим нумерацию, начнём её с 0
	fmt.Fscan(in, &n, &m)
	for ; m > 0; m-- {
		fmt.Fscan(in, &k1, &k2)
		inputData = append(inputData, Pair{k1 - 1, k2 - 1}) // нумерация была с 1
	}
	// fmt.Println("-- входные данные --\n", inputData, "\n-------------")
 
	// соберём срезы друзей по каждому пользователю
	var outp []int
	for i := 0; i < n; i++ {
		// var tmp []int
		tmp := make([]int, 0, 5)
		for _, id := range inputData {
			if id.a1 == i {
				tmp = append(tmp, id.a2)
				} else if id.a2 == i {
					tmp = append(tmp, id.a1)
				}
			}
		// sort.SliceStable(tmp, func(i, j int) bool {
			// 	return tmp[i] < tmp[j]
			// })
			// friends = append(friends, sort(tmp))
			friends = append(friends, tmp)
		}
		// fmt.Println("-- срезы друзей по каждому пользователю --\n",friends, "\n-------------")
 
		//поиск потенциальных друзей
		for i := 0; i < n; i++ {
		choseAndPrint(friends, i, &outp)
		if len(outp) == 0 {
			fmt.Fprintln(out, 0)
			// fmt.Println(0)
			// return
		} else {
			for i := 0; i < len(outp); i++ {
				fmt.Fprint(out, outp[i] + 1) // восстановим нумерацию с 1
				// fmt.Print(outp[i] + 1) // восстановим нумерацию с 1
				if i < len(outp) - 1 {
					fmt.Fprint(out, " ")
				}
			}
			fmt.Fprintln(out)
			outp = outp[:0]
		}
	}
}

