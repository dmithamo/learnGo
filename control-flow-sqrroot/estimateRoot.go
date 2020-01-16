/**

Computers typically compute the square root of x using a loop. Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:

z -= (z*z - x) / (2*z)

Repeating this adjustment makes the guess better and better until we reach an answer that is as close to the actual square root as can be.

Implement this in the func Sqrt provided. A decent starting guess for z is 1, no matter what the input. To begin with, repeat the calculation 10 times and print each z along the way. See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.

**/

package main

import "fmt"

func main() {
	estimateRoot(225)
}

func estimateRoot(number int) {
	numAsFloat := float64(number)
	accuracy := 0.0000001
	guess := 1.0
	numberOfGuesses := 0
	// This algorithm is a blackbox to me :}
	fmt.Printf("------\nCalculating the squareroot of %v by pure guesswork\n", number)

	for guess*guess-numAsFloat > accuracy || numAsFloat-guess*guess > accuracy {

		guess -= (guess*guess - numAsFloat) / (2 * guess)
		numberOfGuesses++

		printTheProcess(numberOfGuesses, guess)
	}
	fmt.Printf("------\nThe squareroot of %v approaches %v\n", number, guess)
	fmt.Printf("We figured that out in %v guesses\n-------\n", numberOfGuesses)
}

func printTheProcess(timesGuessed int, currentGuess float64) {
	fmt.Printf("\tGuess number: %v <<>> Value %v\n", timesGuessed, currentGuess)
}
