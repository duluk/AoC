package Day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func SolveDay2_1() (int, error) {
	const in_file = "/home/jab3/src/mine/AoC/2024/pkg/day2/day2.input"
	fd, err := os.Open(in_file)
	if err != nil {
		return 0, err
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	var safe int
	for scanner.Scan() {
		line := scanner.Text()

		nums := strings.Fields(line)
		if isValid(nums) {
			safe++
		}
	}

	return safe, nil
}

func SolveDay2_2() (int, error) {
	const in_file = "/home/jab3/src/mine/AoC/2024/pkg/day2/day2.input"
	fd, err := os.Open(in_file)
	if err != nil {
		return 0, err
	}
	defer fd.Close()

	var safe int
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)

		// Have to create a new slice to avoid modifying the original (see
		// block comment below). There's probably a better way to do
		// this...still learning Go.
		new_slice := make([]string, 0, len(nums)-1)
		for i := 0; i < len(nums); i++ {
			// Clear the slice by setting the length to 0, causing appends to
			// start at beginning, overwriting old data.
			new_slice = new_slice[:0]
			new_slice = append(new_slice, nums[:i]...)
			new_slice = append(new_slice, nums[i+1:]...)

			if isValid(new_slice) {
				safe++
				break
			}
		}
	}

	return safe, nil
}

/*

This was my first attempt. It's not working because it is modifying the
underlying nums array (ie, s1 := nums[:i]). The solution was to createa a new
slice each loop to avoid modifying nums itself, which caused unexpected results
on each iteration.

for i := 0; i < len(nums); i++ {
     s1 := nums[:i]
     s2 := nums[i+1:]
     new_nums := append(s1, s2...)
     if isValid(new_nums) {
             safe++
             break
     }
}

*/

func isValid(nums []string) bool {
	first_num, _ := strconv.Atoi(nums[0])
	second_num, _ := strconv.Atoi(nums[1])

	var increasing bool
	if first_num < second_num {
		increasing = true
	} else {
		increasing = false
	}

	for i := 1; i < len(nums); i++ {
		n1, _ := strconv.Atoi(nums[i-1])
		n2, _ := strconv.Atoi(nums[i])

		if increasing && n1 > n2 {
			return false
		}
		if !increasing && n1 < n2 {
			return false
		}

		diff := abs(n1 - n2)
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
