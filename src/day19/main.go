package main

import (
	"aoc-2017-go/src/util"
	"fmt"
	"strings"
)

type walker struct {
	roadmap               []string
	line, pos, dir, count int
	result                string
}

func (w *walker) turn() bool {
	switch w.dir {
	case 0, 2:
		if w.pos > 0 && w.roadmap[w.line][w.pos-1] != ' ' {
			w.pos--
			w.dir = 1
		} else if (w.pos+1) < len(w.roadmap[w.line]) && w.roadmap[w.line][w.pos+1] != ' ' {
			w.pos++
			w.dir = 3
		} else {
			return false
		}
	case 1, 3:
		if w.line > 0 && w.roadmap[w.line-1][w.pos] != ' ' {
			w.line--
			w.dir = 2
		} else if len(w.roadmap) > (w.line+1) && w.roadmap[w.line+1][w.pos] != ' ' {
			w.line++
			w.dir = 0
		} else {
			return false
		}
	}
	return true
}

func (w *walker) walk() bool {
	ok := true
	if util.IsLetter(string(w.roadmap[w.line][w.pos])) {
		w.result += string(w.roadmap[w.line][w.pos])
	}
	switch w.dir {
	case 0:
		if len(w.roadmap) > (w.line+1) && w.roadmap[w.line+1][w.pos] != ' ' {
			w.line++
		} else {
			ok = w.turn()
		}
	case 1:
		if w.pos > 0 && w.roadmap[w.line][w.pos-1] != ' ' {
			w.pos--
		} else {
			ok = w.turn()
		}
	case 2:
		if w.line > 0 && w.roadmap[w.line-1][w.pos] != ' ' {
			w.line--
		} else {
			ok = w.turn()
		}
	case 3:
		if w.pos < (len(w.roadmap[w.line])-1) && w.roadmap[w.line][w.pos+1] != ' ' {
			w.pos++
		} else {
			ok = w.turn()
		}
	}
	w.count++
	return ok
}

func solve(rm []string, line, p int) (string, int) {
	w := walker{roadmap: rm, pos: p}
	for w.walk() {
	}
	return w.result, w.count
}

func main() {
	lines := util.ReadLines("resources/day19_input.txt")
	part1, part2 := solve(lines, 0, strings.Index(lines[0], "|"))
	fmt.Printf("Part 1: %v\nPart 2: %v\n", part1, part2)
}
