package main

import (
	"fmt"

	"github.com/duluk/AoC/2024/pkg/day1"
)

func main() {
	fmt.Printf("Day 1_1. ")
	sum, err := Day1.SolveDay1_1()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Sum of differences: %d\n", sum)

	fmt.Printf("Day 1_2. ")
	sum, err = Day1.SolveDay1_2()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Sum of products: %d\n", sum)
}
