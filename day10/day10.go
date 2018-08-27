package main

import (
	"fmt"
	"jasilven/aoc-2017-go/hash"
	"log"
	"strconv"
	"strings"
)

func parseInts(input, sep string) []int {
	result := []int{}
	for _, val := range strings.Split(input, sep) {
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, int(num))
	}
	return result
}

func solve1(lens []int) int {
	h := hash.KnotHashSparse(lens, 1)
	return h[0] * h[1]
}

func solve2(input string) (result string) {
	return hash.KnotHash(input)
}

func main() {
	input := "197,97,204,108,1,29,5,71,0,50,2,255,248,78,254,63"
	part1, part2 := make(chan int, 1), make(chan string, 1)
	go func(lens []int) { part1 <- solve1(lens) }(parseInts(input, ","))
	go func(i string) { part2 <- solve2(i) }(input)
	fmt.Printf("part 1: %v\npart 2: %v\n", <-part1, <-part2)
}
