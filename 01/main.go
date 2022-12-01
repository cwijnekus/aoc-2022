package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	inventories, err := ioutil.ReadFile("input") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	splitInventories := strings.Split(string(inventories), "\n\n")

	var elves []int = make([]int, len(splitInventories))
	for _, v := range splitInventories {
		calorieStrings := strings.Split(v, "\n")
		total := 0
		for _, c := range calorieStrings {
			calorie, _ := strconv.ParseInt(c, 10, 0)
			total += int(calorie)
		}
		elves = append(elves, total)
	}

	for k := range elves {
		for l := range elves {
			if elves[k] < elves[l] {
				elves[k], elves[l] = elves[l], elves[k]
			}
		}
	}

	fmt.Printf("Elf that has the biggest calorie stack has: %d ", elves[len(elves)-1])

	together := 0
	for _, v := range elves[len(elves)-3:] {
		together += v
	}

	fmt.Printf("The three elves carrying the most are carrying %d together", together)

}
