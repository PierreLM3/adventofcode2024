package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, err := os.Open("advent3grep2.txt")

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	readFile, err := os.Open("advent3grep2.txt")

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

	rGroup, _ := regexp.Compile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)")

	sumAll := 0

	for _, line := range fileLines {
		sumAll += newFunction(rGroup, line)
	}

	fmt.Println(sumAll)
}

func newFunction(rGroup *regexp.Regexp, line string) int {
	fmt.Printf("\nf: %v\n", line)
	res := rGroup.FindAllStringSubmatch(line, -1)
	sumAll := 0
	for j := range res {
		i0, _ := strconv.Atoi(res[j][1])
		i1, _ := strconv.Atoi(res[j][2])
		sumAll += i0 * i1
	}
	fmt.Printf("f=> %v\n", sumAll)
	return sumAll
}
