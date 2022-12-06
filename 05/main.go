package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

/**
* [Q] [J]                         [H]
* [G] [S] [Q]     [Z]             [P]
* [P] [F] [M]     [F]     [F]     [S]
* [R] [R] [P] [F] [V]     [D]     [L]
* [L] [W] [W] [D] [W] [S] [V]     [G]
* [C] [H] [H] [T] [D] [L] [M] [B] [B]
* [T] [Q] [B] [S] [L] [C] [B] [J] [N]
* [F] [N] [F] [V] [Q] [Z] [Z] [T] [Q]
*  1   2   3   4   5   6   7   8   9
**/

var stacks [][]string = [][]string{
	1: {"F", "T", "C", "L", "R", "P", "G", "Q"},
	2: {"N", "Q", "H", "W", "R", "F", "S", "J"},
	3: {"F", "B", "H", "W", "P", "M", "Q"},
	4: {"V", "S", "T", "D", "F"},
	5: {"Q", "L", "D", "W", "V", "F", "Z"},
	6: {"Z", "C", "L", "S"},
	7: {"Z", "B", "M", "V", "D", "F"},
	8: {"T", "J", "B"},
	9: {"Q", "N", "B", "G", "L", "S", "P", "H"},
}

var exp *regexp.Regexp = regexp.MustCompile(`move (?P<amount>\d+) from (?P<from>\d+) to (?P<to>\d+)`)

func main() {

	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal("Error reading input:", err)
	}
	moves := strings.Split(string(input), "\n")
	for _, move := range moves {
		executeMove(move)
	}

	printStacks()
}

func executeMove(cmd string) {
	objectives := exp.FindStringSubmatch(cmd)
	amount, _ := strconv.ParseInt(objectives[1], 0, 0)
	from, _ := strconv.ParseInt(objectives[2], 0, 0)
	to, _ := strconv.ParseInt(objectives[3], 0, 0)

	// for i := 0; i < int(amount); i++ {
	// 	moveCrate(int(from), int(to))
	// }
	moveCrates(int(amount), int(from), int(to))
}

func moveCrates(amount int, oldStack int, newStack int) {
	stackLength := len(stacks[oldStack])
	crates := stacks[oldStack][stackLength-amount:]
	stacks[oldStack] = stacks[oldStack][:stackLength-amount]
	stacks[newStack] = append(stacks[newStack], crates...)
}

func moveCrate(oldStack int, newStack int) {
	crate := getTopCrate(oldStack)
	stacks[newStack] = append(stacks[newStack], crate)
}

func getTopCrate(stack int) string {
	stackSize := len(stacks[stack])
	crate := stacks[stack][stackSize-1:]
	stacks[stack] = stacks[stack][:stackSize-1]
	return crate[0]
}

func printStacks() {
	biggestStack := 0
	for _, v := range stacks {
		if len(v) > biggestStack {
			biggestStack = len(v)
		}
	}

	for i := biggestStack - 1; i >= 0; i-- {
		for s := 1; s <= 9; s++ {
			if len(stacks[s]) > i {
				if s == 9 {
					fmt.Printf("[%s]", stacks[s][i])
				} else {
					fmt.Printf("[%s] ", stacks[s][i])
				}
			} else {
				fmt.Printf("    ")
			}
		}
		fmt.Printf("\n")
	}
}
