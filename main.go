package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inFile, err := os.Open("input.txt")
	if err != nil {
		panic("err" + err.Error())
	}
	tracker := map[string]bool{}
	count := 0

	reader := bufio.NewReader(inFile)

	knotPositions := [][]int{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}
	for {

		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("count=%d\n", count)

			return
		}

		split := strings.Split(string(line), " ")
		character := split[0]
		move, _ := strconv.Atoi(split[1])

		for i := 1; i <= move; i++ {

			switch character {
			case "L":
				//move the head first
				knotPositions[0] = []int{knotPositions[0][0], knotPositions[0][1] - 1}
			case "R":
				knotPositions[0] = []int{knotPositions[0][0], knotPositions[0][1] + 1}
			case "U":
				knotPositions[0] = []int{knotPositions[0][0] + 1, knotPositions[0][1]}
			case "D":
				knotPositions[0] = []int{knotPositions[0][0] - 1, knotPositions[0][1]}
			}
			//traverse for all the tails
			for i := 1; i <= 9; i++ {
				r, c := getPosToAdjust(knotPositions[i-1], knotPositions[i])
				knotPositions[i] = []int{knotPositions[i][0] + r, knotPositions[i][1] + c}
			}
			if _, ok := tracker[fmt.Sprintf("%d,%d", knotPositions[9][0], knotPositions[9][1])]; !ok {
				tracker[fmt.Sprintf("%d,%d", knotPositions[9][0], knotPositions[9][1])] = true
				count++
			}

		}
	}

}

func getPosToAdjust(headPosition, tailPosition []int) (int, int) {
	rowDifference := headPosition[0] - tailPosition[0]
	colDifference := headPosition[1] - tailPosition[1]
	r := 0
	c := 0
	if isAdjacent(rowDifference, colDifference) {
		return 0, 0
	}

	if rowDifference > 0 {
		r = 1
	} else if rowDifference < 0 {
		r = -1
	}
	if colDifference > 0 {
		c = 1
	} else if colDifference < 0 {
		c = -1
	}
	return r, c

}

func isAdjacent(rowDifference, colDifference int) bool {

	//The commented part is a bug that made me quit AOC and took almost lot of time.(Not to fix the bug, but a transient issue on submiting the solution)
	/* if rowDifference == 0 && colDifference == 0 {
		return true
	} else if rowDifference == 0 && (colDifference == -1 || colDifference == 1) {
		return true
	} else if colDifference == 0 && (rowDifference == -1 || rowDifference == 1) {
		return true
	} else if colDifference+rowDifference == 0 || colDifference-rowDifference == 0 {
		return true
	}
	return false */
	if math.Abs(float64(rowDifference)) <= 1 && math.Abs(float64(colDifference)) <= 1 {
		return true
	}
	return false
}
