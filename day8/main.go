package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type matrix []int

func (m matrix) size() int {
	return int(math.Sqrt(float64(len(m))))
}

func (m matrix) PrevCols(i int) []int {
	var out []int
	size := m.size()
	pos := i - size
	for {
		if pos < 0 {
			break
		}
		out = append(out, m[pos])
		pos -= size
	}
	return out
}

func (m matrix) NextCols(i int) []int {
	var out []int
	size := m.size()
	pos := i + size
	for {
		if pos >= len(m) {
			break
		}
		out = append(out, m[pos])
		pos += size
	}
	return out
}

func (m matrix) PrevRow(i int) []int {
	size := m.size()
	start := i - (i % size)
	if start < 0 {
		start = 0
	}
	return reverse(m[start:i])
}

func (m matrix) NextRow(i int) []int {
	size := m.size()
	end := size - (i % size)
	return m[i+1 : i+end]
}

func main() {
	var m matrix
	captureLines("input.txt", func(v string) {
		for _, r := range v {
			v := int(r - '0')
			m = append(m, v)
		}
	})

	// a
	a := 0
	for i, cur := range m {
		row := isLarger(cur, m.PrevRow(i)) || isLarger(cur, m.NextRow(i))
		col := isLarger(cur, m.PrevCols(i)) || isLarger(cur, m.NextCols(i))
		if row || col {
			a++
		}
	}
	fmt.Println("a:", a)

	// b
	b := 0
	for i, cur := range m {
		left := m.PrevRow(i)
		right := m.NextRow(i)
		up := m.PrevCols(i)
		down := m.NextCols(i)

		if len(left) == 0 || len(right) == 0 || len(up) == 0 || len(down) == 0 {
			continue
		}

		score := 1
		for _, dir := range [][]int{up, left, down, right} {
			l := 0
			for _, tree := range dir {
				if tree >= cur {
					l++
					break
				}
				l++
			}
			if l > 0 {
				score *= l
			}
		}

		if score > b {
			b = score
		}
	}
	fmt.Println("b:", b)
}

func captureLines(path string, f func(v string)) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f(scanner.Text())
	}
}

func isLarger(have int, arr []int) bool {
	for _, v := range arr {
		if v >= have {
			return false
		}
	}
	return true
}

func reverse(arr []int) []int {
	out := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		out = append(out, arr[i])
	}
	return out
}
