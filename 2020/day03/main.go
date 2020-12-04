package main

import (
	"bytes"
	"flag"
	"fmt"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

type Forest [][]bool

func NewForest(data []byte) (Forest, error) {
	lines := bytes.Split(data, []byte("\n"))
	forest := make([][]bool, 0, len(lines))

	for i, line := range lines {
		if bytes.Equal(line, []byte("")) {
			continue
		}

		row := make([]bool, len(line))

		for j, char := range line {
			switch char {
			case '.':
				row[j] = false
			case '#':
				row[j] = true
			default:
				return nil, fmt.Errorf("unknown symbol %s at %d:%d", []byte{char}, i, j)
			}
		}

		forest = append(forest, row)
	}

	return forest, nil
}

type Position struct {
	x      int
	y      int
	trees  int
	deltaX int
	deltaY int
}

func (p *Position) traverse(forest Forest) {
	for p.y < len(forest) {
		if forest[p.y][p.x] {
			p.trees++
		}

		row := forest[(p.y+p.deltaY)%len(forest)]
		p.x = (p.x + p.deltaX) % len(row)
		p.y += p.deltaY
	}
}

func main() {
	path := flag.String("input", "", "path to input, reads from stdin if empty")

	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	forest, err := NewForest(data)
	cli.ExitOnError(err)

	pos := &Position{
		deltaX: 3,
		deltaY: 1,
	}

	positions := []*Position{
		{
			deltaX: 1,
			deltaY: 1,
		},
		pos,
		{
			deltaX: 5,
			deltaY: 1,
		},
		{
			deltaX: 7,
			deltaY: 1,
		},
		{
			deltaX: 1,
			deltaY: 2,
		},
	}

	product := 1

	for _, position := range positions {
		position.traverse(forest)
		product *= position.trees
	}

	fmt.Println(pos.trees)
	fmt.Println(product)
}
