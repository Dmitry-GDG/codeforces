package main

import (
	"fmt"
	"os"
	"bufio"
)

type Pair struct {
	ap, bp int // ap - порядковый номер, bp - уровень
}

func modDiff(a, b int) int {
	dif := a - b
	if dif < 0 {
		dif *= -1
	}
	return dif
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n, aTmp int

	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		// Создадим срезы пар:
		// для a (срез работников): ap - порядковый номер, bp - уровень
		// для outp (итоговый срез): ab - номер первого работника, bp - номер второго работника
		var a, outp []Pair

		fmt.Fscan(in, &n)
		// заполняем срез работников
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &aTmp)
			a = append (a, Pair{i + 1, aTmp})
		}

		// заполняем итоговый срез парами порядковых номеров работников
		for i := 0; i < n / 2; i++ {
			next := 1
			dif := modDiff(a[0].bp, a[next].bp)
			for j := 2; j < len(a); j++ {
				dif1 := modDiff(a[0].bp, a[j].bp)
				if dif1 < dif {
					next = j
					dif = dif1
				}
			}
			p := Pair{a[0].ap, a[next].ap}
			outp = append(outp, p)
			// удаляем использованных работников из среза работников
			if i < n / 2 - 1 {
				a = append(a[:next], a[(next + 1):]...)
				a = a[1:]
			}
		}

		// печатаем итоговый срез
		for _, o := range outp {
			fmt.Println(o.ap, o.bp)
		}

		// печатаем пустую линию (если это не последний элемент)
		if t > 1 {
			fmt.Println()
		}
	}
}