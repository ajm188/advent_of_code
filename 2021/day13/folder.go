package main

import (
	"fmt"

	"github.com/ajm188/advent_of_code/pkg/sets"
)

var (
	_ Folder = (*HorizontalFolder)(nil)
	_ Folder = (*VerticalFolder)(nil)
)

type Folder interface {
	Fold(*Paper) *Paper
	String() string
}

type HorizontalFolder struct {
	y int
}

func (hf *HorizontalFolder) Fold(paper *Paper) *Paper {
	foldedPaper := &Paper{
		points: sets.NewStrings(),
	}

	maxX, maxY := paper.MaxX(), paper.MaxY()

	var foldCutoff *int
	for y := maxY; y > hf.y; y-- {
		for x := 0; x <= maxX; x++ {
			p := &point{x: x, y: y}
			distance := y - hf.y
			opposite := &point{x: x, y: hf.y - distance}

			if opposite.y < 0 {
				continue
			}

			if foldCutoff == nil || *foldCutoff > opposite.y {
				foldCutoff = &opposite.y
			}

			if paper.points.Has(p.String()) || paper.points.Has(opposite.String()) {
				foldedPaper.points.Insert(opposite.String())
				foldedPaper._points = append(foldedPaper._points, opposite)
			}
		}
	}

	for y := 0; y < *foldCutoff; y++ {
		for x := 0; x <= maxX; x++ {
			p := &point{x: x, y: y}
			if paper.points.Has(p.String()) {
				foldedPaper.points.Insert(p.String())
				foldedPaper._points = append(foldedPaper._points, p)
			}
		}
	}

	return foldedPaper
}

func (hf HorizontalFolder) String() string { return fmt.Sprintf("fold along y=%d", hf.y) }

type VerticalFolder struct {
	x int
}

func (vf *VerticalFolder) Fold(paper *Paper) *Paper {
	foldedPaper := &Paper{
		points: sets.NewStrings(),
	}

	maxX, maxY := paper.MaxX(), paper.MaxY()

	var foldCutoff *int
	for x := maxX; x > vf.x; x-- {
		for y := 0; y <= maxY; y++ {
			p := &point{x: x, y: y}
			distance := x - vf.x
			opposite := &point{x: vf.x - distance, y: y}

			if opposite.x < 0 {
				continue
			}

			if foldCutoff == nil || *foldCutoff > opposite.x {
				foldCutoff = &opposite.x
			}

			if paper.points.Has(p.String()) || paper.points.Has(opposite.String()) {
				foldedPaper.points.Insert(opposite.String())
				foldedPaper._points = append(foldedPaper._points, opposite)
			}
		}
	}

	for x := 0; x < *foldCutoff; x++ {
		for y := 0; y <= maxY; y++ {
			p := &point{x: x, y: y}
			if paper.points.Has(p.String()) {
				foldedPaper.points.Insert(p.String())
				foldedPaper._points = append(foldedPaper._points, p)
			}
		}
	}

	return foldedPaper
}

func (vf *VerticalFolder) String() string { return fmt.Sprintf("fold along x=%d", vf.x) }
