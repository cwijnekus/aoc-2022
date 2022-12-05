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

var winningMap map[string]string

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Panic("Error reading input!", err)
	}
	rounds := strings.Split(string(input), "\n")

	score := 0
	secretScore := 0
	for _, v := range rounds {
		plays := strings.Split(v, " ")
		score += determineScore(plays[0], plays[1])
		secretScore += determineOutcome(plays[0], plays[1])
	}

	fmt.Printf("Total end score is: %d", score)
	fmt.Printf("Total end secretscore is: %d", secretScore)

}

func determineOutcome(play string, outcome string) int {
	if strings.EqualFold(outcome, "Y") {
		return playScore(play) + 3
	} else if strings.EqualFold(outcome, "X") {
		if strings.EqualFold(play, "A") {
			return playScore("C")
		} else if strings.EqualFold(play, "B") {
			return playScore("A")
		} else {
			return playScore("B")
		}
	} else {
		if strings.EqualFold(play, "A") {
			return playScore("B") + 6
		} else if strings.EqualFold(play, "B") {
			return playScore("C") + 6
		} else {
			return playScore("A") + 6
		}
	}
}

func determineScore(play string, opposed string) int {
	return scoringMap[play][opposed] + playScore(opposed)
}

func playScore(play string) int {
	switch play {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	default:
		log.Panic("Could not determine score from play:", play)
		return 0
	}
}
