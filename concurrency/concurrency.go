package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\nFibonacci(%d) = %d\n", n, fibN)
}
func spinner(delay time.Duration) {
	percentageDone := 0.
	for {
			fmt.Printf("\r[%v  %%]", percentageDone)
			percentageDone += 1.52
			time.Sleep(delay)
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
