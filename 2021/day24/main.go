package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ajm188/advent_of_code/2021/day24/internal/alu"
	"github.com/ajm188/advent_of_code/pkg/cli"
)

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	var instructions []alu.Instruction
	for i, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		inst, err := alu.ParseInstruction(line)
		cli.ExitOnErrorf(err, "%s (line %d)", err, i)

		instructions = append(instructions, inst)
	}

	for _, n := range []int{-2, -1, 0, 1, 2} {
		ALU := alu.ALU{Input: []int{n}}
		ALU.Execute(instructions)

		fmt.Printf("%+v\n", ALU)
	}
}
