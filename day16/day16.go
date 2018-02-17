package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func spin(input []byte, x int) []byte {
	if x <= 0 || x > len(input) {
		log.Fatal("index out of array bounds")
	}
	if x == len(input) {
		return input
	}
	return append(input[len(input)-x:], input[:len(input)-x]...)
}

func exchange(input []byte, a, b int) []byte {
	switch {
	case a < 0 || b < 0:
		log.Fatal("index out of array bounds")
	case a == b:
		return input
	default:
		input[a], input[b] = input[b], input[a]
	}
	return input
}

func swap(input []byte, a, b byte) []byte {
	ia, ib := -1, -1
	for i := 0; i < len(input); i++ {
		if input[i] == a {
			ia = i
		}
		if input[i] == b {
			ib = i
		}
	}
	if a < 0 || b < 0 {
		log.Fatal("not found")
	}
	return exchange(input, ia, ib)
}

func solve1(input string, cmds []string) string {
	bs := []byte(input)
	for _, cmd := range cmds {
		com := []byte(strings.TrimSpace(cmd))
		switch com[0] {
		case 's':
			i, err := strconv.Atoi(string(com[1:]))
			if err != nil {
				log.Fatal(err)
			}
			bs = spin(bs, i)
		case 'x':
			inds := strings.Split(string(com[1:]), "/")
			a, err := strconv.Atoi(string(inds[0]))
			if err != nil {
				log.Fatal(err)
			}
			b, err := strconv.Atoi(string(inds[1]))
			if err != nil {
				log.Fatal(err)
			}
			bs = exchange(bs, a, b)
		case 'p':
			bs = swap(bs, com[1], com[3])
		default:
			log.Fatal("wrong input")
		}
	}
	return string(bs)
}

func solve2(input string, cmds []string, limit int) string {
	bs, i := []byte(input), 1
	for ; i <= limit; i++ {
		bs = []byte(solve1(string(bs), cmds))
		if string(bs) == input {
			break
		}
	}
	for j := 0; j < limit%i; j++ {
		input = solve1(input, cmds)
	}
	return input
}

func main() {
	input := "abcdefghijklmnop"
	data, err := ioutil.ReadFile("day16_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	cmds := strings.Split(string(data), ",")
	part1 := solve1(input, cmds)
	part2 := solve2(input, cmds, 1000000000)
	fmt.Printf("Part 1: %v\nPart 2: %v\n", string(part1), string(part2))
}
