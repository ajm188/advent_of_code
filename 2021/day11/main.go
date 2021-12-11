package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

type coordinate struct{ row, col int }

func neighbors(c *coordinate) []*coordinate {
	return []*coordinate{
		{c.row - 1, c.col},
		{c.row + 1, c.col},
		{c.row, c.col - 1},
		{c.row, c.col + 1},
		{c.row - 1, c.col - 1},
		{c.row - 1, c.col + 1},
		{c.row + 1, c.col - 1},
		{c.row + 1, c.col + 1},
	}
}

func main() {
	path := flag.String("path", "input.txt", "")
	steps := flag.Int("steps", 100, "")
	debug := flag.Bool("debug", false, "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	var grid [][]int
	for i, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		if len(line) != 10 {
			cli.ExitOnError(fmt.Errorf("malformed input: line %d is not length 10 (have length=%d)", i, len(line)))
		}

		var row []int
		for j, d := range line {
			num, err := strconv.ParseInt(string(d), 10, 64)
			cli.ExitOnErrorf(err, "%s on line %d:%d", err, i, j)

			row = append(row, int(num))
		}

		grid = append(grid, row)
	}

	if len(grid) != 10 {
		cli.ExitOnError(fmt.Errorf("malformed input: did not receive 10x10 grid as input"))
	}

	var (
		flashes   int
		firstSync *int
	)

	doStep := func(step int) {
		var (
			stepFlashes  int
			chainFlashes []*coordinate
		)
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				grid[i][j]++
				if grid[i][j] > 9 {
					stepFlashes++
					chainFlashes = append(chainFlashes, neighbors(&coordinate{i, j})...)
				}
			}
		}

		for len(chainFlashes) != 0 {
			var chain []*coordinate
			for _, c := range chainFlashes {
				if c.row < 0 || c.row >= len(grid) {
					continue
				}

				if c.col < 0 || c.col >= len(grid[c.row]) {
					continue
				}

				grid[c.row][c.col]++
				if grid[c.row][c.col] == 10 { // if we're above 10, we've already flashed once this step
					chain = append(chain, neighbors(c)...)
					stepFlashes++
				}
			}

			chainFlashes = chain
		}

		flashes += stepFlashes
		if *debug {
			fmt.Printf("after step %d: %d flashes\n", step, flashes)
			for _, row := range grid {
				fmt.Printf("\t%v\n", row)
			}
		}

		if firstSync == nil && stepFlashes == len(grid)*len(grid) {
			s := step + 1
			firstSync = &s
			if *debug {
				fmt.Printf("first synchronization after step %d", s)
			}
		}

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid); j++ {
				if grid[i][j] > 9 {
					grid[i][j] = 0
				}
			}
		}
	}
	for step := 1; step <= *steps; step++ {
		doStep(step)
	}

	fmt.Println(flashes)

	for step := *steps; firstSync == nil; step++ {
		doStep(step)
	}
	fmt.Println(*firstSync)
}
