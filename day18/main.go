package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type vec3 struct {
	x, y, z int
}

func main() {
	cubes := []vec3{}
	captureLines("input.txt", func(v string) {
		parts := strings.Split(v, ",")
		cubes = append(cubes, vec3{
			x: toInt(parts[0]),
			y: toInt(parts[1]),
			z: toInt(parts[2]),
		})

	})

	grid := map[vec3]bool{}

	min, max := minMax(cubes)
	for x := min.x; x <= max.x; x++ {
		for y := min.y; y <= max.y; y++ {
			for z := min.z; z <= max.z; z++ {
				grid[vec3{x, y, z}] = false
			}
		}
	}
	for _, cube := range cubes {
		grid[cube] = true
	}

	// part a
	a := 0
	for _, cube := range cubes {
		a += 6 - blockedPaths(grid, cube)
	}
	fmt.Println("a:", a)

	// part b
	for vec, isCube := range grid {
		if isCube {
			continue
		}
		if !canReachOutside(grid, vec, map[vec3]bool{}) {
			grid[vec] = true
		}
	}

	b := 0
	for _, cube := range cubes {
		b += 6 - blockedPaths(grid, cube)
	}
	fmt.Println("b:", b)
}

func blockedPaths(grid map[vec3]bool, vec vec3) int {
	sum := 0
	for _, v := range []vec3{
		{vec.x + 1, vec.y, vec.z},
		{vec.x - 1, vec.y, vec.z},
		{vec.x, vec.y + 1, vec.z},
		{vec.x, vec.y - 1, vec.z},
		{vec.x, vec.y, vec.z + 1},
		{vec.x, vec.y, vec.z - 1},
	} {
		isCube, ok := grid[v]
		if !ok {
			continue
		}
		if isCube {
			sum++
		}
	}
	return sum
}

func canReachOutside(grid map[vec3]bool, vec vec3, visited map[vec3]bool) bool {
	visited[vec] = true
	for _, v := range []vec3{
		{vec.x + 1, vec.y, vec.z},
		{vec.x - 1, vec.y, vec.z},
		{vec.x, vec.y + 1, vec.z},
		{vec.x, vec.y - 1, vec.z},
		{vec.x, vec.y, vec.z + 1},
		{vec.x, vec.y, vec.z - 1},
	} {
		if visited[v] {
			continue
		}
		isCube, ok := grid[v]
		if !ok {
			return true
		}
		if isCube {
			continue
		}
		if ok := canReachOutside(grid, v, visited); ok {
			return true
		}
	}
	return false
}

func minMax(cubes []vec3) (vec3, vec3) {
	min := vec3{math.MaxInt32, math.MaxInt32, math.MaxInt32}
	max := vec3{-math.MaxInt32, -math.MaxInt32, -math.MaxInt32}
	for _, c := range cubes {
		if c.x < min.x {
			min.x = c.x
		}
		if c.x > max.x {
			max.x = c.x
		}
		if c.y < min.y {
			min.y = c.y
		}
		if c.y > max.y {
			max.y = c.y
		}
		if c.z < min.z {
			min.z = c.z
		}
		if c.z > max.z {
			max.z = c.z
		}
	}
	return min, max
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
