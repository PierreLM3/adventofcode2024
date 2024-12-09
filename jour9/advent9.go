package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

var fileName = "advent9.txt"

func main() {
	input0 := read()

	input1 := make([]int, 0)
	id := 0

	dotCount := 0

	for i := 0; i < len(input0); i++ {
		if i%2 == 0 {
			for j := 0; j < input0[i]; j++ {
				input1 = append(input1, id)
			}
			id++
		} else {
			for range input0[i] {
				input1 = append(input1, -1)
				dotCount++
			}
		}
	}

	//print(input1)
	println(checkSum(move(input1, dotCount)))

}

func checkSum(input []int) int {
	sumAll := 0
	for i := 0; i < len(input); i++ {
		if input[i] == -1 {
			break
		}
		sumAll += i * input[i]
	}
	return sumAll
}

func move(input []int, dotCount int) []int {
	count := 0

	for i := len(input) - 1; i > 0; i-- {
		//	for i := len(input) - 1; i > len(input) - 2; i-- {
		if input[i] != -1 {
			writeFirstEmpty(&input, input[i])
			count++
		}
		if count == dotCount-1 {
			break
		}
		//print(input)
	}
	return input
}

func print(input []int) {
	for i := 0; i < len(input); i++ {
		if input[i] == -1 {
			fmt.Printf(".")
		} else {
			fmt.Printf("%d", input[i])
		}
	}
	println("")
}

func writeFirstEmpty(input *[]int, r int) {
	firstDot := slices.Index(*input, -1)
	(*input)[firstDot] = r

	replaceLast(input, r)
}

func replaceLast(input *[]int, y int) {
	for i := len(*input) - 1; i > 0; i-- {
		if (*input)[i] == y {
			(*input)[i] = -1
			break
		}
	}
}

func read() []int {
	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	println(fileLines[0])

	inputs := make([]int, 0)

	for _, i := range fileLines[0] {
		j, _ := strconv.Atoi(string(i))
		inputs = append(inputs, j)
	}

	return inputs
}
