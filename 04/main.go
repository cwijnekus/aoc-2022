package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type assignment struct {
	min int
	max int
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	assignments := strings.Split(string(input), "\n")

	containginPairs := 0
	overlappingPairs := 0
	for _, v := range assignments {
		assign1, assign2 := splitAssignments(v)
		if checkContains(assign1, assign2) || checkContains(assign2, assign1) {
			containginPairs++
		}
		if checkOverlap(assign1, assign2) || checkOverlap(assign2, assign1) {
			overlappingPairs++
		}
	}

	fmt.Printf("Amount of containing groups are %d", containginPairs)
	fmt.Printf("Amount of overlapping groups are %d", overlappingPairs)
}

func checkOverlap(range1 assignment, range2 assignment) bool {
	if range2.min >= range1.min && range2.min <= range1.max {
		return true
	} else if range2.max >= range1.min && range2.max <= range1.max {
		return true
	} else {
		fmt.Printf("Range 1: %v | Range 2: %v\n", range1, range2)
		return false
	}
}

func checkContains(range1 assignment, range2 assignment) bool {
	return range1.min >= range2.min && range1.max <= range2.max
}

func splitAssignments(assignment string) (assignment, assignment) {
	ranges := strings.Split(assignment, ",")
	range1 := splitRange(ranges[0])
	range2 := splitRange(ranges[1])
	return range1, range2
}

func splitRange(assignedRange string) assignment {
	limits := strings.Split(assignedRange, "-")
	minLimit, err := strconv.ParseInt(limits[0], 0, 0)
	if err != nil {
		log.Fatal("Error strconv: ", err)
	}
	maxLimit, err := strconv.ParseInt(limits[1], 0, 0)
	if err != nil {
		log.Fatal("Error strconv: ", err)
	}
	return assignment{
		min: int(minLimit),
		max: int(maxLimit),
	}
}
