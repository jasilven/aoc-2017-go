package main

import (
	"fmt"
	"jasilven/aoc-2017-go/util"
	"log"
	"strconv"
	"strings"
)

type comp struct {
	a, b int
}

type bridge struct {
	comps map[comp]bool
	end   int
}

func loadComps(fname string) map[int][]comp {
	comps := make(map[int][]comp)
	for _, line := range util.ReadLines(fname) {
		line = strings.TrimSpace(line)
		ports := strings.Split(line, "/")
		a, err := strconv.Atoi(ports[0])
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.Atoi(ports[1])
		if err != nil {
			log.Fatal(err)
		}
		comps[a] = append(comps[a], comp{a: a, b: b})
		if a != b {
			comps[b] = append(comps[b], comp{a: a, b: b})
		}
	}
	return comps
}

func findBridges(bridges []bridge, components map[int][]comp) []bridge {
	result := []bridge{}
	for {
		if len(bridges) > 0 {
			b, comps := bridges[len(bridges)-1], []comp{}
			bridges = bridges[:len(bridges)-1]
			for _, cand := range components[b.end] {
				_, ok := b.comps[cand]
				if ok {
					continue
				} else {
					comps = append(comps, cand)
				}
			}
			if len(comps) == 0 {
				result = append(result, b)
			} else {
				for _, c := range comps {
					var end int
					if c.a == b.end {
						end = c.b
					} else {
						end = c.a
					}
					newb := bridge{end: end}
					newb.comps = make(map[comp]bool)
					newb.comps[c] = true
					for oldc := range b.comps {
						newb.comps[oldc] = true
					}
					bridges = append(bridges, newb)
				}
			}
		} else {
			break
		}
	}
	return result
}

func solve(bridges []bridge) (int, int) {
	strongest1, strongest2, longest := 0, 0, 0
	for _, bridge := range bridges {
		strenght, lenght := 0, 0
		for component := range bridge.comps {
			strenght += component.a + component.b
			lenght++
		}
		if lenght >= longest {
			longest = lenght
			if strenght >= strongest2 {
				strongest2 = strenght
			}
		}
		if strenght > strongest1 {
			strongest1 = strenght
		}
	}
	return strongest1, strongest2
}

func main() {
	comps := loadComps("day24_input.txt")
	roots, bridges := comps[0], []bridge{}
	if len(roots) == 0 {
		log.Fatal("root components not found")
	}
	for _, root := range roots {
		var end int
		if root.a == 0 {
			end = root.b
		} else {
			end = root.a
		}
		nbridge := bridge{end: end}
		nbridge.comps = make(map[comp]bool)
		nbridge.comps[root] = true
		bridges = append(bridges, nbridge)
	}
	part1, part2 := solve(findBridges(bridges, comps))
	fmt.Printf("Part 1: %v\nPart 2: %v\n", part1, part2)
}
