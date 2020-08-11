package main

import (
	"aoc-2017-go/src/util"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type program struct {
	Instructions []string
	Regs         map[string]int
	Index        int
	Result       int
	Out          chan int
	In           chan int
	Exit         bool
	Waiting      bool
}

func (p *program) done() bool {
	if p.Index >= 0 && p.Index < len(p.Instructions) && !p.Exit {
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
		case "snd":
			p.Result = p.getVal(ins[1])
			p.Index++
		case "set":
			p.Regs[ins[1]] = p.getVal(ins[2])
			p.Index++
		case "mul":
			p.Regs[ins[1]] = p.getReg(ins[1]) * p.getVal(ins[2])
			p.Index++
		case "add":
			p.Regs[ins[1]] = p.getReg(ins[1]) + p.getVal(ins[2])
			p.Index++
		case "mod":
			p.Regs[ins[1]] = p.getReg(ins[1]) % p.getVal(ins[2])
			p.Index++
		case "rcv":
			if p.getReg(ins[1]) != 0 {
				p.Exit = true
			} else {
				p.Index++
			}
		case "jgz":
			if p.getVal(ins[1]) > 0 {
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

func (p *program) run2() {
	for !p.done() {
		ins := strings.Split(p.Instructions[p.Index], " ")
		switch ins[0] {
		case "snd":
			p.Out <- p.getVal(ins[1])
			p.Result++
			p.Index++
		case "set":
			p.Regs[ins[1]] = p.getVal(ins[2])
			p.Index++
		case "mul":
			p.Regs[ins[1]] = p.getReg(ins[1]) * p.getVal(ins[2])
			p.Index++
		case "add":
			p.Regs[ins[1]] = p.getReg(ins[1]) + p.getVal(ins[2])
			p.Index++
		case "mod":
			p.Regs[ins[1]] = p.getReg(ins[1]) % p.getVal(ins[2])
			p.Index++
		case "rcv":
			p.Waiting = true
			p.Regs[ins[1]] = <-p.In
			p.Waiting = false
			p.Index++
		case "jgz":
			if p.getVal(ins[1]) > 0 {
				p.Index += p.getVal(ins[2])
			} else {
				p.Index++
			}
		default:
			log.Fatal("unknown instruction:", ins)
		}
	}
	p.Exit = true
}

func solve1(ins []string) int {
	p := program{Instructions: ins, Regs: make(map[string]int)}
	return p.run()
}

func solve2(ins []string) int {
	ch1, ch0 := make(chan int, 70), make(chan int, 70)
	p0 := program{Instructions: ins, Regs: make(map[string]int), In: ch0, Out: ch1}
	p1 := program{Instructions: ins, Regs: make(map[string]int), In: ch1, Out: ch0}
	p0.Regs["p"], p1.Regs["p"] = 0, 1
	go p0.run2()
	go p1.run2()
	for !(p0.Exit && p1.Exit) && !(p0.Waiting && p1.Waiting) {
		time.Sleep(10 * time.Millisecond)
	}
	return p1.Result
}

func main() {
	lines := util.ReadLines("resources/day18_input.txt")
	fmt.Println("Part 1:", solve1(lines))
	fmt.Println("Part 2:", solve2(lines))
}
