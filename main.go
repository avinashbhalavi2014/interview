package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	timeRegex = regexp.MustCompile("^([01]?[0-9]|2[0-3]):[0-5][0-9]")
)

// permutation function will return all the possible integer combinations using recursion
func permutation(integers []int) [][]int {

	var helper func([]int, int, int)
	res := [][]int{}

	helper = func(array []int, index int, lastIndex int) {

		if index == lastIndex {
			tmp := make([]int, len(array))
			copy(tmp, array)
			res = append(res, tmp)
		} else {

			for i := index; i <= lastIndex; i++ {

				// swap the integers
				array[index], array[i] = array[i], array[index]

				// recursion
				helper(array, index+1, lastIndex)

				// backtrack -> to original array
				array[index], array[i] = array[i], array[index]
			}
		}

	}

	helper(integers, 0, len(integers)-1)
	return res

}

// validate24HoursTime function will validate the integer combination and returns the counts
func validate24HoursTime(combinationArray [][]int) (int, error) {

	// declare variable count and initialize to 0
	count := 0

	// iterate over combinationArray
	for _, seq := range combinationArray {

		if len(seq) == 4 {
			// convert integers to string time format
			str := strconv.Itoa(seq[0]) + strconv.Itoa(seq[1]) + ":" + strconv.Itoa(seq[2]) + strconv.Itoa(seq[3])
			// match the string with regex
			valid := timeRegex.MatchString(str)

			if valid {
				count++ // increment the counter if string is valid time
			}
		}
	}
	return count, nil
}

func validateInput(input []int) error {
	visited := make(map[int]bool, 0)
	for _, val := range input {

		if visited[val] == true {
			return fmt.Errorf("duplicate value found: %v", val)
		} else {
			visited[val] = true
		}

		if val < 0 {
			return fmt.Errorf("input values cannot be negative")
		}
	}

	return nil
}

func main() {

	// declare integers values
	integers := []int{1, 2, 3, 3}

	// validate input integers
	err := validateInput(integers)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	// get all the possible integers arrangements
	result := permutation(integers)

	// validate the combination
	count, err := validate24HoursTime(result)
	if err != nil {
		fmt.Println("errors: %w", err)
		return
	}

	// Print the result in console
	fmt.Println("result:", count)

}
