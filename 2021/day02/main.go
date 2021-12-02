package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

type position struct{ x, y int64 }

type command func(p position) position

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	lines := strings.Split(string(data), "\n")
	commands := make([]command, len(lines)-1)

	for i, line := range lines[:len(lines)-1] {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			cli.ExitOnError(fmt.Errorf("input %q (line:%d) did not match format '<command> <units>'", line, i))
		}

		units, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			cli.ExitOnError(fmt.Errorf("failed to parse units (%s) for input %q (line:%d): %w", parts[1], line, i, err))
		}

		switch parts[0] {
		case "forward":
			commands[i] = func(p position) position {
				return position{p.x + units, p.y}
			}
		case "down":
			commands[i] = func(p position) position {
				return position{p.x, p.y + units}
			}
		case "up":
			commands[i] = func(p position) position {
				return position{p.x, p.y - units}
			}
		default:
			cli.ExitOnError(fmt.Errorf("invalid command (%s) for input %q (line:%d); must be (forward|down|up)", parts[0], line, i))
		}
	}

	var p position
	for _, command := range commands {
		p = command(p)
	}

	fmt.Println(p.x * p.y)
}
