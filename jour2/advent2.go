package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("advent2.csv")

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	readFile, err := os.Open("advent2.csv")

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

	var safe = 0

	for _, line := range fileLines {
		columns := strings.Split(line, " ")

		errorFull := newFunction(columns)

		if !errorFull {
			safe += 1
			log.Printf("all: %v error:%v safe:%v\n", columns, errorFull, safe)
			continue
		}

		errorCol := true
		for i := 0; i < len(columns); i++ {

			columnsMinus := RemoveIndex(columns, i)

			errorCol = newFunction(columnsMinus)
			if !errorCol {
				break
			}
		}

		if !errorCol {
			safe += 1
		}

		log.Printf("%v error:%v safe:%v\n", columns, errorCol, safe)
	}
}

func RemoveIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func newFunction(columns []string) bool {
	i0, _ := strconv.Atoi(columns[0])
	i1, _ := strconv.Atoi(columns[1])

	var error = false

	if i0 == i1 {
		error = true
	} else {
		var asc = i0 < i1

		for i := 0; i < len(columns)-1; i++ {
			left, _ := strconv.Atoi(columns[i])
			right, _ := strconv.Atoi(columns[i+1])

			if asc {
				if left >= right || (right-left) > 3 {
					error = true
					break
				}
			} else {
				if left <= right || (left-right) > 3 {
					error = true
					break
				}
			}
		}
	}

	return error
}
