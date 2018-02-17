package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const bcount = 16
const input_file = "day6_input.txt"

func readBanks(fname string) (result [bcount]int) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	for i, bank := range strings.Split(strings.TrimSpace(string(data)), "\t") {
		j, err := strconv.Atoi(bank)
		if err != nil {
			log.Fatal(err)
		}
		result[i] = j
	}
	return
}

func largestBank(banks [bcount]int) (index, value int) {
	for i, v := range banks {
		if v > value {
			index = i
			value = v
		}
	}
	return
}

func contains(col [][bcount]int, item [bcount]int) bool {
	for _, v := range col {
		if v == item {
			return true
		}
	}
	return false
}

func reallocate(banks [bcount]int) [bcount]int {
	index, blocks := largestBank(banks)
	banks[index] = 0
	for i, count := (index+1)%bcount, 1; count <= blocks; count, i = count+1, (i+1)%bcount {
		banks[i] = banks[i] + 1
	}
	return banks
}

func solve1(fname string) (cycles int) {
	banks := readBanks(fname)
	seen := [][bcount]int{banks}
	for {
		banks, cycles = reallocate(banks), cycles+1
		if contains(seen, banks) {
			break
		}
		seen = append(seen, banks)
	}
	return
}

func solve2(fname string) (result int) {
	banks := readBanks(fname)
	seen := [][bcount]int{banks}
	seenIndex := make(map[string]int)
	for {
		seenIndex[fmt.Sprintf("%v", banks)] = result
		banks, result = reallocate(banks), result+1
		if contains(seen, banks) {
			result = result - seenIndex[fmt.Sprintf("%v", banks)]
			break
		}
		seen = append(seen, banks)
	}
	return
}

func main() {
	part1, part2 := make(chan int, 1), make(chan int, 1)
	go func(f string) { part1 <- solve1(f) }(input_file)
	go func(f string) { part2 <- solve2(f) }(input_file)
	fmt.Printf("part 1: %v\npart 2: %v\n", <-part1, <-part2)
}
