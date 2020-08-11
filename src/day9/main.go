package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func solve(str string) (int, int) {
	escape := false
	inGarbage := false
	score := 0
	garbageCnt := 0
	level := 0

	for _, r := range str {
		s := string(r)
		switch {
		case inGarbage && escape:
			escape = false
		case inGarbage && s == ">":
			inGarbage = false
		case inGarbage && s == "!":
			escape = true
		case inGarbage:
			garbageCnt++
		case s == "<":
			inGarbage = true
		case s == "{":
			level++
		case s == "}":
			score += level
			level--
		default:
			continue
		}
	}
	return score, garbageCnt
}

func main() {
	bs, err := ioutil.ReadFile("resources/day9_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	part1, part2 := solve(string(bs))
	fmt.Printf("Part 1: %v\nPart 2: %v\n", part1, part2)
}
