package main

import (
	"fmt"
)

func generateNum(seed, factor, cnt uint64, out chan<- uint64, valid func(uint64) bool) {
	result := seed
	for i := 0; i < int(cnt); {
		result = (result * factor) % 2147483647
		if valid(result) {
			out <- result
			i++
		}
	}
	close(out)
}

func solve(count int, fva, fvb func(i uint64) bool) (result int) {
	bufsize := 30
	cha, chb := make(chan uint64, bufsize), make(chan uint64, bufsize)
	go generateNum(277, 16807, uint64(count), cha, fva)
	go generateNum(349, 48271, uint64(count), chb, fvb)
	for i := 0; i < count; i++ {
		if uint16(<-cha) == uint16(<-chb) {
			result++
		}
	}
	return
}

func fvalid1(i uint64) bool { return true }

func fvalid2a(i uint64) bool { return i%4 == 0 }

func fvalid2b(i uint64) bool { return i%8 == 0 }

func main() {
	fmt.Println("part 1:", solve(40000000, fvalid1, fvalid1))
	fmt.Println("part 2:", solve(5000000, fvalid2a, fvalid2b))
}
