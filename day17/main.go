package main

import (
	"bufio"
	"fmt"
	"os"
)

type vec2 struct {
	x int
	y int
}

type rock []vec2

func (r rock) Copy() rock {
	out := []vec2{}
	for _, v := range r {
		out = append(out, v)
	}
	return out
}

func (r rock) MoveFromOrigin(x, y int) {
	for i := 0; i < len(r); i++ {
		r[i].x += x
		r[i].y += y
	}
}

func (r rock) Left(xMin int, g grid) {
	for i := 0; i < len(r); i++ {
		x := r[i].x - 1
		if x < xMin {
			return
		}
		if g[vec2{x, r[i].y}] {
			return
		}
	}
	for i := 0; i < len(r); i++ {
		r[i].x--
	}
}

func (r rock) Right(xMax int, g grid) {
	for i := 0; i < len(r); i++ {
		x := r[i].x + 1
		if x > xMax {
			return
		}
		if g[vec2{x, r[i].y}] {
			return
		}
	}
	for i := 0; i < len(r); i++ {
		r[i].x++
	}
}

func (r rock) Down(g grid) bool {
	for i := 0; i < len(r); i++ {
		if g[vec2{r[i].x, r[i].y - 1}] {
			return false
		}
	}
	for i := 0; i < len(r); i++ {
		r[i].y--
	}
	return true
}

type jet struct {
	cursor int
	moves  []int
}

func (j *jet) Add(v int) {
	j.moves = append(j.moves, v)
}

func (j *jet) Next() vec2 {
	if j.cursor >= len(j.moves) {
		j.cursor = 0
	}
	out := vec2{j.moves[j.cursor], 0}
	j.cursor++
	return out
}

type grid map[vec2]bool

func (g grid) Add(vecs []vec2) {
	for _, v := range vecs {
		g[v] = true
	}
}

func (g grid) Max() int {
	max := 0
	for v := range g {
		if v.y > max {
			max = v.y
		}
	}
	return max
}

func (g grid) ToString(yMin int, yMax int) string {
	all := ""
	for i := yMin; i <= yMax; i++ {
		line := ""
		for j := 0; j <= 6; j++ {
			if g[vec2{j, i}] {
				line += "#"
			} else {
				line += "."
			}
		}
		all += line
	}
	return all
}

func main() {
	moves := jet{}
	captureLines("input.txt", func(v string) {
		for _, r := range v {
			if r == '<' {
				moves.Add(-1)
				continue
			}
			moves.Add(1)
		}
	})

	rocks := []rock{
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 1}},
		{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
	}

	g := grid{
		{0, 0}: true,
		{1, 0}: true,
		{2, 0}: true,
		{3, 0}: true,
		{4, 0}: true,
		{5, 0}: true,
		{6, 0}: true,
	}

	seen := map[string][2]int64{}

	yMax := 0
	yskip := int64(0)

	i := 0
	iskip := int64(0)

	for {
		r := rocks[i%len(rocks)].Copy()
		r.MoveFromOrigin(2, yMax+4)

		for {
			dir := moves.Next()
			if dir.x == 1 {
				r.Right(6, g)
			} else {
				r.Left(0, g)
			}
			if !r.Down(g) {
				break
			}
		}
		g.Add(r)
		yMax = g.Max()

		if i == 2021 {
			fmt.Println("a:", yMax)
		}

		i++
		iskip++

		// 60 == random interval that contains cycle
		str := g.ToString(yMax-60, yMax)
		if v, ok := seen[str]; ok {
			di := int64(i) - v[0]    // number of rocks in interval
			dy := int64(yMax) - v[1] // height of interval

			// how many intervals until 1 trillion
			size := (1000000000000 - iskip) / di
			iskip += size * di
			yskip += size * dy
		} else {
			seen[str] = [2]int64{int64(i), int64(yMax)}
		}

		// hmm...
		if iskip == 1000000000000 {
			fmt.Println("b:", yskip+int64(yMax))
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
