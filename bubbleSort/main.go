// bubbleSort sorts an array by bubbling the largest value
// to the end of the array

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(bubbleSort())
}

func bubbleSort() interface{} {
	input := os.Args

	if len(input) == 1 {
		return "No values provided"
	}

	arrOfNums := make([]int, len(input[1:]))

	// convert each cli arg to an int first
	for i, val := range input[1:] {
		converted, _ := strconv.Atoi(val)
		arrOfNums[i] = converted
	}

	fmt.Println("Unsorted values: ", arrOfNums)

	for i := 0; i < len(arrOfNums); i++ {
		for j := len(arrOfNums) - 1; j >= 0; j-- {

			// If i >= j comparison is unnecessary
			if i >= j {
				fmt.Println("Skipping [i/j]: ", i, j)
				continue
			}

			fmt.Println(arrOfNums, "i/j", i, j)
			if arrOfNums[i] > arrOfNums[j] {
				arrOfNums[i], arrOfNums[j] = arrOfNums[j], arrOfNums[i]
			}
		}
	}

	return fmt.Sprint("Sorted values: ", arrOfNums)
}
