package main

import (
	"bytes"
	"flag"
	"fmt"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

// Group represents a set of answers from a group. The zero-byte key is reserved
// to track group size.
type Group map[byte]int

func (g Group) Add(b byte) {
	x, ok := g[b]
	if !ok {
		x = 0
	}

	g[b] = x + 1
}

func (g Group) Total() int {
	return len(g)
}

func (g Group) UnanimousTotal() int {
	total := 0

	for k, v := range g {
		if k == 0 {
			continue
		}

		if v == g[0] {
			total++
		}
	}

	return total
}

func parse(data []byte) []Group {
	var groups []Group

	cur := Group(map[byte]int{})

	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		line := bytes.TrimSpace(line)

		if bytes.Equal(line, []byte{}) {
			groups = append(groups, cur)
			cur = map[byte]int{}

			continue
		}

		for _, question := range line {
			cur.Add(question)
		}

		cur.Add(0)
	}

	if len(cur) > 0 {
		groups = append(groups, cur)
	}

	return groups
}

func main() {
	path := flag.String("input", "", "path to input, reads from stdin if empty")

	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	total := 0
	unanimousTotal := 0

	for _, group := range parse(data) {
		total += group.Total()
		unanimousTotal += group.UnanimousTotal()
	}

	fmt.Println(total)
	fmt.Println(unanimousTotal)
}
