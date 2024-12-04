package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/duluk/AoC/2024/pkg/day3"
)

func main() {
	const in_file = "/home/jab3/src/mine/AoC/2024/pkg/day3/day3.input"
	fd, err := os.Open(in_file)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	var sum int
	scanner := bufio.NewScanner(fd)
	fmt.Printf("Day 3_1. ")
	sum, err = Day3.SolveDay3_1(scanner)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Sum of products: %d\n", sum)

	fmt.Printf("Day 3_2. ")
	fd.Seek(0, 0)
	scanner = bufio.NewScanner(fd)
	sum, err = Day3.SolveDay3_2(scanner)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Sum of products: %d\n", sum)
}
