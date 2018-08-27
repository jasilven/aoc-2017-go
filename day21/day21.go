package main

import (
	"fmt"
	"jasilven/aoc-2017-go/util"
	"strings"
)

type matrix [][]byte

func (m matrix) String() (result string) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			result += string(m[i][j])
		}
		result += "/"
	}
	return strings.TrimSuffix(result, "/")
}

func (m matrix) flip() matrix {
	result := matrix{}
	for i := len(m) - 1; i >= 0; i-- {
		row := []byte{}
		for j := 0; j < len(m[0]); j++ {
			row = append(row, m[i][j])
		}
		result = append(result, row)
	}
	return result
}

func (m matrix) rotate() matrix {
	result := matrix{}
	for j := 0; j < len(m[0]); j++ {
		row := []byte{}
		for i := len(m) - 1; i >= 0; i-- {
			row = append(row, m[i][j])
		}
		result = append(result, row)
	}
	return result
}

func newMatrix(s string) matrix {
	result := matrix{}
	for _, row := range strings.Split(s, "/") {
		bs := []byte{}
		for j, _ := range row {
			bs = append(bs, row[j])
		}
		result = append(result, bs)
	}
	return result
}

func loadRules(fname string) map[string]string {
	result := make(map[string]string)
	for _, line := range util.ReadLines(fname) {
		ss := strings.Split(line, " => ")
		result[ss[0]] = ss[1]
		m := newMatrix(ss[0])
		result[m.rotate().String()] = ss[1]
		result[m.rotate().rotate().String()] = ss[1]
		result[m.rotate().rotate().rotate().String()] = ss[1]
		result[m.flip().String()] = ss[1]
		result[m.rotate().flip().String()] = ss[1]
		result[m.rotate().rotate().flip().String()] = ss[1]
		result[m.rotate().rotate().rotate().flip().String()] = ss[1]
	}
	return result
}

func solve(input string, rules map[string]string, cnt int) int {
	for c := 0; c < cnt; c++ {
		rows, result, sqSize := strings.Split(input, "/"), "", 3
		if len(rows[0])%2 == 0 {
			sqSize = 2
		}
		sqCnt := len(rows[0]) / sqSize
		for i := 0; i < sqCnt; i++ {
			squares, square := []string{}, ""
			for j := 0; j < sqCnt; j++ {
				if sqSize == 2 {
					square = rows[i*sqSize][j*sqSize:j*sqSize+sqSize] + "/" +
						rows[i*sqSize+1][j*sqSize:j*sqSize+sqSize]
				} else {
					square = rows[i*sqSize][j*sqSize:j*sqSize+sqSize] + "/" +
						rows[i*sqSize+1][j*sqSize:j*sqSize+sqSize] + "/" +
						rows[i*sqSize+2][j*sqSize:j*sqSize+sqSize]
				}
				squares = append(squares, rules[square])
			}
			for l := 0; l < sqSize+1; l++ {
				for _, sq := range squares {
					lines := strings.Split(sq, "/")
					result += lines[l]
				}
				result += "/"
			}
		}
		input = strings.TrimSuffix(result, "/")
	}
	return strings.Count(input, "#")
}

func main() {
	rules := loadRules("day21_input.txt")
	fmt.Println("Part 1:", solve(".#./..#/###", rules, 5))
	fmt.Println("Part 2:", solve(".#./..#/###", rules, 18))
}
