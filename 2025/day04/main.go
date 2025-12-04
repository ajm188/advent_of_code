package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

type Grid [][]bool

func ParseGrid(input string) Grid {
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	grid := make([][]bool, len(lines))

	for i, line := range lines {
		grid[i] = make([]bool, len(line))
		for j, c := range line {
			grid[i][j] = c == '@'
		}
	}

	return grid
}

func (g Grid) String() string {
	var buf strings.Builder
	for i := range g {
		for j := range g[i] {
			if g[i][j] {
				buf.WriteString("@")
			} else {
				buf.WriteString(".")
			}
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func (g Grid) Neighbors(x, y int) int {
	count := 0
	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			if i == 0 && j == 0 {
				continue
			}

			if x+i >= 0 && x+i < len(g) && y+j >= 0 && y+j < len(g[x+i]) && g[x+i][y+j] {
				count++
			}
		}
	}

	return count
}

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	grid := ParseGrid(string(data))
	neighborCounts := make([][]int, len(grid))
	for i := range grid {
		neighborCounts[i] = make([]int, len(grid[i]))

		for j := range grid[i] {
			if grid[i][j] {
				neighborCounts[i][j] = grid.Neighbors(i, j)
			}
		}
	}

	// fmt.Println(grid.String())

	rolls := 0
	for i, counts := range neighborCounts {
		for j, neighbors := range counts {
			if grid[i][j] && neighbors < 4 {
				rolls++
				grid[i][j] = false
			}
		}
	}

	fmt.Println(rolls)
}
