// This program parses command line args and prints them
// In other words, it is a clone of Unix's `echo` program

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	measureTime(printRaw)
	measureTime(printWithJoin)
	measureTime(printWithLoop)
}

type printer func()

func measureTime(function printer) {
	start := time.Now()
	function()
	fmt.Println("time taken:", time.Since(start))
	fmt.Print("\n\n")
}

func printRaw() {
	fmt.Println("Raw: ", os.Args)
}

func printWithJoin() {
	fmt.Println("With strings.join: ", strings.Join(os.Args, "<<>>"))
}

func printWithLoop() {
	var s string
	for i, val := range os.Args {

		s += fmt.Sprint(i, " : ", val, " ")
	}

	fmt.Println("With loop: ", s)
}
