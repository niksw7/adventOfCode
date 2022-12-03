package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	top3 := []int{0, 0, 0}
	inFile, err := os.Open("input.txt")
	if err != nil {
		panic("err" + err.Error())
	}

	reader := bufio.NewReader(inFile)
	for {

		currentWeight, err := getWeight(reader)
		if err != nil {
			fmt.Print(top3[2] + top3[1] + top3[0])
			return
		}
		addWeight(top3, currentWeight)
	}

}

func addWeight(top3 []int, num int) {
	if top3[0] < num {
		top3[0] = num
	}
	sort.Ints(top3)

}

func getWeight(reader *bufio.Reader) (int, error) {
	weight := 0
	for {
		s, err := reader.ReadString('\n')

		if err != nil {
			return 0, err
		}
		if len(s) == 1 {
			return weight, nil
		}
		w, err := strconv.Atoi(strings.TrimSuffix(s, "\n"))
		weight += w
	}
}
