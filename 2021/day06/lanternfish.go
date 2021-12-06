package main

import "fmt"

type Lanternfish struct {
	counter int
}

func NewLanternfish(counter int) *Lanternfish {
	if counter < 0 {
		panic(fmt.Errorf("invalid initial counter value (%d); must be non-negative", counter))
	}

	return &Lanternfish{counter: counter}
}

func (lf *Lanternfish) Step() *Lanternfish {
	if lf.counter == 0 {
		lf.counter = 6
		return NewLanternfish(8)
	}

	lf.counter--
	return nil
}

func (lf *Lanternfish) String() string { return fmt.Sprintf("%d", lf.counter) }
