package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	Size int
	Type string
	Prev *node
	Next map[string]*node
}

func (n *node) Add(name string, v *node) {
	if n.Next == nil {
		n.Next = map[string]*node{}
	}
	n.Next[name] = v
}

func (n *node) Move(name string) *node {
	if n.Next == nil {
		return nil
	}
	v, ok := n.Next[name]
	if !ok {
		return nil
	}
	return v
}

func (n *node) Back() *node {
	return n.Prev
}

func (n *node) Sums() ([]int, int) {
	tot := n.Size
	out := []int{}
	for _, next := range n.Next {
		sums, v := next.Sums()
		tot += v
		out = append(out, sums...)
	}
	if n.Type == "dir" {
		out = append(out, tot)
	}
	return out, tot
}

func main() {
	var cmds []string
	captureLines("input.txt", func(v string) {
		cmds = append(cmds, v)
	})

	start := &node{Type: "dir"}
	cursor := start

	for _, cmd := range cmds {
		switch {
		case strings.HasPrefix(cmd, "$ cd"):
			path := cmd[5:]
			if path == ".." {
				cursor = cursor.Back()
			} else if path == "/" {
				cursor = start
			} else {
				cursor = cursor.Move(path)
			}
		case strings.HasPrefix(cmd, "dir"):
			file := cmd[4:]
			cursor.Add(file, &node{Prev: cursor, Type: "dir"})
		default:
			parts := strings.Split(cmd, " ")
			cursor.Add(parts[1], &node{Prev: cursor, Type: "file", Size: toInt(parts[0])})
		}
	}

	ints, _ := start.Sums()

	// a
	a := 0
	for _, sum := range ints[:len(ints)-1] {
		if sum <= 100000 {
			a += sum
		}
	}
	fmt.Println("a:", a)

	// b
	sort.Ints(ints)
	free := 70000000 - ints[len(ints)-1]
	for _, v := range ints {
		current := free + v
		if current >= 30000000 {
			fmt.Println("b:", v)
			break
		}
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
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return n
}
