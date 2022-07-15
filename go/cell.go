package main

import "fmt"

type Cell map[int]bool

func NewCell(value *int) Cell {
	if value != nil {
		return Cell{*value: false}
	}
	return Cell{
		1: false,
		2: false,
		3: false,
		4: false,
		5: false,
		6: false,
		7: false,
		8: false,
		9: false,
	}
}

func (c Cell) Remove(i int) {
	delete(c, i)
}

func (c Cell) Set(i int, b bool) {
	for n := range c {
		delete(c, n)
	}
	c[i] = b
}

func (c Cell) Print() {
	for n := range c {
		fmt.Print(n)
	}
}
