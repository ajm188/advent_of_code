package main

import (
	"strings"

	"github.com/ajm188/advent_of_code/pkg/sets"
)

type Paper struct {
	points  *sets.Strings
	_points []*point
}

func NewPaper(_points []*point) *Paper {
	points := sets.NewStrings()
	for _, p := range _points {
		points.Insert(p.String())
	}

	return &Paper{
		points:  points,
		_points: _points,
	}
}

func (paper *Paper) MaxX() int {
	var x int
	for _, p := range paper._points {
		if p.x > x {
			x = p.x
		}
	}

	return x
}

func (paper *Paper) MaxY() int {
	var y int
	for _, p := range paper._points {
		if p.y > y {
			y = p.y
		}
	}

	return y
}

func (paper *Paper) String() string {
	rows, cols := paper.MaxY(), paper.MaxX()

	var buf strings.Builder
	for y := 0; y <= rows; y++ {
		for x := 0; x <= cols; x++ {
			if paper.points.Has((&point{x: x, y: y}).String()) {
				buf.WriteRune('#')
			} else {
				buf.WriteRune('.')
			}
		}

		if y != rows {
			buf.WriteString("\n")
		}
	}

	return buf.String()
}
