package main

import (
	"fmt"
	"strconv"
)

type rule struct {
	value, move int
	nextState   string
}

type turing struct {
	tape   map[int]int
	rules  map[string]rule
	cursor int
	state  string
}

func (t *turing) execute() {
	rule := t.rules[t.state]
	if rule.value == 1 {
		t.tape[t.cursor] = 1
	} else {
		delete(t.tape, t.cursor)
	}
	t.cursor += rule.move
	val, ok := t.tape[t.cursor]
	if !ok {
		val = 0
	}
	t.state = rule.nextState + strconv.Itoa(val)
}

func (t *turing) checksum() int {
	return len(t.tape)
}

func main() {
	rules := make(map[string]rule)
	rules["A0"] = rule{value: 1, move: 1, nextState: "B"}
	rules["A1"] = rule{value: 0, move: -1, nextState: "C"}
	rules["B0"] = rule{value: 1, move: -1, nextState: "A"}
	rules["B1"] = rule{value: 1, move: 1, nextState: "C"}
	rules["C0"] = rule{value: 1, move: 1, nextState: "A"}
	rules["C1"] = rule{value: 0, move: -1, nextState: "D"}
	rules["D0"] = rule{value: 1, move: -1, nextState: "E"}
	rules["D1"] = rule{value: 1, move: -1, nextState: "C"}
	rules["E0"] = rule{value: 1, move: 1, nextState: "F"}
	rules["E1"] = rule{value: 1, move: 1, nextState: "A"}
	rules["F0"] = rule{value: 1, move: 1, nextState: "A"}
	rules["F1"] = rule{value: 1, move: 1, nextState: "E"}
	t := turing{tape: make(map[int]int), rules: rules, state: "A0"}
	for i := 0; i < 12134527; i++ {
		t.execute()
	}
	fmt.Println("Part 1:", t.checksum())
}
