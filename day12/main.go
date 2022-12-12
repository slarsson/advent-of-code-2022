package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type matrix struct {
	values []int32
	width  int
	height int
}

func (m matrix) Left(i int) (int, bool) {
	p := i % m.width
	if p-1 < 0 {
		return -1, false
	}
	return i - 1, true
}

func (m matrix) Right(i int) (int, bool) {
	p := i % m.width
	if p+1 >= m.width {
		return -1, false
	}
	return i + 1, true
}

func (m matrix) Up(i int) (int, bool) {
	if i-m.width < 0 {
		return -1, false
	}
	return i - m.width, true
}

func (m matrix) Down(i int) (int, bool) {
	if i+m.width >= m.width*m.height {
		return -1, false
	}
	return i + m.width, true
}

func (m matrix) OK(from, to int) bool {
	tv := m.values[to]
	fv := m.values[from]
	return tv <= fv+1
}

type edge struct {
	from int
	to   int
}

func main() {
	var values []int32
	var width int
	var height int
	var start int
	var end int
	captureLines("input.txt", func(v string) {
		width = len(v)
		for i, r := range v {
			if r == 'S' {
				r = 'a'
				start = height*width + i
			}
			if r == 'E' {
				r = 'z'
				end = height*width + i
			}
			values = append(values, r-'0')
		}
		height++
	})

	m := matrix{values, width, height}

	startValues := []int{}
	edges := map[int][]int{}
	for from, v := range m.values {
		if v == 49 { // a -> 49
			startValues = append(startValues, from)
		}
		edges[from] = []int{}
		if to, ok := m.Left(from); ok && m.OK(from, to) {
			edges[from] = append(edges[from], to)
		}
		if to, ok := m.Right(from); ok && m.OK(from, to) {
			edges[from] = append(edges[from], to)
		}
		if to, ok := m.Up(from); ok && m.OK(from, to) {
			edges[from] = append(edges[from], to)
		}
		if to, ok := m.Down(from); ok && m.OK(from, to) {
			edges[from] = append(edges[from], to)
		}
	}

	res := map[int]int{}
	for _, init := range startValues {
		vertices := map[int]int{}
		for from := range m.values {
			vertices[from] = math.MaxInt32
		}

		vertices[init] = 0

		queue := []int{init}
		seen := map[int]bool{}
		for {
			if len(queue) == 0 {
				break
			}

			sort.Slice(queue, func(i, j int) bool {
				return vertices[queue[i]] < vertices[queue[j]]
			})

			head := queue[0]
			queue = queue[1:]

			if seen[head] {
				break
			}
			seen[head] = true

			cur := vertices[head]
			for _, e := range edges[head] {
				newScore := cur + 1
				if newScore < vertices[e] {
					vertices[e] = newScore
					queue = append(queue, e)
				}
			}
		}

		res[init] = vertices[end]
	}

	min := math.MaxInt32
	for _, v := range res {
		if v < min {
			min = v
		}
	}

	fmt.Println("a:", res[start])
	fmt.Println("b:", min)
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
