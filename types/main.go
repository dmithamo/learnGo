package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args
	if len(input) == 1 {
		fmt.Println("Try again and provide a value")
		return
	}

	userString := input[1]
	isAlphaNumeric := true

	for i := 0; i < len(userString); i++ {
		char := userString[i]
		if 'a' <= char && char <= 'z' ||
			'A' <= char && char <= 'Z' ||
			'0' <= char && char <= '9' {
			continue
		} else {
			isAlphaNumeric = false
		}
	}

	fmt.Println(userString, ":", "isAlphaNumeric? : ", isAlphaNumeric)
}
