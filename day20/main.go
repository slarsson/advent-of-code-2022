package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	next  *node
	prev  *node
	value int
}

func (n *node) Find(count int, length int) int {
	cursor := n
	for i := 0; i < count%length; i++ {
		cursor = cursor.next
	}
	return cursor.value
}

func (n *node) Reorder(length int) {
	var steps int
	if n.value > 0 {
		steps = n.value % (length - 1)
	} else {
		steps = -(-n.value % (length - 1))
	}

	cursor := n
	if steps < 0 {
		for i := 0; i > steps; i-- {
			prev := cursor.prev
			next := cursor.next

			cursor.next = prev
			cursor.prev = prev.prev
			prev.prev.next = cursor
			prev.prev = cursor
			prev.next = next
			next.prev = prev
		}
	} else {
		for i := 0; i < steps; i++ {
			prev := cursor.prev
			next := cursor.next

			cursor.next = next.next
			cursor.prev = next
			next.next.prev = cursor
			next.next = cursor
			next.prev = prev
			prev.next = next
		}
	}
}

func main() {
	var anumbers []int
	var bnumbers []int
	captureLines("input.txt", func(v string) {
		n := toInt(v)
		anumbers = append(anumbers, n)
		bnumbers = append(bnumbers, 811589153*n)
	})

	size := len(anumbers)

	// part a
	var zero *node
	_, anodes := build(anumbers)
	for _, node := range anodes {
		if node.value == 0 {
			zero = node
		}
		node.Reorder(size)
	}
	fmt.Println("a:", zero.Find(1000, size)+zero.Find(2000, size)+zero.Find(3000, size))

	// part b
	zero = nil
	_, bnodes := build(bnumbers)
	for i := 0; i < 10; i++ {
		for _, node := range bnodes {
			if node.value == 0 {
				zero = node
			}
			node.Reorder(size)
		}
	}
	fmt.Println("b:", zero.Find(1000, size)+zero.Find(2000, size)+zero.Find(3000, size))
}

func build(numbers []int) (*node, []*node) {
	start := &node{next: nil, prev: nil, value: numbers[0]}
	nodelist := []*node{start}
	cursor := start
	for _, n := range numbers[1:] {
		next := &node{next: nil, prev: cursor, value: n}
		nodelist = append(nodelist, next)
		cursor.next = next
		cursor = next
	}
	cursor.next = start
	start.prev = cursor
	return start, nodelist
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
