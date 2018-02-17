package main

import (
	"fmt"
	"jasilven/aoc17/util"
	"strconv"
	"strings"
)

type Program struct {
	Weight      int
	Children    []string
	TotalWeight int
}

type Programs map[string]Program

func loadData(fname string) Programs {
	result := make(map[string]Program)
	for _, line := range util.ReadLines(fname) {
		line = strings.Replace(line, "(", "", -1)
		line = strings.Replace(line, ")", "", -1)
		line = strings.Replace(line, "-> ", "", -1)
		line = strings.Replace(line, ",", "", -1)
		words := strings.Split(line, " ")
		weight, err := strconv.Atoi(words[1])
		if err != nil {
			panic(err)
		}
		result[words[0]] = Program{weight, words[2:], 0}
	}
	return calcTotals(result)
}

func calcTotals(progs Programs) Programs {
	nodes := []string{solve1(progs)}
	for len(nodes) > 0 {
		cur := nodes[len(nodes)-1]
		childs := progs[cur].Children
		if len(childs) > 0 && progs[childs[0]].TotalWeight > 0 {
			prog := progs[cur]
			prog.TotalWeight = prog.Weight
			for _, c := range childs {
				prog.TotalWeight += progs[c].TotalWeight
			}
			progs[cur] = prog
			nodes = nodes[:len(nodes)-1]
		} else if len(childs) > 0 {
			nodes = append(nodes, childs...)
		} else {
			prog := progs[cur]
			prog.TotalWeight = prog.Weight
			progs[cur] = prog
			nodes = nodes[:len(nodes)-1]
		}
	}
	return progs
}

func solve1(progs Programs) (result string) {
	parents, childs := []string{}, []string{}
	for name, prog := range progs {
		if len(prog.Children) > 0 {
			parents = append(parents, name)
			childs = append(childs, progs[name].Children...)
		}
	}
	for _, parent := range parents {
		if util.Contains(childs, parent) {
			continue
		} else {
			result = parent
			break
		}
	}
	return
}

func nextRoot(progs Programs, root string) string {
	freqs := make(map[int][]string)
	for _, name := range progs[root].Children {
		w := progs[name].TotalWeight
		freqs[w] = append(freqs[w], name)
	}
	for _, names := range freqs {
		if len(names) == 1 {
			return names[0]
		}
	}
	return root
}

func calcDelta(progs Programs, root string) int {
	freqs := make(map[int][]string)
	var wrong, right int
	for _, name := range progs[root].Children {
		w := progs[name].TotalWeight
		freqs[w] = append(freqs[w], name)
	}
	for w, names := range freqs {
		if len(names) == 1 {
			wrong = w
		} else {
			right = w
		}
	}
	return right - wrong
}

func solve2(progs Programs) int {
	cur := solve1(progs)
	delta := calcDelta(progs, cur)
	next := nextRoot(progs, cur)
	for next != cur {
		cur = next
		next = nextRoot(progs, cur)
	}
	return progs[cur].Weight + delta
}

func main() {
	progs := loadData("day7_input.txt")
	part1, part2 := make(chan string, 1), make(chan int, 1)
	go func(p Programs) { part1 <- solve1(p) }(progs)
	go func(p Programs) { part2 <- solve2(p) }(progs)
	fmt.Printf("part 1: %v\npart 2: %v\n", <-part1, <-part2)
}
