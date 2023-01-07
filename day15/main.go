package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type vec2 struct {
	x, y int
}

func main() {
	dists := map[vec2]int{}
	beacons := map[vec2]bool{}
	sensors := map[vec2]bool{}

	var min vec2
	var max vec2

	captureLines("input.txt", func(line string) {
		var x1, x2, y1, y2 int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &x1, &y1, &x2, &y2)
		sensor := vec2{x1, y1}
		beacon := vec2{x2, y2}

		dist := manhattan(sensor, beacon)
		dists[sensor] = dist

		if sensor.x-dist < min.x {
			min.x = sensor.x - dist
		}
		if sensor.x+dist > max.x {
			max.x = sensor.x + dist
		}

		if sensor.y-dist < min.y {
			min.y = sensor.y - dist
		}
		if sensor.y+dist > max.y {
			max.y = sensor.y + dist
		}

		beacons[beacon] = true
		sensors[sensor] = true
	})

	// part a
	var notreachable int
	y := 2000000
	for x := min.x; x <= max.x; x++ {
		vec := vec2{x, y}

		var reachable bool
		if !beacons[vec] && !sensors[vec] {
			for sensor := range sensors {
				reach := dists[sensor]
				if reach >= manhattan(vec, sensor) {
					reachable = true
					break
				}
			}
		}
		if reachable {
			notreachable++
		}
	}
	fmt.Println("a:", notreachable)

	// part b
	vecs := map[vec2]int{}
	start := 0
	end := 4000000

	for sensor := range sensors {
		dist := dists[sensor]
		for _, vec := range frame(sensor, dist) {
			if vec.x < start || vec.x > end {
				continue
			}
			if vec.y < start || vec.y > end {
				continue
			}
			vecs[vec]++
		}
	}

	var maybe []vec2
	for vec, count := range vecs {
		if count > 1 {
			maybe = append(maybe, vec)
		}
	}

	for _, vec := range maybe {
		var reachable bool
		for sensor := range sensors {
			reach := dists[sensor]
			dist := manhattan(vec, sensor)
			if reach >= dist {
				reachable = true
				break
			}
		}
		if !reachable {
			fmt.Println("b:", vec.x*4000000+vec.y)
			break
		}
	}
}

func manhattan(v1 vec2, v2 vec2) int {
	dx := float64(v1.x - v2.x)
	dy := float64(v1.y - v2.y)
	return int(math.Abs(dx) + math.Abs(dy))
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

func frame(vec vec2, dist int) []vec2 {
	top := vec2{vec.x, vec.y - dist - 1}
	right := vec2{vec.x + dist + 1, vec.y}
	bottom := vec2{vec.x, vec.y + dist + 1}
	left := vec2{vec.x - dist - 1, vec.y}

	out := []vec2{top, right, bottom, left}
	for i := 0; i < dist; i++ {
		out = append(out, vec2{top.x + i, top.y + i})
		out = append(out, vec2{right.x - i, right.y + i})
		out = append(out, vec2{bottom.x - i, bottom.y + i})
		out = append(out, vec2{left.x + i, left.y + i})
	}
	return out
}
