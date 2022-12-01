package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var elves map[int]int = make(map[int]int)

func main() {
	inventories, err := ioutil.ReadFile("input") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	splitInventories := strings.Split(string(inventories), "\n\n")

	for k, v := range splitInventories {
		calorieStrings := strings.Split(v, "\n")
		for _, c := range calorieStrings {
			calorie, _ := strconv.ParseInt(c, 10, 0)
			elves[k] = elves[k] + int(calorie)
		}
	}

	var elf int = 0
	var calories int = 0
	for k, v := range elves {
		if v > calories {
			calories = v
			elf = k
		}
	}

	fmt.Printf("Elf %d has the biggest calorie stack with: %d ", elf, calories)
}
