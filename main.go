package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	matrix := buildMatrix()
	max := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			s := score(matrix, i, j)
			if s > max {
				max = s
			}
		}
	}
	fmt.Println(max)
}
func score(m [][]int, row, col int) int {
	rscore1 := 0
	//go ahead
	for c := col + 1; c < len(m[0]); c++ {
		if m[row][c] >= m[row][col] {
			rscore1++
			break
		}
		rscore1++
	}
	rscore2 := 0
	//go behind
	for c := col - 1; c >= 0; c-- {
		if m[row][c] >= m[row][col] {
			rscore2++
			break
		}
		rscore2++
	}
	cscore1 := 0
	//go down
	for r := row + 1; r < len(m); r++ {
		if m[r][col] >= m[row][col] {
			cscore1++
			break
		}
		cscore1++
	}
	//go up
	cscore2 := 0
	for r := row - 1; r >= 0; r-- {
		if m[r][col] >= m[row][col] {
			cscore2++
			break
		}
		cscore2++
	}
	//fmt.Println(rscore1, rscore2, cscore1, cscore2)
	return rscore1 * rscore2 * cscore1 * cscore2
}

func buildMatrix() [][]int {
	inFile, err := os.Open("input.txt")
	if err != nil {
		panic("err" + err.Error())
	}
	matrix := [][]int{}
	reader := bufio.NewReader(inFile)
	for {

		line, _, err := reader.ReadLine()
		if err != nil {
			return matrix
		}
		arr := []int{}
		for _, b := range line {
			asInt := int(b) - 48
			arr = append(arr, asInt)
		}
		matrix = append(matrix, arr)
	}
}
