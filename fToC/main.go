// fToC computes the Fahrenheit conversion of a temp in Celsius
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Temp in F: ", fToC())
}

func fToC() interface{} {
	input := os.Args
	if len(input) == 1 {
		fmt.Println("Please enter a valid number:")
		return "Error: No input"
	}

	temp, _ := strconv.Atoi(input[1])

	fmt.Println("Temp in C: ", temp)
	return temp*9/5 + 32
}
