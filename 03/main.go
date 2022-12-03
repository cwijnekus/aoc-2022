package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Panic("Error reading input", err)
	}

	inventories := strings.Split(string(input), "\n")

	priorities := 0
	for _, v := range inventories {
		priorities += intersectingItems(splitInventories(v))
	}

	fmt.Printf("The total amount of priorities is: %d", priorities)
}

func splitInventories(inv string) (string, string) {
	return inv[:(len(inv) / 2)], inv[(len(inv) / 2):]
}

func intersectingItems(part1 string, part2 string) int {
	for _, v := range part1 {
		if strings.Index(part2, string(v)) >= 0 {
			fmt.Printf("Intersecting char: %s has prio %d\n", string(v), getPrio(string(v)))
			return getPrio(string(v))
		}
	}
	return 0

}

func getPrio(character string) int {
	return strings.Index("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", character) + 1
}
