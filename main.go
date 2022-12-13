package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	matrix := buildMatrix()
	count := 0
	d := generateVisibility(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if d[i][j].min() < matrix[i][j] {
				count++
			}
		}
	}
	fmt.Println(count)
}

type Data struct {
	left   int
	right  int
	top    int
	bottom int
}

func (d Data) min() int {
	a := []int{d.left, d.right, d.top, d.bottom}
	sort.Ints(a)
	return a[0]
}

func generateVisibility(matrix [][]int) [][]Data {
	d := [][]Data{}
	//rows
	for i := 0; i < len(matrix); i++ {
		d = append(d, []Data{})
		lastBig := -1
		for j := 0; j < len(matrix[0]); j++ {
			d[i] = append(d[i], Data{left: lastBig})
			if matrix[i][j] > lastBig {
				lastBig = matrix[i][j]
			}
		}
	}
	for i := len(matrix) - 1; i >= 0; i-- {
		lastBig := -1
		for j := len(matrix[0]) - 1; j >= 0; j-- {

			d[i][j].right = lastBig
			if matrix[i][j] > lastBig {
				lastBig = matrix[i][j]
			}
		}
	}

	//cols
	for i := 0; i < len(matrix); i++ {
		d = append(d, []Data{})
		lastBig := -1
		for j := 0; j < len(matrix[0]); j++ {

			d[j][i].top = lastBig
			if matrix[j][i] > lastBig {
				lastBig = matrix[j][i]
			}
		}
	}
	for i := len(matrix) - 1; i >= 0; i-- {
		lastBig := -1
		for j := len(matrix[0]) - 1; j >= 0; j-- {

			d[j][i].bottom = lastBig
			if matrix[j][i] > lastBig {
				lastBig = matrix[j][i]
			}
		}
	}
	return d

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
