package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmd struct {
	cycles int
	value  int
}

func main() {
	var cmds []cmd
	var total int
	captureLines("input.txt", func(v string) {
		parts := strings.Split(v, " ")

		var cycles int
		var value int
		switch parts[0] {
		case "noop":
			cycles = 1
			value = 0
		case "addx":
			cycles = 2
			value = toInt(parts[1])
		}

		total += cycles

		cmds = append(cmds, cmd{
			cycles: cycles,
			value:  value,
		})
	})

	all := map[int]int{}
	reg := 1
	cur := 1
	for _, c := range cmds {
		all[cur] = reg
		reg += c.value
		cur += c.cycles
	}

	// part a
	sum := 0
	for i := 20; i <= 220; i += 40 {
		var v int
		var ok bool
		v, ok = all[i]
		if !ok {
			for j := i; j > 0; j-- {
				v, ok = all[j]
				if ok {
					break
				}
			}
		}
		if !ok {
			panic(":/")
		}

		sum += i * v
	}

	fmt.Println("a:", sum)

	// part b
	crt := [6 * 40]string{}
	for i := 0; i < 6*40; i++ {
		var v int
		var ok bool
		v, ok = all[i+1]
		if !ok {
			for j := i; j > 0; j-- {
				v, ok = all[j]
				if ok {
					break
				}
			}
		}
		if !ok {
			panic(":((")
		}

		p := i % 40
		if p >= v-1 && p <= v+1 {
			crt[i] = "#"
		} else {
			crt[i] = "."
		}
	}
	for i := 0; i < 6*40; i += 40 {
		fmt.Println(crt[i : i+40])
	}
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
