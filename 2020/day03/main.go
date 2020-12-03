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
				return nil, fmt.Errorf("unknown symbol %s at %d:%d", char, i, j)
			}
		}

		forest = append(forest, row)
	}

	return forest, nil
}

type Position struct {
	x     int
	y     int
	trees int
}

func (p *Position) traverse(forest Forest) {
	for p.y < len(forest) {
		if forest[p.y][p.x] {
			p.trees++
		}

		row := forest[(p.y+1)%len(forest)]
		p.x = (p.x + 3) % len(row)
		p.y++
	}
}

func main() {
	path := flag.String("input", "", "path to input, reads from stdin if empty")

	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	forest, err := NewForest(data)
	cli.ExitOnError(err)

	pos := &Position{}
	pos.traverse(forest)
	fmt.Println(pos.trees)
}
