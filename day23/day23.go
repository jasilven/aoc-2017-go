package main

import (
	"fmt"
	"jasilven/aoc-2017-go/util"
	"log"
	"strconv"
	"strings"
)

type program struct {
	Instructions []string
	Regs         map[string]int
	Index        int
	Result       int
}

func (p *program) done() bool {
	if p.Index >= 0 && p.Index < len(p.Instructions) {
		return false
	}
	return true
}

func (p *program) getReg(r string) int {
	v, ok := p.Regs[r]
	if ok {
		return v
	}
	return 0
}

func (p *program) getVal(s string) (result int) {
	result, err := strconv.Atoi(s)
	if err != nil {
		return p.getReg(s)
	}
	return
}

func (p *program) run() int {
	for !p.done() {
		ins := strings.Split(p.Instructions[p.Index], " ")
		switch ins[0] {
		case "set":
			p.Regs[ins[1]] = p.getVal(ins[2])
			p.Index++
		case "mul":
			p.Regs[ins[1]] = p.getReg(ins[1]) * p.getVal(ins[2])
			p.Result++
			p.Index++
		case "sub":
			p.Regs[ins[1]] = p.getReg(ins[1]) - p.getVal(ins[2])
			p.Index++
		case "jnz":
			if p.getVal(ins[1]) != 0 {
				p.Index += p.getVal(ins[2])
			} else {
				p.Index++
			}
		default:
			log.Fatal("unknown instruction:", ins)
		}
	}
	return p.Result
}

func solve1(instructions []string) int {
	prog := program{Instructions: instructions, Regs: make(map[string]int)}
	return prog.run()
}

func solve2() (result int) {
	for i := 109900; i < 126901; i += 17 {
		if !util.IsPrime(i) {
			result++
		}
	}
	return
}

func main() {
	fmt.Println("Part 1: ", solve1(util.ReadLines("day23_input.txt")))
	fmt.Println("Part 2: ", solve2())
}
