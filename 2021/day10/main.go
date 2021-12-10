package main

import (
	"flag"
	"fmt"
	"sort"
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
	autocompleteScores = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
)

type NavigationLine struct {
	Line string

	parsed bool

	complete       bool
	autocompletion []rune

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
		nl.valid = true
		nl.autocompletion = stack.Complete()
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

func (nl *NavigationLine) AutocompleteScore() int {
	if !nl.parsed {
		nl.Parse()
	}

	if !nl.valid || nl.complete {
		return 0
	}

	var score int
	for _, r := range nl.autocompletion {
		score = (5 * score) + autocompleteScores[r]
	}
	return score
}

func main() {
	path := flag.String("path", "input.txt", "")
	debug := flag.Bool("debug", false, "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	var (
		errorScore       int
		completionScores []int
	)
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		navline := &NavigationLine{Line: line}
		errorScore += navline.SyntaxErrorScore()
		if score := navline.AutocompleteScore(); score != 0 {
			if *debug {
				fmt.Println(line, "\t", string(navline.autocompletion), "\t", score)
			}
			completionScores = append(completionScores, score)
		}
	}

	fmt.Println(errorScore)
	sort.Ints(completionScores)
	fmt.Println(completionScores[len(completionScores)/2])
}
