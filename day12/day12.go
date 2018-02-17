package main

import (
	"fmt"
	"jasilven/aoc17/util"
	"sort"
	"strings"
)

func solve(source map[string][]string) (int, int) {
	groups := make(map[string][]string)
	keys := util.GetKeys(source)
	sort.Strings(keys)
	for _, key := range keys {
		if len(source[key]) == 0 {
			continue
		}
		groups[key] = append([]string{}, source[key]...)
		groups[key] = append([]string{}, key)
		stack := append([]string{}, source[key]...)
		source[key] = []string{}
		for len(stack) > 0 {
			stack = append(stack, source[stack[0]]...)
			groups[key] = append(groups[key], source[stack[0]]...)
			source[stack[0]] = []string{}
			stack = stack[1:]
		}
		groups[key] = util.RemoveDuplicates(groups[key])
	}
	return len(groups["0"]), len(groups)
}

func main() {
	groups := make(map[string][]string)
	for _, line := range util.ReadLines("day12_input.txt") {
		line = strings.Replace(line, "<-> ", "", 1)
		line = strings.Replace(line, ",", "", -1)
		items := strings.Split(line, " ")
		groups[items[0]] = items[1:]
	}
	part1, part2 := solve(groups)
	fmt.Printf("Part 1: %v\nPart 2: %v\n", part1, part2)
}
