package main

import (
	"fmt"
	"jasilven/aoc17/util"
	"strings"
)

const input_file = "day4_input.txt"

func solve1(fname string) (result int) {
	for _, line := range util.ReadLines(fname) {
		words, valid := make(map[string]bool), true
		for _, word := range strings.Split(line, " ") {
			if words[word] {
				valid = false
				break
			} else {
				words[word] = true
			}
		}
		if valid {
			result++
		}
	}
	return
}

func solve2(fname string) (result int) {
	for _, line := range util.ReadLines(fname) {
		words, valid := make(map[string]bool), true
		for _, word := range strings.Split(line, " ") {
			if words[util.SortString(word)] {
				valid = false
				break
			} else {
				words[util.SortString(word)] = true
			}
		}
		if valid {
			result++
		}
	}
	return
}

func main() {
	part1, part2 := make(chan int, 1), make(chan int, 1)
	go func(f string) { part1 <- solve1(f) }(input_file)
	go func(f string) { part2 <- solve2(f) }(input_file)
	fmt.Printf("part 1: %v\npart 2: %v\n", <-part1, <-part2)
}
