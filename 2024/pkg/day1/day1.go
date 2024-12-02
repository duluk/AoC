package Day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func get_input() ([]int, []int) {
	const in_file = "/home/jab3/src/mine/AoC/2024/pkg/day1/day1.input"
	fd, err := os.Open(in_file)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, nil
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	var first, second []int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		n1, _ := strconv.Atoi(parts[0])
		n2, _ := strconv.Atoi(parts[1])

		first = append(first, n1)
		second = append(second, n2)
	}

	return first, second
}

func SolveDay1_1() error {
	first, second := get_input()
	if first == nil || second == nil {
		return fmt.Errorf("Error reading input")
	}

	sort.Ints(first)
	sort.Ints(second)

	var sum int
	for i := 0; i < len(first); i++ {
		sum += abs(first[i] - second[i])
	}

	fmt.Printf("Sum of differences: %d\n", sum)

	return nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SolveDay1_2() error {
	// My idea: take the right list and turn it into a set/hashmap whih has the
	// number as the key and the value is how many times it appears in the
	// list. Then go through the first list and multiply each number by the
	// value of that number's key in the hashmap.

	first, second := get_input()
	if first == nil || second == nil {
		return fmt.Errorf("Error reading input")
	}

	second_map := make(map[int]int)
	for _, n := range second {
		second_map[n]++
	}

	var sum int
	for _, n := range first {
		sum += n * second_map[n]
	}

	fmt.Printf("Sum of products: %d\n", sum)

	return nil
}
