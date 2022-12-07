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
		// fmt.Printf("Starting comparison!\n\n")
		// i := 0
		if i < len(input)-4 {
			marker := input[i : i+4]
			hasDups := hasDuplicates(marker)
			if !hasDups {
				fmt.Printf("Marker %s at index %d", marker, i+4)
				break
			}
		}
		// fmt.Printf("-------------------------------------------\n")
	}
}

func hasDuplicates(marker []byte) bool {
	for i := range marker {
		duplicate := -1
		otherCharacters := string(marker[:i]) + string(marker[i+1:])
		// fmt.Println("Comparing ", otherCharacters)
		// fmt.Println("To: ", string(marker[i]))
		duplicate = strings.Index(otherCharacters, string(marker[i]))
		if duplicate != -1 {
			return true
		}
	}
	return false
}
