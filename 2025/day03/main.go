package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

type BatteryBank []rune

func (bb BatteryBank) MaxJoltage() int {
	var (
		d1    rune
		d1idx int
	)
	for i, b := range bb[0 : len(bb)-1] {
		if b > d1 {
			d1 = b
			d1idx = i
		}
	}

	var d2 rune
	for _, b := range bb[d1idx+1:] {
		if b > d2 {
			d2 = b
		}
	}

	i, _ := strconv.Atoi(fmt.Sprintf("%c%c", d1, d2))
	return i
}

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	var banks []BatteryBank
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		banks = append(banks, []rune(line))
	}

	var totalJoltage int
	for _, bank := range banks {
		totalJoltage += bank.MaxJoltage()
	}
	fmt.Println(totalJoltage)
}
