package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var dirExp *regexp.Regexp = regexp.MustCompile(`dir .+`)
var fileExp *regexp.Regexp = regexp.MustCompile(`\d+ .+`)
var cdInExp *regexp.Regexp = regexp.MustCompile(`\$ cd [^\.\.]+`)
var cdOutExp *regexp.Regexp = regexp.MustCompile(`\$ cd \.\.`)
var lsExp *regexp.Regexp = regexp.MustCompile(`\$ ls`)

type FSNode struct {
	Contents map[string]*FSNode
	Name     string
	Parent   *FSNode
	Size     int
}

func main() {
	root := FSNode{
		Name:     "/",
		Contents: make(map[string]*FSNode),
	}
	currentDir := &root
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		cmd := strings.Split(line, " ")
		if cmd[0] == "$" {
			if cmd[1] == "ls" {
				continue
			}

			arg := cmd[2]

			switch arg {
			case "/":
				currentDir = &root
			case "..":
				currentDir = currentDir.Parent
			default:
				currentDir = currentDir.Contents[arg]
			}
		} else {
			newNode := currentDir.createChild(cmd)
			currentDir.Contents[newNode.Name] = &newNode
		}
	}

	fmt.Println("Calculated size...")
	root.calculateSize()
	result := root.getSub100000Sum()
	fmt.Println(result)
	totalSize := 70000000
	minSize := 30000000
	freeSize := totalSize - root.Size
	reqSize := minSize - freeSize
	fmt.Printf("%d/%d Free | %d remaining\n", root.Size, totalSize, freeSize)
	fmt.Printf("%d needed for %d\n", reqSize, minSize)
	fmt.Println("Smallest: ", root.isBigEnough(reqSize))
}

func (node *FSNode) isBigEnough(reqSize int) int {
	minSize := 70000000
	for _, c := range node.Contents {
		if len(c.Contents) == 0 {
			continue
		} else {
			if c.Size < minSize && c.Size >= reqSize {
				minSize = c.Size
			}
			subDirSize := c.isBigEnough(reqSize)
			if subDirSize < minSize && subDirSize >= reqSize {
				minSize = subDirSize
			}
		}
	}
	return minSize
}

func (node *FSNode) calculateSize() int {
	for _, c := range node.Contents {
		if len(c.Contents) == 0 {
			fmt.Println("Encountered file: ", c.Name)
			node.Size += c.Size
		} else {
			fmt.Println("Calculating for dir: ", c.Name)
			node.Size += c.calculateSize()
		}
	}
	return node.Size
}

func (node *FSNode) getSub100000Sum() int {
	sum := 0
	for _, c := range node.Contents {
		if len(c.Contents) == 0 {
			continue
		} else {
			sum += c.getSub100000Sum()
			if c.Size <= 100000 {
				sum += c.Size
			}
		}
	}
	return sum
}

func (cd *FSNode) createChild(cmd []string) FSNode {
	if cmd[0] == "dir" {
		fmt.Println("Creating dir: ", cmd[1])
		return FSNode{
			Name:     cmd[1],
			Contents: make(map[string]*FSNode),
			Parent:   cd,
		}
	} else {
		fmt.Println("Creating file: ", cmd[1])
		return FSNode{
			Name:     cmd[1],
			Contents: nil,
			Parent:   cd,
			Size:     convertSize(cmd[0]),
		}
	}
}

func convertSize(sizeStr string) int {
	size, _ := strconv.ParseInt(sizeStr, 0, 0)
	return int(size)
}
