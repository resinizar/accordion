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

type accordion struct {
	roots [20]int
	maj3s [20]int
	majs  [20][3]int
	mins  [20][3]int
	dom7s [20][3]int
	dim7s [20][3]int
}

func newAccordion() *accordion {
	a := accordion{}

	for i := range a.roots {
		if i == 0 {
			a.roots[i] = 9 // bottom note
		} else {
			a.roots[i] = fifth(a.roots[i-1])
		}
	}

	for i := range a.maj3s {
		r := a.roots[i]
		a.maj3s[i] = maj3(r)
	}

	for i := range a.majs {
		r := a.roots[i]
		a.majs[i] = [3]int{r, a.maj3s[i], fifth(r)}
	}

	for i := range a.mins {
		r := a.roots[i]
		a.mins[i] = [3]int{r, min3(r), fifth(r)}
	}

	for i := range a.dom7s {
		r := a.roots[i]
		a.dom7s[i] = [3]int{r, maj3(r), maj7(r)}
	}

	for i := range a.dim7s {
		r := a.roots[i]
		a.dim7s[i] = [3]int{min7(r), r, min3(r)}
	}

	return &a
}

func main() {
	a := newAccordion()

	fmt.Println(a.roots)
	fmt.Println(a.maj3s)
	fmt.Println(a.majs)
	fmt.Println(a.mins)
	fmt.Println(a.dom7s)
	fmt.Println(a.dim7s)
}
