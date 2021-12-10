package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

var (
	errorScores = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
)

type NavigationLine struct {
	Line string

	parsed    bool
	complete  bool
	valid     bool
	corrupted rune
}

func (nl *NavigationLine) Parse() {
	stack := NewPairStack([]string{
		"()",
		"[]",
		"{}",
		"<>",
	})

	var corrupted *rune
	for _, r := range nl.Line {
		if _, _, err := stack.Push(r); err != nil {
			corrupted = &r
			break
		}
	}

	if corrupted != nil {
		nl.valid = false
		nl.corrupted = *corrupted
	} else if stack.Len() != 0 {
		nl.complete = false
	} else {
		nl.complete = true
		nl.valid = true
	}

	nl.parsed = true
}

func (nl *NavigationLine) SyntaxErrorScore() int {
	if !nl.parsed {
		nl.Parse()
	}

	if nl.valid {
		return 0
	}

	return errorScores[nl.corrupted]
}

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	var errorScore int
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		navline := &NavigationLine{Line: line}
		errorScore += navline.SyntaxErrorScore()
	}

	fmt.Println(errorScore)
}
