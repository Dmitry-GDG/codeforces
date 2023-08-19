package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// переменные
	// n - количество процессоров
	// m - количество задач
	// a - временная переменная
	// t - момент прихода i-й задачи
	// l - время выполнения i-й задачи
	// res - выходная сумма
	var n, m int
	var a, t, l, res uint64
	
	fmt.Fscan(in, &n, &m)

	// срез всех процессоров
	proc := make([][]uint64, n, n)

	for i:= 0; i < n; i++ {
		fmt.Fscan(in, &a)
		proc[i] = []uint64{0, a}
	}
	// fmt.Println(proc)

	// отсортируем процессоры по энергопотреблению
	sort.Slice(proc, func(i, j int) bool {
		return proc[i][1] < proc[j][1]
	})
	// fmt.Println(proc)

	// подсчет времени
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &t, &l)
		for j := 0; j < n; j++ {
			if t >= proc[j][0] {
				proc[j][0] = t + l
				res += proc[j][1] * l
				break
			}
		}
	}
	fmt.Fprintln(out, res)
}