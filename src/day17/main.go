package main

import "fmt"

const stepCnt = 355

type cbuffer struct {
	Buffer  []int
	Index   int
	StepCnt int
	Count   int
}

func NewCbuffer() cbuffer {
	return cbuffer{Buffer: []int{0}, Index: 0, StepCnt: stepCnt, Count: 1}
}

func (cb *cbuffer) insert(val int) {
	cb.Index = (cb.Index + cb.StepCnt + 1) % len(cb.Buffer)
	head := cb.Buffer[:cb.Index]
	tail := make([]int, len(cb.Buffer[cb.Index:]))
	copy(tail, cb.Buffer[cb.Index:])
	head = append(head, val)
	head = append(head, tail...)
	cb.Buffer = head
}

func (cb *cbuffer) getNext() int {
	return cb.Buffer[(cb.Index+1)%len(cb.Buffer)]
}

func solve1(buf cbuffer, count int) int {
	for i := 1; i <= count; i++ {
		buf.insert(i)
	}
	return buf.getNext()
}

func solve2(count int) (result int) {
	index := 0
	for i := 1; i <= count; i++ {
		index = (index + stepCnt + 1) % i
		if index == 0 {
			result = i
		}
	}
	return
}

func main() {
	fmt.Printf("Part 1: %v\n", solve1(NewCbuffer(), 2017))
	fmt.Printf("Part 2: %v\n", solve2(50000000))
}
