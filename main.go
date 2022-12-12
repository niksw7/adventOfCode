package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	tree := BuildTree()
	fmt.Println(recurse(tree.files["/"]))

}

func recurse(t *Tree) int {
	size := 0
	if t == nil || t.size != 0 {
		return 0
	}
	if t.getSize() <= 100000 {
		size += t.getSize()
	}

	for _, v := range t.files {
		if v.size == 0 {
			size += recurse(v)
		}
	}
	return size
}

type Tree struct {
	name      string
	files     map[string]*Tree
	size      int
	parentDir *Tree
}

func (t Tree) getSize() int {
	if t.size != 0 {
		return t.size
	}
	size := 0
	for _, v := range t.files {
		size += v.getSize()
	}
	return size
}

func BuildTree() *Tree {

	inFile, err := os.Open("input.txt")
	if err != nil {
		panic("err" + err.Error())
	}

	reader := bufio.NewReader(inFile)
	root := &Tree{name: "root"}
	current := root
	cwd := "/"

	for {

		line, _, err := reader.ReadLine()
		if err != nil {
			return root
		}
		elements := strings.Split(string(line), " ")

		if elements[0] == "$" {
			//its a command
			switch elements[1] {
			case "cd":
				cwd = elements[2]
				if cwd == ".." {
					//move up
					current = current.parentDir
				} else {
					//move down
					if _, ok := current.files[cwd]; !ok {
						if current.files == nil {
							current.files = map[string]*Tree{}
						}
						current.files[cwd] = &Tree{name: cwd, parentDir: current}
					}
					current = current.files[cwd]
				}
			case "ls":
				//do nothing
			default:
				panic("unrecognised" + elements[1])
			}
		} else {
			//non command
			if elements[0] == "dir" {
				if _, ok := current.files[elements[1]]; !ok {
					if current.files == nil {
						current.files = map[string]*Tree{}
					}
					current.files[elements[1]] = &Tree{name: elements[1], parentDir: current}
				}
			} else {
				size, _ := strconv.Atoi(elements[0])
				if current.files == nil {
					current.files = map[string]*Tree{}
				}
				current.files[elements[1]] = &Tree{name: elements[1], parentDir: current, size: size}
			}

		}
	}
}
