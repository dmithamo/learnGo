/**

Implement WordCount. It should return a map of the counts of each “word” in the string s.
**/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(countWords("This is not a drill. It's a not drill"))
}

func countWords(sentence string) map[string]int {

	wordCounter := map[string]int{}

	words := strings.Fields(sentence)
	for _, word := range words {
		cleanedWord := cleanUp(word)
		_, isPresent := wordCounter[cleanedWord]

		if isPresent {
			wordCounter[cleanedWord]++
		} else {
			wordCounter[cleanedWord] = 1
		}
	}

	return wordCounter

}

/*
* @description Remove non-alphanumeric x-ters fromeach word
 */
func cleanUp(str string) string {
	nonAlphanumerics := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	unwanted := nonAlphanumerics.FindAllString(str, -1)
	unwantedAsString := strings.Join(unwanted, " ")

	return strings.Trim(str, unwantedAsString)

}
