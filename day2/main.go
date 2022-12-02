package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	score := map[string]int{"A": 1, "B": 2, "C": 3}
	trans := map[string]string{"A": "A", "B": "B", "C": "C", "X": "A", "Y": "B", "Z": "C"}
	prio := map[string]string{"A": "C", "B": "A", "C": "B"}

	win := map[string]string{"A": "B", "B": "C", "C": "A"}
	loose := map[string]string{"A": "C", "B": "A", "C": "B"}

	input := [][]string{}
	captureLines("input.txt", func(v string) {
		input = append(input, strings.Split(v, " "))
	})

	a := 0
	b := 0
	for _, row := range input {
		enemy := row[0]
		self := trans[row[1]]

		// calc a
		switch {
		case self == enemy:
			a += score[self] + 3
		case prio[self] == enemy:
			a += score[self] + 6
		default:
			a += score[self]
		}

		// calc b
		switch row[1] {
		case "X":
			self := loose[enemy]
			b += score[self]
		case "Y":
			self := enemy
			b += score[self] + 3
		case "Z":
			self := win[enemy]
			b += score[self] + 6
		}
	}

	fmt.Println("a:", a)
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
