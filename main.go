package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	inFile, err := os.Open("input.txt")
	if err != nil {
		panic("err" + err.Error())
	}

	reader := bufio.NewReader(inFile)

	matrix := getMatrix(reader)
	for {

		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Print("********")
			for i := 0; i < 9; i++ {
				fmt.Print(matrix[i][0])
			}

			return
		}
		elems := strings.Split(string(line), " ")
		totalParts, _ := strconv.Atoi(elems[1])
		from, _ := strconv.Atoi(elems[3])
		to, _ := strconv.Atoi(elems[5])
		move(matrix, totalParts, from-1, to-1)

	}

}

func move(matrix [][]string, totalParts, from, to int) {
	fmt.Println("matrix state", matrix)
	fmt.Printf("moving %d parts from %d to %d\n", totalParts, from, to)

	elementsRemoved := matrix[from][:totalParts]
	fmt.Println("Elements removed ", elementsRemoved)
	matrix[from] = matrix[from][totalParts:]
	matrix[to] = getArrayAfterPrefix(matrix[to], copy(elementsRemoved))
	fmt.Println("Matrix now is ", matrix)

}

// t,r,e . p,k ==> t,r,e
func getArrayAfterPrefix(m1, toAdd []string) []string {
	return append(toAdd, m1...)

}

func copy(e []string) []string {
	e1 := make([]string, len(e))
	for i := 0; i < len(e); i++ {
		e1[i] = e[i]
	}
	return e1
}
func getMatrix(reader *bufio.Reader) [][]string {

	buf := make([]byte, 4)
	matrix := make([][]string, 9)

	index := 0
	for {
		_, err := io.ReadFull(reader, buf)
		if err != nil {
			return nil
		}
		str := strings.TrimSpace(string(buf))
		fmt.Println("s", str, len(str), len(string(buf)))
		if str == "" {

		} else if string(buf)[1:2] == "1" {
			reader.ReadLine()
			reader.ReadLine()
			return matrix
		} else {
			matrix[index%9] = append(matrix[index%9], string(buf)[1:2])
		}
		index++
	}
}
