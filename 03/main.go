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

	fmt.Printf("The total amount of priorities is: %d\n", priorities)

	groups := make([][]string, 0, len(inventories)/3)
	for i := 0; i < len(inventories); i++ {
		group := inventories[:3]
		inventories = inventories[3:]
		groups = append(groups, group)
		i = 0
	}

	groupPrio := 0

	for _, group := range groups {
		for _, v := range group[0] {
			if strings.Index(group[1], string(v)) >= 0 && strings.Index(group[2], string(v)) >= 0 {
				fmt.Printf("We got a match: %s\n", string(v))
				groupPrio += getPrio(string(v))
				break
			}
		}
	}

	fmt.Printf("The total amount of group prio's: %d", groupPrio)
}

func splitInventories(inv string) (string, string) {
	return inv[:(len(inv) / 2)], inv[(len(inv) / 2):]
}

func intersectingItems(part1 string, part2 string) int {
	for _, v := range part1 {
		if strings.Index(part2, string(v)) >= 0 {
			return getPrio(string(v))
		}
	}
	return 0
}

func getPrio(character string) int {
	return strings.Index("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", character) + 1
}
