package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
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

var crates [][]string = [][]string{
	1: {"Q", "J", "", "", "", "", "", "", "H"},
	2: {"G", "S", "Q", "", "Z", "", "", "", "P"},
	3: {"P", "F", "M", "", "F", "", "F", "", "S"},
	4: {"R", "R", "P", "F", "V", "", "D", "", "L"},
	5: {"L", "W", "W", "D", "W", "S", "V", "", "G"},
	6: {"C", "H", "H", "T", "D", "L", "M", "B", "B"},
	7: {"T", "Q", "B", "S", "L", "C", "B", "J", "N"},
	8: {"F", "N", "F", "V", "Q", "Z", "Z", "T", "Q"},
}

func main() {

	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal("Error reading input:", err)
	}
	moves := strings.Split(string(input), "\n")
	objectives := regexp.MustCompile(`(move (?P<amount>\d+) )(from (?P<from>\d+) )(to (?P<to>\d+))`).FindStringSubmatch(moves[0])

	fmt.Println(objectives)
}
