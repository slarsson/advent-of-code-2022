package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := [][]rune{}
	captureLines("input.txt", func(v string) {
		var row []rune
		for _, r := range v {
			row = append(row, r)
		}
		input = append(input, row)
	})

	// part a
	a := 0
	for _, row := range input {
		size := len(row) / 2
		seen := map[rune]bool{}
		for _, lhs := range row[:size] {
			for _, rhs := range row[size:] {
				if lhs == rhs && !seen[lhs] {
					a += score(lhs)
					seen[lhs] = true
				}
			}
		}
	}
	fmt.Println("a:", a)

	// part b
	b := 0
	for i := 0; i < len(input); i += 3 {
		freq := map[rune]int{}
		for _, row := range input[i : i+3] {
			seen := map[rune]bool{}
			for _, r := range row {
				if seen[r] {
					continue
				}
				freq[r]++
				seen[r] = true
			}
		}
		for r, count := range freq {
			if count > 2 {
				b += score(r)
			}
		}
	}
	fmt.Println("b:", b)
}

func score(r rune) int {
	if v := r - int32(96); v > 0 {
		return int(v)
	}
	return 26 + int(r-int32(64))
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
