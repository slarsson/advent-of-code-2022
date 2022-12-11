package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var re = regexp.MustCompile("\\d+")

type monkey struct {
	Items   []int
	OpFunc  func(old, v int) int
	OpValue int
	Div     int
	True    int
	False   int
}

func (m monkey) Get() []int {
	return m.Items
}

func (m monkey) Overflow(v int) {
	for i := 0; i < len(m.Items); i++ {
		m.Items[i] %= v
	}
}

func (m *monkey) Push(item int) {
	m.Items = append(m.Items, item)
}

func (m *monkey) Pop(div int) (worry int, to int, ok bool) {
	if len(m.Items) == 0 {
		return
	}

	ok = true
	item := m.Items[0]
	m.Items = m.Items[1:]

	worry = m.OpFunc(item, m.OpValue) / div
	if worry%m.Div == 0 {
		to = m.True
	} else {
		to = m.False
	}
	return
}

func main() {
	amonkeys := load()
	bmonkeys := load()

	acounts := make([]int, len(amonkeys))
	bcounts := make([]int, len(bmonkeys))

	// a
	for i := 0; i < 20; i++ {
		for j := 0; j < len(amonkeys); j++ {
			for {
				worry, to, ok := amonkeys[j].Pop(3)
				if !ok {
					break
				}
				acounts[j]++
				amonkeys[to].Push(worry)
			}
		}
	}
	sort.Ints(acounts)
	fmt.Println("a:", acounts[len(acounts)-1]*acounts[len(acounts)-2])

	// b
	overflow := 1
	for j := 0; j < len(bmonkeys); j++ {
		overflow *= bmonkeys[j].Div
	}

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(bmonkeys); j++ {
			bmonkeys[j].Overflow(overflow)
			for {
				worry, to, ok := bmonkeys[j].Pop(1)
				if !ok {
					break
				}
				bcounts[j]++
				bmonkeys[to].Push(worry)
			}
		}
	}
	sort.Ints(bcounts)
	fmt.Println("b:", bcounts[len(bcounts)-1]*bcounts[len(bcounts)-2])
}

func toIntz(arr []string) []int {
	var out []int
	for _, v := range arr {
		intv, err := strconv.Atoi(v)
		if err != nil {
			panic("wtf")
		}
		out = append(out, intv)
	}
	return out
}

func load() []monkey {
	d, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var monkeys []monkey
	for _, block := range strings.Split(string(d), "\n\n") {
		lines := strings.Split(block, "\n")

		var opFunc func(old, v int) int
		var opValue int

		switch {
		case strings.Contains(lines[2], "old * old"):
			opFunc = func(old, v int) int { return old * old }
		case strings.Contains(lines[2], "*"):
			opFunc = func(old, v int) int { return old * v }
			opValue = toIntz(re.FindAllString(lines[2], 1))[0]
		case strings.Contains(lines[2], "+"):
			opFunc = func(old, v int) int { return old + v }
			opValue = toIntz(re.FindAllString(lines[2], 1))[0]
		default:
			panic(lines[2])
		}

		monkeys = append(monkeys, monkey{
			Items:   toIntz(re.FindAllString(lines[1], -1)),
			OpFunc:  opFunc,
			OpValue: opValue,
			Div:     toIntz(re.FindAllString(lines[3], 1))[0],
			True:    toIntz(re.FindAllString(lines[4], 1))[0],
			False:   toIntz(re.FindAllString(lines[5], 1))[0],
		})
	}

	return monkeys
}
