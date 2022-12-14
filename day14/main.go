package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

type cave map[coord]int

func (c cave) search(v coord, maxy int, miny int) bool {
	if v.y > maxy || v.y < miny {
		return false
	}
	if _, ok := c[v]; ok {
		return false
	}

	down := coord{v.x, v.y + 1}
	if _, ok := c[down]; !ok {
		return c.search(down, maxy, miny)
	}

	left := coord{v.x - 1, v.y + 1}
	if _, ok := c[left]; !ok {
		return c.search(left, maxy, miny)
	}

	right := coord{v.x + 1, v.y + 1}
	if _, ok := c[right]; !ok {
		return c.search(right, maxy, miny)
	}

	c[v] = -1
	return true
}

func (c cave) minMaxY() (int, int) {
	min := math.MaxInt32
	max := -math.MaxInt32
	for v := range c {
		if v.y < min {
			min = v.y
		}
		if v.y > max {
			max = v.y
		}
	}
	return min, max
}

func (c cave) minMaxX() (int, int) {
	min := math.MaxInt32
	max := -math.MaxInt32
	for v := range c {
		if v.x < min {
			min = v.x
		}
		if v.x > max {
			max = v.x
		}
	}
	return min, max
}

func main() {
	var coords [][]coord
	captureLines("input.txt", func(v string) {
		c := []coord{}
		for _, number := range strings.Split(v, " -> ") {
			parts := strings.Split(number, ",")
			x := toInt(parts[0])
			y := toInt(parts[1])
			c = append(c, coord{x, y})
		}
		coords = append(coords, c)
	})

	acave := cave{}
	bcave := cave{}
	for _, line := range coords {
		for i := 0; i < len(line)-1; i++ {
			for _, c := range coordsBetween(line[i], line[i+1]) {
				acave[c] = 1
				bcave[c] = 1
			}
		}
	}

	// part a
	before := len(acave)
	_, maxy := acave.minMaxY()
	for {
		if !acave.search(coord{500, 0}, maxy, 0) {
			break
		}
	}
	after := len(acave)
	fmt.Println("a:", after-before)

	// part b
	end := maxy + 2
	for _, c := range coordsBetween(coord{0, end}, coord{1000, end}) {
		bcave[c] = 1
	}
	for {
		if !bcave.search(coord{500, 0}, math.MaxInt32, 0) {
			break
		}
	}
	count := 0
	for _, v := range bcave {
		if v == -1 {
			count++
		}
	}
	fmt.Println("b:", count)

}

func coordsBetween(start coord, end coord) []coord {
	var coords []coord
	if start.x == end.x {
		min, max := minMax(start.y, end.y)
		for y := min; y <= max; y++ {
			coords = append(coords, coord{start.x, y})
		}
		return coords
	}
	if start.y == end.y {
		min, max := minMax(start.x, end.x)
		for x := min; x <= max; x++ {
			coords = append(coords, coord{x, start.y})
		}
		return coords
	}
	panic("should not happen")
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

func toInt(v string) int {
	intv, err := strconv.Atoi(v)
	if err != nil {
		return -1
	}
	return intv
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
