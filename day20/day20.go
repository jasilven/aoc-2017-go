package main

import (
	"fmt"
	"jasilven/aoc-2017-go/util"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

const aSize = 9

func loadData(fname string) [][aSize]int {
	result := [][aSize]int{}
	for _, line := range util.ReadLines(fname) {
		line = line[3 : len(line)-1]
		line = strings.Replace(line, ">, v=<", ",", 1)
		line = strings.Replace(line, ">, a=<", ",", 1)
		nums := strings.Split(line, ",")
		var fs [aSize]int
		for k, n := range nums {
			i, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			fs[k] = i
		}
		result = append(result, fs)
	}
	return result
}

func solve1(data [][aSize]int) int {
	type dva struct{ d, v, a, id int }
	dvas := []dva{}
	for i, ia := range data {
		fs := []float64{}
		for j := 0; j < aSize; j++ {
			fs = append(fs, float64(ia[j]))
		}
		d := int(math.Abs(fs[0]) + math.Abs(fs[1]) + math.Abs(fs[2]))
		v := int(math.Sqrt(fs[3]*fs[3]) + math.Sqrt(fs[4]*fs[4]) + math.Sqrt(fs[5]*fs[5]))
		a := int(math.Sqrt(fs[6]*fs[6]) + math.Sqrt(fs[7]*fs[7]) + math.Sqrt(fs[8]*fs[8]))
		dvas = append(dvas, dva{d: d, v: v, a: a, id: i})
	}
	sort.Slice(dvas, func(i, j int) bool {
		switch {
		case dvas[i].a < dvas[j].a:
			return true
		case dvas[i].a == dvas[j].a && dvas[i].v < dvas[j].v:
			return true
		case dvas[i].a == dvas[j].a && dvas[i].v == dvas[j].v && dvas[i].d < dvas[j].d:
			return true
		default:
			return false
		}
	})
	return dvas[0].id
}

func solve2(data [][aSize]int) int {
	pmap := make(map[string]int)
	for i := 0; i < len(data); i++ {
		pmap[fmt.Sprintf("%d.%d.%d", data[i][0], data[i][1], data[i][2])] = i
	}
	for i := 0; i < 1000; i++ {
		for _, val := range pmap {
			data[val][3] += data[val][6]
			data[val][4] += data[val][7]
			data[val][5] += data[val][8]
			data[val][0] += data[val][3]
			data[val][1] += data[val][4]
			data[val][2] += data[val][5]
		}
		nmap, collides := make(map[string]int), []string{}
		for _, v := range pmap {
			key := fmt.Sprintf("%d.%d.%d", data[v][0], data[v][1], data[v][2])
			_, ok := nmap[key]
			if ok {
				collides = append(collides, key)
			} else {
				nmap[key] = v
			}
		}
		for _, val := range collides {
			delete(nmap, val)
		}
		pmap = nmap
	}
	return len(pmap)
}

func main() {
	data := loadData("day20_input.txt")
	fmt.Println(solve1(data))
	fmt.Println(solve2(data))
}
