package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

var (
	path = flag.String("path", "input.txt", "")

	debug = flag.Bool("debug", false, "")
)

func main() {
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	var grid [][]int
	for i, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		row := make([]int, 0, len(line))
		for j, n := range line {
			num, err := strconv.ParseInt(string(n), 10, 64)
			cli.ExitOnErrorf(err, "%s line %d:%d", err, i, j)

			row = append(row, int(num))
		}

		if len(grid) > 0 && len(grid[0]) != len(row) {
			cli.ExitOnError(fmt.Errorf("bad input: row %d has different length (%d)", i, len(row)))
		}

		grid = append(grid, row)
	}

	if *debug {
		var buf strings.Builder
		for i, row := range grid {
			for _, num := range row {
				fmt.Fprintf(&buf, "%d", num)
			}

			if i < len(grid)-1 {
				buf.WriteString("\n")
			}
		}

		log.Println(buf.String())
	}

	var (
		routes = []*route{NewRoute(&coordinate{0, 0}, grid)}
		end    = &coordinate{
			row: len(grid) - 1,
			col: len(grid[len(grid)-1]) - 1,
		}
		scores = map[string]int{}
	)
	for len(routes) != 0 {
		if *debug {
			log.Printf("routes to check: %d", len(routes))
		}

		r := routes[0]
		routes = routes[1:]

		if s, ok := scores[r._current.String()]; ok && s < r.Score() {
			continue
		}

		scores[r._current.String()] = r.Score()
		for _, neighbor := range r._current.Neighbors() {
			if r2 := r.Add(neighbor); r2 != nil {
				if s, ok := scores[r2._current.String()]; ok && s < r2.Score() {
					continue
				}

				scores[r2._current.String()] = r2.Score()
				routes = append(routes, r2)
			}
		}
	}

	fmt.Println(scores[end.String()])
}
