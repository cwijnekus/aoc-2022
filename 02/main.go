package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var scoringMap map[string]map[string]int = map[string]map[string]int{
	"A": {"X": 3, "Y": 6, "Z": 0},
	"B": {"X": 0, "Y": 3, "Z": 6},
	"C": {"X": 6, "Y": 0, "Z": 3},
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Panic("Error reading input!", err)
	}
	rounds := strings.Split(string(input), "\n")

	score := 0
	for _, v := range rounds {
		plays := strings.Split(v, " ")
		score += determineScore(plays[0], plays[1])
	}

	fmt.Printf("Total end score is: %d", score)
}

func determineScore(play string, opposed string) int {
	return scoringMap[play][opposed] + playScore(opposed)
}

func playScore(play string) int {
	switch play {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	default:
		log.Panic("Could not determine score from play:", play)
		return 0
	}
}
