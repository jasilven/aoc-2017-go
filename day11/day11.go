package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

type vertex struct {
	X, Y, Z int
}

func (v *vertex) distance() int {
	result := (math.Abs(float64(v.X)) + math.Abs(float64(v.Y)) + math.Abs(float64(v.Z))) / 2
	return (int(result))
}

func (v *vertex) n() {
	v.Y++
	v.Z--
}

func (v *vertex) s() {
	v.Y--
	v.Z++
}

func (v *vertex) ne() {
	v.X++
	v.Z--
}

func (v *vertex) sw() {
	v.X--
	v.Z++
}

func (v *vertex) nw() {
	v.X--
	v.Y++
}

func (v *vertex) se() {
	v.X++
	v.Y--
}

func solve(steps []string) (int, int) {
	v := vertex{}
	maxDistance := 0
	for _, step := range steps {
		switch step {
		case "n":
			v.n()
		case "s":
			v.s()
		case "ne":
			v.ne()
		case "nw":
			v.nw()
		case "se":
			v.se()
		case "sw":
			v.sw()
		default:
			log.Fatal("unknown step:", step)
		}
		d := v.distance()
		if d > maxDistance {
			maxDistance = d
		}
	}
	return v.distance(), maxDistance
}

func main() {
	bs, err := ioutil.ReadFile("day11_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	steps := strings.Split(strings.TrimSpace(string(bs)), ",")
	part1, part2 := solve(steps)
	fmt.Printf("part 1: %v\npart 2: %v\n", part1, part2)
}
