// swap swaps the values of two string variables
// passed to it

package main

import "fmt"

func main() {
	fmt.Println("Swapping vars...")
	a := "dennis"
	b := "bundi"

	a, b = swap(a, b)

	fmt.Printf("After swap...%v %v\n", a, b)
}

func swap(first, second string) (string, string) {
	fmt.Println("Before swap: ", first, second)
	// temp := &first // won't work
	temp := first
	first = second
	// second = *temp // won't work: the value contained in the pointer changes
	second = temp
	return first, second
}
