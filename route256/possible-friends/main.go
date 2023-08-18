package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func notInSlice(n int, ss []int) bool {
	for _, s := range ss {
		if n == s {
			return false
		}
	}
	return true
}

func chose(friends [][]int, id int, outp *[]int) {
	res := make(map[int]int)
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

	// отсортируем
	sort.SliceStable(*outp, func(i, j int) bool {
		return (*outp)[i] < (*outp)[j]
	})
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, k1, k2 int // n - количество пользователей, m - количество пар друзей

	fmt.Fscan(in, &n, &m)
	friends := make([][]int, n, n) // срез друзей каждого пользователя (0-первый пользователь, 1-2й пользователь и тд)
	// распарсим входные данные в срез по каждому пользователю
	// при этом исправим нумерацию, начнём её с 0
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &k1, &k2)
		friends[k1 - 1] = append(friends[k1 - 1], k2 - 1) // нумерация во входных данных была с 1
		friends[k2 - 1] = append(friends[k2 - 1], k1 - 1) // нумерация во входных данных была с 1
	}
	// fmt.Println("-- срезы друзей по каждому пользователю --\n",friends, "\n-------------")

	//поиск потенциальных друзей
	outp := make([]int, 0, 5)
	for i := 0; i < n; i++ {
		chose(friends, i, &outp)
		// выведем на печать максимально рекомендованных друзей
		// при этом  восстановим нумерацию с 1
		if len(outp) == 0 {
			fmt.Fprintln(out, 0)
		} else {
			for i := 0; i < len(outp); i++ {
				fmt.Fprint(out, outp[i]+1) // восстановим нумерацию с 1
				if i < len(outp)-1 {
					fmt.Fprint(out, " ")
				}
			}
			fmt.Fprintln(out)
			outp = outp[:0]
		}
	}
}
