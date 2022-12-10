package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vec2 struct {
	x int
	y int
}

func (v vec2) isBehind(other vec2) bool {
	x := abs(v.x - other.x)
	y := abs(v.y - other.y)
	return x > 1 || y > 1
}

func (v *vec2) MoveOneStepCloser(other vec2) {
	if !v.isBehind(other) {
		return
	}

	if v.x > other.x {
		v.x--
	} else if v.x < other.x {
		v.x++
	}

	if v.y > other.y {
		v.y--
	} else if v.y < other.y {
		v.y++
	}
}

type move struct {
	dir    string
	length int
}

func main() {
	var moves []move
	captureLines("input.txt", func(v string) {
		parts := strings.Split(v, " ")
		m := move{dir: parts[0], length: toInt(parts[1])}
		moves = append(moves, m)
	})

	knots := [10]vec2{}
	visited := map[int]map[vec2]bool{}

	for i := range knots {
		visited[i] = map[vec2]bool{}
	}

	for _, m := range moves {
		switch m.dir {
		case "L":
			start := knots[0].x
			end := knots[0].x - m.length
			for j := start; j > end; j-- {
				knots[0].x--
				for i := 0; i < len(knots)-1; i++ {
					knots[i+1].MoveOneStepCloser(knots[i])
					visited[i+1][knots[i+1]] = true
				}
			}
		case "R":
			start := knots[0].x
			end := knots[0].x + m.length
			for j := start; j < end; j++ {
				knots[0].x++
				for i := 0; i < len(knots)-1; i++ {
					knots[i+1].MoveOneStepCloser(knots[i])
					visited[i+1][knots[i+1]] = true
				}
			}
		case "U":
			start := knots[0].y
			end := knots[0].y + m.length
			for j := start; j < end; j++ {
				knots[0].y++
				for i := 0; i < len(knots)-1; i++ {
					knots[i+1].MoveOneStepCloser(knots[i])
					visited[i+1][knots[i+1]] = true
				}
			}
		case "D":
			start := knots[0].y
			end := knots[0].y - m.length
			for j := start; j > end; j-- {
				knots[0].y--
				for i := 0; i < len(knots)-1; i++ {
					knots[i+1].MoveOneStepCloser(knots[i])
					visited[i+1][knots[i+1]] = true
				}
			}
		default:
			panic(":(((")
		}
	}

	fmt.Println("a:", len(visited[1]))
	fmt.Println("b:", len(visited[9]))
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

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
