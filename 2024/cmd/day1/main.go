package main

import (
	"fmt"

	"github.com/duluk/AoC/2024/pkg/day1"
)

func main() {
	fmt.Printf("Day 1_1. ")
	err := Day1.SolveDay1_1()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Day 1_2. ")
	err = Day1.SolveDay1_2()
	if err != nil {
		panic(err)
	}
}
