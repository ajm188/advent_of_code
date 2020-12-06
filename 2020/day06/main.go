package main

import (
	"bytes"
	"flag"
	"fmt"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

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
	for _, group := range parse(data) {
		total += group.Total()
	}

	fmt.Println(total)
}
