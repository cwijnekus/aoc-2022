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

var root map[string]FSEntity = make(map[string]FSEntity)

type FSEntity interface {
	GetSize() int
}

type File struct {
	Size int
}

func (f File) GetSize() int {
	return f.Size
}

type Directory struct {
	Size    int
	Content map[string]FSEntity
}

func (d Directory) GetSize() int {
	return d.Size
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	dirs, _ := processDirectory(lines)
	if v, ok :=
}

func processDirectory(lines []string) (FSEntity, int) {
	dir := Directory{}
	for i := 0; i < len(lines); i++ {
		if strings.EqualFold("$ cd ..", lines[i]) {
			return dir, i
		} else if strings.HasPrefix("$ cd ", lines[i]) {
			nested, index := processDirectory(lines[i:])
			dir.Content[getDirNameFromCmd(lines[i])] = nested
			i = index
		} else if strings.HasPrefix("dir ", lines[i]) {
			dir.Content[getDirNameFromLs(lines[i])] = Directory{}
		} else {
			file := strings.Split(lines[i], " ")
			dir.Content[file[1]] = File{Size: parseFileSize(file[0])}
		}
	}
	return dir, 0
}

func getDirNameFromCmd(cmd string) string {
	return strings.Split(cmd, " ")[2]
}

func getDirNameFromLs(line string) string {
	return strings.Split(line, " ")[1]
}

func parseFileSize(strSize string) int {
	size, err := strconv.ParseInt(strSize, 0, 0)
	if err != nil {
		panic(err)
	}
	return int(size)
}
