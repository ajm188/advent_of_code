package main

import (
	"bytes"
	"flag"
	"fmt"
	"regexp"
	"strconv"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

type Policy struct {
	min    int
	max    int
	letter []byte
}

func (p *Policy) Match(b []byte) bool {
	count := bytes.Count(b, p.letter)
	return p.min <= count && count <= p.max
}

func (p *Policy) Match2(b []byte) bool {
	if len(b) < p.max {
		return false
	}

	low := b[p.min-1 : p.min]
	high := b[p.max-1 : p.max]

	if bytes.Equal(low, p.letter) && bytes.Equal(high, p.letter) {
		return false
	}

	return bytes.Equal(low, p.letter) || bytes.Equal(high, p.letter)
}

type Password struct {
	policy   *Policy
	password []byte
}

// nolint:gochecknoglobals
var passwordRegexp = regexp.MustCompile(
	`^(?P<min>\d+)-(?P<max>\d+) (?P<letter>[a-z]): (?P<password>.*)$`,
)

func readPasswords(data []byte) ([]*Password, error) {
	lines := bytes.Split(data, []byte("\n"))
	passwords := make([]*Password, 0, len(lines))

	for _, line := range lines {
		if bytes.Equal(line, []byte("")) {
			continue
		}

		match := passwordRegexp.FindSubmatch(line)
		if match == nil {
			continue
		}

		min, err := strconv.ParseInt(string(match[passwordRegexp.SubexpIndex("min")]), 10, 64)
		if err != nil {
			return nil, err
		}

		max, err := strconv.ParseInt(string(match[passwordRegexp.SubexpIndex("max")]), 10, 64)
		if err != nil {
			return nil, err
		}

		pw := &Password{
			password: match[passwordRegexp.SubexpIndex("password")],
			policy: &Policy{
				min:    int(min),
				max:    int(max),
				letter: match[passwordRegexp.SubexpIndex("letter")],
			},
		}

		passwords = append(passwords, pw)
	}

	return passwords, nil
}

func main() {
	path := flag.String("input", "", "path to input, reads from stdin if empty")

	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	passwords, err := readPasswords(data)
	cli.ExitOnError(err)

	count := 0
	count2 := 0

	for _, pw := range passwords {
		if pw.policy.Match(pw.password) {
			count++
		}

		if pw.policy.Match2(pw.password) {
			count2++
		}
	}

	fmt.Println(count)
	fmt.Println(count2)
}
