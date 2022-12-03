package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	maxWeight := 0
	inFile, err := os.Open("input.txt")
	if err != nil {
		panic("err" + err.Error())
	}

	reader := bufio.NewReader(inFile)
	for {

		currentWeight, err := getWeight(reader)
		if err != nil {
			fmt.Print(maxWeight)
			return
		}
		if currentWeight > maxWeight {
			maxWeight = currentWeight
		}

	}

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
