package main

import "github.com/ajm188/advent_of_code/pkg/sets"

type route struct {
	route    []*coordinate
	_seen    *sets.Strings
	_current *coordinate
	_score   int
	_grid    [][]int
}

func NewRoute(start *coordinate, grid [][]int) *route {
	return &route{
		route:    []*coordinate{start},
		_seen:    sets.NewStrings(),
		_current: start,
		_score:   0,
		_grid:    grid,
	}
}

func (r *route) Add(coord *coordinate) *route {
	if !coord.In(r._grid) {
		return nil
	}

	if r._seen.Has(coord.String()) {
		return nil
	}

	return &route{
		route:    append(append([]*coordinate{}, r.route...), coord),
		_seen:    sets.NewStrings(append(r._seen.Items(), coord.String())...),
		_current: coord,
		_score:   r._score + r._grid[coord.row][coord.col],
		_grid:    r._grid,
	}
}

func (r *route) Score() int { return r._score }
