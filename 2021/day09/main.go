package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

type Point struct {
	Height    int
	Neighbors []int
}

func (p *Point) IsLow() bool {
	return p.Height < p.Neighbors[0]
}

func (p *Point) Risk() int {
	return p.Height + 1
}

func NewPoint(i, j int, heightmap [][]int) *Point {
	neighbors := make([]int, 0, 4)
	if i > 0 {
		neighbors = append(neighbors, heightmap[i-1][j])
	}

	if j > 0 {
		neighbors = append(neighbors, heightmap[i][j-1])
	}

	if i < len(heightmap)-1 {
		neighbors = append(neighbors, heightmap[i+1][j])
	}

	if j < len(heightmap[i])-1 {
		neighbors = append(neighbors, heightmap[i][j+1])
	}

	sort.Ints(neighbors)
	return &Point{
		Height:    heightmap[i][j],
		Neighbors: neighbors,
	}
}

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	var heightmap [][]int
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if line == "" {
			continue
		}

		var row []int
		for j, d := range line {
			digit, err := strconv.ParseInt(string(d), 10, 64)
			cli.ExitOnErrorf(err, "could not parse digit at %d:%d: %s", i, j, err)
			row = append(row, int(digit))
		}

		heightmap = append(heightmap, row)
	}

	// make sure it's even
	rowLengths := map[int][]int{}
	for i, row := range heightmap {
		rowsWithLength := rowLengths[len(row)]
		rowsWithLength = append(rowsWithLength, i)
		rowLengths[len(row)] = rowsWithLength
	}

	if len(rowLengths) > 1 {
		var (
			buf   strings.Builder
			count int
		)
		for l, rows := range rowLengths {
			fmt.Fprintf(&buf, "rows %v have length %d", rows, l)
			count++
			if count != len(rowLengths) {
				buf.WriteString(", ")
			}
		}

		cli.ExitOnError(fmt.Errorf("rows do not all have same length: %s", buf.String()))
	}

	var points []*Point
	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			points = append(points, NewPoint(i, j, heightmap))
		}
	}

	var totalRisk int
	for _, p := range points {
		if p.IsLow() {
			totalRisk += p.Risk()
		}
	}

	fmt.Println(totalRisk)
}
