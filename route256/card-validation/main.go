package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// как предполагается проверять карту:
// в каждую ячейки запишем дополнительный флаг (равный 0)
// При этом в каждой первой попавшейся ячейке каждого цвета запишем 1 и будем искать соседей для этих ячеек (для ячеек с записанной 1)
// Если ячейка соседствует со знакомой ячейкой -
// изменим её флаг на 1 и уменьшим общее кол-во оставшихся ячеек соответствующего цвета
// Будем делать такие итерации, пока при проверках хотя бы одна ячейка найдёт нового соседа по вышеуказанным критериям
// результат будет "NO", если останутся несоседствующие ячейки

func checkLeft(sss *[][][]int, i, j int, lenColor *[]int) bool {
	if (*sss)[i][j-2][1] == 1 && (*sss)[i][j-2][0] == (*sss)[i][j][0] {
		(*sss)[i][j][1] = 1
		(*lenColor)[(*sss)[i][j][0]]--
		return true
	}
	return false
}

func checkLeftUp(sss *[][][]int, i, j int, lenColor *[]int) bool {
	// fmt.Println("checkLeftUp ", (*sss)[i][j][0], (*sss)[i-1][j-1][0])
	if (*sss)[i-1][j-1][1] == 1 && (*sss)[i-1][j-1][0] == (*sss)[i][j][0] {
		// fmt.Println("FOUND")
		(*sss)[i][j][1] = 1
		(*lenColor)[(*sss)[i][j][0]]--
		return true
	}
	return false
}

func checkLeftDown(sss *[][][]int, i, j int, lenColor *[]int) bool {
	if (*sss)[i+1][j-1][1] == 1 && (*sss)[i+1][j-1][0] == (*sss)[i][j][0] {
		(*sss)[i][j][1] = 1
		(*lenColor)[(*sss)[i][j][0]]--
		return true
	}
	return false
}

func checkRightUp(sss *[][][]int, i, j int, lenColor *[]int) bool {
	if (*sss)[i-1][j+1][1] == 1 && (*sss)[i-1][j+1][0] == (*sss)[i][j][0] {
		(*sss)[i][j][1] = 1
		(*lenColor)[(*sss)[i][j][0]]--
		return true
	}
	return false
}

func checkRightDown(sss *[][][]int, i, j int, lenColor *[]int) bool {
	if (*sss)[i+1][j+1][1] == 1 && (*sss)[i+1][j+1][0] == (*sss)[i][j][0] {
		(*sss)[i][j][1] = 1
		(*lenColor)[(*sss)[i][j][0]]--
		return true
	}
	return false
}

func checkRight(sss *[][][]int, i, j int, lenColor *[]int) bool {
	if (*sss)[i][j+2][1] == 1 && (*sss)[i][j+2][0] == (*sss)[i][j][0] {
		(*sss)[i][j][1] = 1
		(*lenColor)[(*sss)[i][j][0]]--
		return true
	}
	return false
}

func check(ss [][]byte, wg *sync.WaitGroup, done chan string) {
	defer wg.Done()
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// for _, s := range ss {
	// 	fmt.Fprintln(out, s)
	// }
	// fmt.Fprintln(out, "-------")

	n := len(ss)
	m := len(ss[0])

	lenColor := make([]int, 26, 26)
	var sss [][][]int
	for i := 0; i < len(ss); i++ {
		var tmpSss [][]int
		for j := 0; j < len(ss[i]); j++ {
			if ss[i][j] == 46 {
				tmpSss = append(tmpSss, []int{46, 0})
			} else if lenColor[ss[i][j]-65] == 0 {
				tmpSss = append(tmpSss, []int{int(ss[i][j]) - 65, 1}) // 1 - флаг проверки соседей
				lenColor[ss[i][j]-65]++
			} else {
				tmpSss = append(tmpSss, []int{int(ss[i][j]) - 65, 0})
				lenColor[ss[i][j]-65]++
			}
		}
		sss = append(sss, tmpSss)
	}
	// fmt.Println("lenColor initial: ", lenColor)
	// уменьшим на 1 количество каждого имеющегося цвета, чтоб избежать ошибок, связанных с единственным или единственным оставшимся элементом
	lenC := 0
	for i := 0; i < 26; i++ {
		if lenColor[i] > 0 {
			lenColor[i]--
		}
	}
	// fmt.Println("lenColor after correction: ", lenColor)
	// for i := 0; i < len(sss); i++ {
	// 	fmt.Println(sss[i])
	// }

	for rcount := 0; rcount < m*n*m*n; rcount++ {
		foundInCircle := false
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if sss[i][j][1] == 1 || sss[i][j][0] == 46 {
					continue
				}
				if j > 1 {
					foundInCircle = foundInCircle || checkLeft(&sss, i, j, &lenColor)
				}
				if i > 0 && j > 0 {
					foundInCircle = foundInCircle || checkLeftUp(&sss, i, j, &lenColor)
				}
				if i < n-1 && j > 0 {
					foundInCircle = foundInCircle || checkLeftDown(&sss, i, j, &lenColor)
				}
				if j < m-2 {
					foundInCircle = foundInCircle || checkRight(&sss, i, j, &lenColor)
				}
				if i > 0 && j < m-1 {
					foundInCircle = foundInCircle || checkRightUp(&sss, i, j, &lenColor)
				}
				if i < n-1 && j < m-1 {
					foundInCircle = foundInCircle || checkRightDown(&sss, i, j, &lenColor)
				}
			}
		}

		if !foundInCircle {
			break
		}
	}

	lenC = 0
	for i := 0; i < 26; i++ {
		lenC += lenColor[i]
	}
	if lenC == 0 {
		done <- "YES"
	} else {
		done <- "NO"
	}
	close(done)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n, m int
	var str string

	fmt.Fscan(in, &t)

	wg := new(sync.WaitGroup)
	wg.Add(t)
	outp := make([]string, t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &m)
		ss := make([][]byte, n, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &str)
			ss[i] = []byte(str)
		}
		// for _, s := range ss {
		// 	fmt.Fprintln(out, s)
		// }
		// fmt.Fprintln(out, "-------")
		done := make(chan string)
		go check(ss, wg, done)
		outp[i] = <-done
	}
	wg.Wait()
	for _, o := range outp {
		fmt.Fprintln(out, o)
	}
}
