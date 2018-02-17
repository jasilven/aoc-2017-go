package main

import (
	"fmt"
	"jasilven/aoc17/util"
	"strconv"
)

const input_file = "day1_input.txt"

func solve1(input string) (result int) {
	str := input + string(input[0])
	for i, v := range input {
		if str[i] == str[i+1] {
			num, _ := strconv.Atoi(string(v))
			result += num
		}
	}
	return
}

func solve2(input string) (result int) {
	len := len(input)
	for i, v := range input {
		if input[i] == input[(i+len/2)%len] {
			num, _ := strconv.Atoi(string(v))
			result += num
		}
	}
	return
}

func main() {
	data := util.ReadLines(input_file)
	part1, part2 := make(chan int, 1), make(chan int, 1)
	go func(s string) { part1 <- solve1(s) }(data[0])
	go func(s string) { part2 <- solve2(s) }(data[0])
	fmt.Printf("part 1: %v\npart 2: %v\n", <-part1, <-part2)
}
