package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

var re = regexp.MustCompile("\\d+")

type blueprint struct {
	id       int
	ore      [3]int
	clay     [3]int
	obsidian [3]int
	geode    [3]int
}

var cache = map[string]int{}

func main() {
	rand.Seed(time.Now().UnixNano())

	var bps []blueprint
	captureLines("input.txt", func(v string) {
		numbers := re.FindAllString(v, -1)
		bps = append(bps, blueprint{
			id:       toInt(numbers[0]),
			ore:      [3]int{toInt(numbers[1]), 0, 0},
			clay:     [3]int{toInt(numbers[2]), 0, 0},
			obsidian: [3]int{toInt(numbers[3]), toInt(numbers[4]), 0},
			geode:    [3]int{toInt(numbers[5]), 0, toInt(numbers[6])},
		})
	})

	// a
	t := time.Now()
	sum := 0
	for i, bp := range bps {
		fmt.Println(i, "of", len(bps))
		cache = map[string]int{}
		v := solve(24, 0, &bp, 0, 0, 0, 1, 0, 0, 0, 0)
		fmt.Println(v, v*bp.id)
		sum += v * bp.id
		fmt.Println(time.Since(t).Seconds())
		t = time.Now()
	}
	fmt.Println("a:", sum)

	// b
	b := 1
	for i, bp := range bps[:3] {
		fmt.Println(i, "of", len(bps))
		cache = map[string]int{}
		v := solve(24, 0, &bp, 0, 0, 0, 1, 0, 0, 0, 0)
		b *= v
		fmt.Println(v)
		fmt.Println(time.Since(t).Seconds())
		t = time.Now()
	}
	fmt.Println("b:", b)
}

// lol... this is very very very slow and stupid
func solve(end int, tick int, bp *blueprint, c1 int, c2 int, c3 int, r1 int, r2 int, r3 int, r4 int, open int) int {
	if tick == end {
		return open
	}

	tick++

	fingerprint := fmt.Sprintf("%d%d%d%d%d%d%d%d", tick, r1, r2, r3, r4, c1, c2, c3)
	if v, ok := cache[fingerprint]; ok {
		return v
	}

	max := -math.MaxInt32

	maxore := bp.ore[0]
	if bp.clay[0] > maxore {
		maxore = bp.clay[0]
	}
	if bp.obsidian[0] > maxore {
		maxore = bp.obsidian[0]
	}
	if bp.geode[0] > maxore {
		maxore = bp.geode[0]
	}

	if check(c1, c2, c3, bp.geode) {
		if v := solve(end, tick, bp, c1-bp.geode[0]+r1, c2-bp.geode[1]+r2, c3-bp.geode[2]+r3, r1, r2, r3, r4+1, open+r4); v > max {
			max = v
		}
	} else {
		if r3 < bp.geode[2] {
			if check(c1, c2, c3, bp.obsidian) {
				if v := solve(end, tick, bp, c1-bp.obsidian[0]+r1, c2-bp.obsidian[1]+r2, c3-bp.obsidian[2]+r3, r1, r2, r3+1, r4, open+r4); v > max {
					max = v
				}
			}
		}
		if r1 < maxore {
			if check(c1, c2, c3, bp.ore) {
				if v := solve(end, tick, bp, c1-bp.ore[0]+r1, c2-bp.ore[1]+r2, c3-bp.ore[2]+r3, r1+1, r2, r3, r4, open+r4); v > max {
					max = v
				}
			}
		}
		if r2 < bp.obsidian[1] {
			if check(c1, c2, c3, bp.clay) {
				if v := solve(end, tick, bp, c1-bp.clay[0]+r1, c2-bp.clay[1]+r2, c3-bp.clay[2]+r3, r1, r2+1, r3, r4, open+r4); v > max {
					max = v
				}
			}
		}
	}

	if v := solve(end, tick, bp, c1+r1, c2+r2, c3+r3, r1, r2, r3, r4, open+r4); v > max {
		max = v
	}

	if max > 0 {
		cache[fingerprint] = max
	}

	return max
}

func check(c1, c2, c3 int, want [3]int) bool {
	return c1 >= want[0] && c2 >= want[1] && c3 >= want[2]
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

// Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
// Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.
