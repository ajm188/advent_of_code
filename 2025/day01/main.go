package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

type Direction int

const (
	Left Direction = iota
	Right
)

type Rotation struct {
	Direction Direction
	Amount    int64
}

func ParseRotation(input string) (Rotation, error) {
	if len(input) < 2 {
		return Rotation{}, fmt.Errorf("invalid rotation: %s", input)
	}

	amount, err := strconv.ParseInt(input[1:], 10, 64)
	if err != nil {
		return Rotation{}, fmt.Errorf("%w: invalid rotation amount: %s", err, input[1:])
	}

	switch input[0] {
	case 'L':
		return Rotation{Direction: Left, Amount: amount}, nil
	case 'R':
		return Rotation{Direction: Right, Amount: amount}, nil
	default:
		return Rotation{}, fmt.Errorf("invalid rotation: %s", input)
	}
}

func (r Rotation) Apply(d Dial) Dial {
	amount := r.Amount
	if r.Direction == Left {
		amount = -amount
	}

	pos := d.pos + amount
	if pos < 0 {
		pos = 100 + pos
	}

	pos = (pos % 100)

	dial := Dial{pos: pos, zeroCount: d.zeroCount}
	if pos == 0 {
		dial.zeroCount++
	}
	return dial
}

type Dial struct {
	pos       int64
	zeroCount int
}

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	lines := strings.Split(string(data), "\n")
	rotations := make([]Rotation, len(lines)-1)
	for i, line := range lines {
		if line == "" {
			continue
		}

		rotations[i], err = ParseRotation(line)
		cli.ExitOnError(err)
	}

	dial := Dial{pos: 50, zeroCount: 0}
	for _, r := range rotations {
		dial = r.Apply(dial)
	}

	fmt.Println(dial.zeroCount)
}
