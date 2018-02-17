package main

import (
	"fmt"
	"jasilven/aoc17/util"
	"sort"
	"strconv"
	"strings"
)

const InputFile = "day2_input.txt"

func solve1(fname string) (result int) {
	for _, line := range util.ReadLines(fname) {
		var nums []int
		for _, s := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(s)
			nums = append(nums, num)
		}
		sort.Ints(nums)
		result += nums[len(nums)-1] - nums[0]
	}
	return
}

func findResult(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums); j++ {
			if (nums[j]%nums[i] == 0) && (i != j) {
				return nums[j] / nums[i]
			}
		}
	}
	return 0
}

func solve2(fname string) (result int) {
	for _, line := range util.ReadLines(fname) {
		var nums []int
		for _, s := range strings.Split(line, ",") {
			num, err := strconv.Atoi(s)
			if err == nil {
				nums = append(nums, num)
			}
		}
		sort.Ints(nums)
		result += findResult(nums)
	}
	return
}

func main() {
	part1, part2 := make(chan int, 1), make(chan int, 1)
	go func(f string) { part1 <- solve1(f) }(InputFile)
	go func(f string) { part2 <- solve2(f) }(InputFile)
	fmt.Printf("part 1: %v\npart 2: %v\n", <-part1, <-part2)
}
