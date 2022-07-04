/*
Implement WordCount. It should return a map of the counts of each
“word” in the string s. The wc.Test function runs a test suite
against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)

	word_list := strings.Fields(s)

	for _, value := range word_list {
		_, ok := res[value]

		if ok == true {
			res[value] = res[value] + 1
		} else {
			res[value] = 1
		}
	}

	return res
}

func main() {
	wc.Test(WordCount)
}
