package main

import (
	"fmt"

	"github.com/duluk/AoC/2024/pkg/day2"
)

func main() {
	fmt.Printf("Day 2_1. ")
	safe, err := Day2.SolveDay2_1()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Safe: %d\n", safe)

	fmt.Printf("Day 2_2. ")
	safe, err = Day2.SolveDay2_2()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Safe: %d\n", safe)
}
