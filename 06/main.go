package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	for i := range input {
		if i < len(input)-4 {
			marker := input[i : i+14]
			hasDups := hasDuplicates(marker)
			if !hasDups {
				fmt.Printf("Marker %s at index %d", marker, i+14)
				break
			}
		}
	}
}

func hasDuplicates(marker []byte) bool {
	for i := range marker {
		duplicate := -1
		otherCharacters := string(marker[:i]) + string(marker[i+1:])
		duplicate = strings.Index(otherCharacters, string(marker[i]))
		if duplicate != -1 {
			return true
		}
	}
	return false
}
