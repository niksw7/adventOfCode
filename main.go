package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var smallestElement int = 100000000

func main() {

	tree := BuildTree()
	totalDiskSpace := 70000000

	rootDir := tree.files["/"]
	totalSize := rootDir.getSize()
	available := totalDiskSpace - totalSize
	needed := 30000000 - available
	recurse(rootDir, needed)
	fmt.Println(smallestElement)

}

func recurse(t *Tree, sizeToClean int) {
	if t == nil || t.size != 0 {
		return
	}
	dirSize := t.getSize()
	if dirSize >= sizeToClean && smallestElement > dirSize {
		smallestElement = dirSize
	}

	for _, v := range t.files {
		if v.size == 0 {
			recurse(v, sizeToClean)
		}
	}

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
