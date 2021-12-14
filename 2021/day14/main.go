package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

var (
	path  = flag.String("path", "input.txt", "")
	steps = flag.Int("steps", 10, "")
	debug = flag.Bool("debug", false, "")
)

func main() {
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	var (
		polymer    string
		rules      = map[string]string{}
		ruleRegexp = regexp.MustCompile(`^([A-Z]+) -> ([A-Z]+)$`)
	)
	for i, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		if m := ruleRegexp.FindStringSubmatch(line); m != nil {
			if len(m) != 3 {
				cli.ExitOnError(fmt.Errorf("bad input: polymer insertion rule (%q, line %d) does not match %s", line, i, ruleRegexp))
			}

			if _, ok := rules[m[1]]; ok {
				cli.ExitOnError(fmt.Errorf("bad input: two different replacements specified for %s; first is `%s -> %s`, second is `%s -> %s` on line %d", m[1], m[1], rules[m[1]], m[1], m[2], i))
			}

			rules[m[1]] = m[2]
		} else {
			if polymer != "" {
				cli.ExitOnError(fmt.Errorf("bad input: two polymer templates given (first: %q), second (%q, on line %d)", polymer, line, i))
			}

			polymer = line
		}
	}

	if *debug {
		log.Printf("Template: %s", polymer)
	}

	for step := 1; step <= *steps; step++ {
		var buf strings.Builder
		for i := 0; i < len(polymer)-1; i++ {
			pair := polymer[i : i+2]
			buf.WriteByte(pair[0])

			if insertion, ok := rules[pair]; ok {
				buf.WriteString(insertion)
			}
		}

		buf.WriteByte(polymer[len(polymer)-1])

		polymer = buf.String()
		if *debug {
			log.Printf("After step %d: %s", step, polymer)
		}
	}

	counts := map[string]int{}
	for _, r := range polymer {
		counts[string(r)]++
	}

	if *debug {
		log.Println(counts)
	}

	var max, min *int
	for _, count := range counts {
		count := count // prevent loop shadowing
		if max == nil || *max < count {
			max = &count
		}

		if min == nil || *min > count {
			min = &count
		}
	}

	fmt.Println(*max - *min)
}
