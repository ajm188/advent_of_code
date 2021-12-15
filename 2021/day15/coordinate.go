package main

import "fmt"

type coordinate struct{ row, col int }

func (c *coordinate) In(grid [][]int) bool {
	if c.row < 0 || c.row >= len(grid) {
		return false
	}

	if c.col < 0 || c.col >= len(grid[c.row]) {
		return false
	}

	return true
}

func (c *coordinate) Neighbors() []*coordinate {
	return []*coordinate{
		{c.row - 1, c.col},
		{c.row + 1, c.col},
		{c.row, c.col - 1},
		{c.row, c.col + 1},
	}
}

func (c *coordinate) Equals(other *coordinate) bool {
	return c.row == other.row && c.col == other.col
}

func (c *coordinate) String() string { return fmt.Sprintf("(%d, %d)", c.row, c.col) }
