package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile("\\d+")

func main() {
	// part a
	moves, stackz := load("input.txt")
	for _, mv := range moves {
		for i := 0; i < mv.Size; i++ {
			v, err := stackz[mv.From].Pop()
			if err != nil {
				continue
			}
			stackz[mv.To].Push(v)
		}
	}

	a := ""
	for i := 1; i <= len(stackz); i++ {
		a += stackz[i].Peek()
	}
	fmt.Println("a:", a)

	// part b
	moves, stackz = load("input.txt")
	for _, mv := range moves {
		var tmp []string
		for i := 0; i < mv.Size; i++ {
			v, err := stackz[mv.From].Pop()
			if err != nil {
				continue
			}
			tmp = append([]string{v}, tmp...)
		}
		for _, v := range tmp {
			stackz[mv.To].Push(v)
		}
	}

	b := ""
	for i := 1; i <= len(stackz); i++ {
		b += stackz[i].Peek()
	}
	fmt.Println("b:", b)
}

type move struct {
	From int
	To   int
	Size int
}

func load(path string) ([]move, map[int]*stack[string]) {
	d, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	text := string(d)

	var moves []move
	parts := strings.Split(text, "\n\n")
	for _, row := range strings.Split(parts[1], "\n") {
		numbers := re.FindAllString(row, 3)
		if len(numbers) != 3 {
			continue
		}
		moves = append(moves, move{
			Size: toInt(numbers[0]),
			From: toInt(numbers[1]),
			To:   toInt(numbers[2]),
		})
	}

	stackz := map[int]*stack[string]{}
	parts = strings.Split(parts[0], "\n")
	for _, row := range parts[:len(parts)-1] {
		for start := 0; start < len(row); start += 4 {
			end := start + 4
			if end >= len(row) {
				end = len(row) - 1
			}
			str := strings.TrimSpace(row[start:end])
			str = strings.Trim(str, "[")
			str = strings.Trim(str, "]")
			if str == "" {
				continue
			}

			pos := start/4 + 1
			if _, ok := stackz[pos]; !ok {
				stackz[pos] = NewStack[string]()
			}
			stackz[pos].Last(str)
		}
	}

	return moves, stackz
}

func toInt(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return v
}

type stack[T any] struct {
	list []T
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) Push(v T) {
	s.list = append(s.list, v)
}

func (s *stack[T]) Last(v T) {
	s.list = append([]T{v}, s.list...)
}

func (s *stack[T]) Pop() (T, error) {
	if len(s.list) == 0 {
		return *new(T), fmt.Errorf("empty :(")
	}
	last := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return last, nil
}

func (s *stack[T]) Peek() T {
	if len(s.list) == 0 {
		return *new(T)
	}
	return s.list[len(s.list)-1]
}
