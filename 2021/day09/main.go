package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

type Coordinate struct{ Row, Col int }

func (c *Coordinate) String() string { return fmt.Sprintf("(%d, %d)", c.Row, c.Col) }

type Point struct {
	Height    int
	Neighbors []int

	Coordinate          *Coordinate
	NeighborCoordinates []*Coordinate
}

func (p *Point) IsLow() bool {
	return p.Height < p.Neighbors[0]
}

func (p *Point) Risk() int {
	return p.Height + 1
}

func NewPoint(i, j int, heightmap [][]int) *Point {
	neighbors := make([]int, 0, 4)
	coordinates := make([]*Coordinate, 0, 4)
	if i > 0 {
		neighbors = append(neighbors, heightmap[i-1][j])
		coordinates = append(coordinates, &Coordinate{Row: i - 1, Col: j})
	}

	if j > 0 {
		neighbors = append(neighbors, heightmap[i][j-1])
		coordinates = append(coordinates, &Coordinate{Row: i, Col: j - 1})
	}

	if i < len(heightmap)-1 {
		neighbors = append(neighbors, heightmap[i+1][j])
		coordinates = append(coordinates, &Coordinate{Row: i + 1, Col: j})
	}

	if j < len(heightmap[i])-1 {
		neighbors = append(neighbors, heightmap[i][j+1])
		coordinates = append(coordinates, &Coordinate{Row: i, Col: j + 1})
	}

	sort.Ints(neighbors)
	return &Point{
		Height:              heightmap[i][j],
		Neighbors:           neighbors,
		Coordinate:          &Coordinate{Row: i, Col: j},
		NeighborCoordinates: coordinates,
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

	var (
		// part1
		totalRisk int

		// part2
		points    = make([][]*Point, len(heightmap))
		lowPoints []*Point
	)
	for i := 0; i < len(heightmap); i++ {
		row := make([]*Point, len(heightmap[i]))

		for j := 0; j < len(heightmap[i]); j++ {
			p := NewPoint(i, j, heightmap)
			row[j] = p
			if p.IsLow() {
				lowPoints = append(lowPoints, p)
				totalRisk += p.Risk()
			}
		}

		points[i] = row
	}

	fmt.Println(totalRisk)

	var basinSizes []int
	for _, center := range lowPoints {
		var (
			size int
			seen = map[string]struct{}{
				center.Coordinate.String(): {},
			}
			current = center
			next    []*Point
		)

		for next /*covers first iteration */ == nil || len(next) != 0 {
			if current.Height == 9 {
				goto iter
			}

			for _, coord := range current.NeighborCoordinates {
				if _, ok := seen[coord.String()]; ok {
					continue
				}

				next = append(next, points[coord.Row][coord.Col])
				seen[coord.String()] = struct{}{}
			}
			size++

		iter:
			current, next = next[0], next[1:]
		}

		basinSizes = append(basinSizes, size)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	fmt.Println(basinSizes[0] * basinSizes[1] * basinSizes[2])
}
