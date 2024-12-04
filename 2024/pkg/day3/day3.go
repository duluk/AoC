package Day3

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func addMul(line string) (int, error) {
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

	var sum int
	var matches = re.FindAllString(line, -1)
	for _, match := range matches {
		parts := strings.Split(match, ",")

		// Parts Ex: [mul(123 456)]
		num1, err := strconv.Atoi(parts[0][4:])
		if err != nil {
			return 0, err
		}
		num2, err := strconv.Atoi(parts[1][:len(parts[1])-1])
		if err != nil {
			return 0, err
		}

		product := num1 * num2
		sum += product
	}

	return sum, nil
}

func SolveDay3_1(scanner *bufio.Scanner) (int, error) {
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		s, err := addMul(line)
		if err != nil {
			return 0, err
		}
		sum += s
	}

	return sum, nil
}

// I cheated and found how someone else did it. In Python no less.
// https://www.youtube.com/watch?v=uBup4-4uPBI, Neil Thistlethwaite
// I didn't realize 'FindAllString' would return the matches line by line in
// order. Pretty clever. (I don't know what I thought it did but...)
func SolveDay3_2(scanner *bufio.Scanner) (int, error) {
	var s string
	for scanner.Scan() {
		s += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	matches := re.FindAllString(s, -1)

	sum := 0
	good := true
	for _, match := range matches {
		switch match {
		case "do()":
			good = true
		case "don't()":
			good = false
		default:
			if good {
				s, _ := addMul(match)
				sum += s
			}
		}
	}

	return sum, nil
}
