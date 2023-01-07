package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile("\\d+")

var cache = map[string]uint64{}

type valve struct {
	rate int
	to   []int
}

func main() {
	valves := map[int]valve{}

	i := 0
	lookup := map[string]int{}
	captureLines("input.txt", func(v string) {
		rate := re.FindAllString(v, -1)[0]
		parts := strings.Split(strings.ReplaceAll(v, ",", ""), " ")

		fromAsString := parts[1]
		from, ok := lookup[fromAsString]
		if !ok {
			from = i
			lookup[fromAsString] = i
			i++
		}

		var tos []int
		for _, toAsString := range parts[9:] {
			to, ok := lookup[toAsString]
			if !ok {
				to = i
				lookup[toAsString] = i
				i++
			}
			tos = append(tos, to)
		}

		valves[from] = valve{toInt(rate), tos}
	})

	start := lookup["AA"]

	// part a
	fmt.Println("a:", solve(valves, start, 30, 0))

	// part b
	var b uint64
	for _, p := range sets(valves) {
		sum := solve(valves, start, 26, p.first) + solve(valves, start, 26, p.second)
		if sum > b {
			b = sum
		}
	}
	fmt.Println("b:", b)
}

func set(bmask uint64, pos int) uint64 {
	bmask |= 1 << pos
	return bmask
}

func isSet(bmask uint64, pos int) bool {
	v := (1 << pos) & bmask
	return v != 0
}

func solve(valves map[int]valve, cur int, minLeft int, state uint64) uint64 {
	if minLeft <= 0 {
		return 0
	}

	fingerprint := fmt.Sprintf("%d;%d;%d", cur, minLeft, state)
	if v, ok := cache[fingerprint]; ok {
		return v
	}

	tos := valves[cur].to
	rate := valves[cur].rate

	var best uint64
	for _, to := range tos {
		v := solve(valves, to, minLeft-1, state)
		if v > best {
			best = v
		}
	}

	if !isSet(state, cur) && rate > 0 {
		newState := set(state, cur)
		tot := uint64((minLeft - 1) * rate)
		for _, to := range tos {
			v := solve(valves, to, minLeft-2, newState)
			if v+tot > best {
				best = v + tot
			}
		}
	}

	cache[fingerprint] = best
	return best
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

func permutate(n int) [][]bool {
	size := int(math.Pow(2, float64(n)))
	var out [][]bool
	for i := 0; i < size; i++ {
		var row []bool
		// lol....
		for _, r := range fmt.Sprintf("%0*v", n, strconv.FormatUint(uint64(i), 2)) {
			row = append(row, r == '1')
		}
		out = append(out, row)
	}
	return out
}

func invert(arr []bool) []bool {
	out := make([]bool, len(arr))
	for i, v := range arr {
		out[i] = !v
	}
	return out
}

type pair struct {
	first  uint64
	second uint64
}

// create all possible sets
func sets(valves map[int]valve) []pair {
	lookup := map[int]int{}
	var size int
	for i, valve := range valves {
		if valve.rate > 0 {
			lookup[size] = i
			size++
		}
	}

	var out []pair
	for _, first := range permutate(size) {
		var firstState uint64
		var secondState uint64
		second := invert(first)
		for i := 0; i < len(first); i++ {
			if first[i] {
				firstState = set(firstState, lookup[i])
			}
			if second[i] {
				secondState = set(secondState, lookup[i])
			}
		}
		out = append(out, pair{firstState, secondState})
	}
	return out
}
