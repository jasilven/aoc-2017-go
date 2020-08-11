package main

import (
	"aoc-2017-go/src/util"
	"fmt"
	"log"
)

type Instruction struct {
	Reg       string
	Com       string
	ComArg    int
	TestLeft  string
	TestOp    string
	TestRight int
}

type Program struct {
	Regs         map[string]int
	Instructions []Instruction
	MaxReg       int
}

func loadData(fname string) Program {
	result := Program{}
	result.Regs = make(map[string]int)
	for _, line := range util.ReadLines(fname) {
		ins := Instruction{}
		var x string
		_, err := fmt.Sscanf(line, "%s%s%d%s%s%s%d", &ins.Reg, &ins.Com, &ins.ComArg, &x, &ins.TestLeft, &ins.TestOp, &ins.TestRight)
		if err != nil {
			log.Fatal(err)
		}
		result.Regs[ins.Reg] = 0
		result.Instructions = append(result.Instructions, ins)
	}
	return result
}

func evalTest(i int, prog Program) bool {
	ins := prog.Instructions[i]
	switch ins.TestOp {
	case "==":
		return prog.Regs[ins.TestLeft] == ins.TestRight
	case "!=":
		return prog.Regs[ins.TestLeft] != ins.TestRight
	case "<":
		return prog.Regs[ins.TestLeft] < ins.TestRight
	case ">":
		return prog.Regs[ins.TestLeft] > ins.TestRight
	case "<=":
		return prog.Regs[ins.TestLeft] <= ins.TestRight
	case ">=":
		return prog.Regs[ins.TestLeft] >= ins.TestRight
	default:
		log.Fatal("unknown operator:", ins.TestOp)
	}
	return false
}

func doCmd(i int, prog *Program) {
	ins := prog.Instructions[i]
	switch ins.Com {
	case "inc":
		prog.Regs[ins.Reg] = prog.Regs[ins.Reg] + ins.ComArg
	case "dec":
		prog.Regs[ins.Reg] = prog.Regs[ins.Reg] - ins.ComArg
	default:
		log.Fatal("unknown command:", ins.Com)
	}
	if prog.Regs[ins.Reg] > prog.MaxReg {
		prog.MaxReg = prog.Regs[ins.Reg]
	}
}

func solve(prog Program) (int, int) {
	part1 := 0
	for i, _ := range prog.Instructions {
		if evalTest(i, prog) {
			doCmd(i, &prog)
		}
	}
	for _, val := range prog.Regs {
		if val > part1 {
			part1 = val
		}
	}
	return part1, prog.MaxReg
}

func main() {
	part1, part2 := solve(loadData("resources/day8_input.txt"))
	fmt.Printf("Part 1: %v\nPart 2: %v\n", part1, part2)
}
