package hash

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const Lsize = 256

func reverse(list *[Lsize]int, start int, lenght int) {
	var subl []int
	for i, j := start, 0; j < lenght; i, j = (i+1)%Lsize, j+1 {
		subl = append(subl, list[i])
	}
	for i, j := start, len(subl)-1; j >= 0; i, j = (i+1)%Lsize, j-1 {
		list[i] = subl[j]
	}
}

func KnotHashSparse(lens []int, rounds int) [Lsize]int {
	var list [Lsize]int
	for i := range list {
		list[i] = i
	}
	var skip, curpos int
	for r := 0; r < rounds; r++ {
		for _, len := range lens {
			reverse(&list, curpos, len)
			curpos = (curpos + int(len) + skip) % Lsize
			skip++
		}
	}
	return list
}

func KnotHash(input string) (result string) {
	bsize := 16
	lens := []int{}
	for _, b := range input {
		lens = append(lens, int(b))
	}
	for _, str := range strings.Split("17, 31, 73, 47, 23", ", ") {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		lens = append(lens, num)
	}
	sparse := KnotHashSparse(lens, 64)
	dense := []int{}
	for n := 0; n < bsize; n++ {
		b := sparse[n*bsize : n*bsize+bsize]
		dense = append(dense, b[0]^b[1]^b[2]^b[3]^b[4]^b[5]^b[6]^b[7]^b[8]^b[9]^b[10]^b[11]^b[12]^b[13]^b[14]^b[15])
	}
	for _, n := range dense {
		result = result + fmt.Sprintf("%0.2x", n)
	}
	return result
}
