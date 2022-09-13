package main

import (
	"fmt"
)

func fifth(root int) int {
	return (root + 7) % 12
}

func maj3(root int) int {
	return (root + 4) % 12
}

func min3(root int) int {
	return (root + 3) % 12
}

func maj7(root int) int {
	return (root + 10) % 12
}

func min7(root int) int {
	return (root + 9) % 12
}

type RowName int

const (
	Maj3s RowName = iota
	Roots
	Majs
	Mins
	Dom7s
	Dim7s
)

func (n RowName) String() string {
	switch n {
	case Roots:
		return "roots"
	case Maj3s:
		return "major 3rds"
	case Majs:
		return "majors"
	case Mins:
		return "minors"
	case Dom7s:
		return "dominant 7ths"
	case Dim7s:
		return "diminished 7ths"
	}
	return "unknown"
}

type accordion struct {
	numRows    int
	numButtons int
	centerInd  int
	rows       map[RowName][20][3]int
}

func newAccordion() *accordion {
	a := accordion{}

	// always a full size accordion
	// control indexing for other sizes
	a.rows = make(map[RowName][20][3]int)

	var new_bs [20][3]int
	for i := 0; i < 20; i++ {
		if i == 0 {
			new_bs[i] = [3]int{9, 9, 9} // bottom note of full size
		} else {
			prev_note := new_bs[i-1][0]
			new_bs[i] = [3]int{fifth(prev_note), fifth(prev_note), fifth(prev_note)}
		}
	}
	a.rows[Roots] = new_bs

	for i := 0; i < 20; i++ {
		r := a.rows[Roots][i][0]
		new_bs[i] = [3]int{maj3(r), maj3(r), maj3(r)}
	}
	a.rows[Maj3s] = new_bs

	for i := 0; i < 20; i++ {
		r := a.rows[Roots][i][0]
		new_bs[i] = [3]int{r, maj3(r), fifth(r)}
	}
	a.rows[Majs] = new_bs

	for i := 0; i < 20; i++ {
		r := a.rows[Roots][i][0]
		new_bs[i] = [3]int{r, min3(r), fifth(r)}
	}
	a.rows[Mins] = new_bs

	for i := 0; i < 20; i++ {
		r := a.rows[Roots][i][0]
		new_bs[i] = [3]int{r, maj3(r), maj7(r)}
	}
	a.rows[Dom7s] = new_bs

	for i := 0; i < 20; i++ {
		r := a.rows[Roots][i][0]
		new_bs[i] = [3]int{min7(r), r, min3(r)}
	}
	a.rows[Dim7s] = new_bs

	// for now only handling full size
	// later pass these in
	a.numButtons = 20
	a.numRows = 6
	a.centerInd = 9
	return &a
}

func (a *accordion) searchNote(searchNote int) {

	for rowName, row := range a.rows {
		for i, ch := range row {
			for _, n := range ch {
				if n == searchNote {
					fromC := i - a.centerInd
					fmt.Printf("%s button %d\n", rowName, fromC)
					break
				}
			}
		}
	}
}

func main() {
	a := newAccordion()
	a.searchNote(0)
}
