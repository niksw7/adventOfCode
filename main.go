package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	inFile, err := os.Open("input.txt")
	if err != nil {
		panic("err" + err.Error())
	}

	reader := bufio.NewReader(inFile)

	line, _, _ := reader.ReadLine()

	//slidingwindow
	oldIndex := 1
	m := map[byte]int{}
	for i := 0; i < len(line); i++ {
		if m[line[i]] != 0 && m[line[i]] >= oldIndex { //this is present in map already,reset the oldIndex
			oldIndex = m[line[i]] + 1
		}
		m[line[i]] = i + 1

		if i+1-oldIndex == 13 {
			fmt.Println("", i+1)
			return
		}
	}
}
