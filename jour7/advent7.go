package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var fileName = "advent7.txt"
var lineTotal = 850 // 9

func main() {

	var key, values = read()

	sumAll := 0

	for i := 0; i < lineTotal; i++ {
		var nbValues = len(values[i])

		var allCombinaisons = int(math.Pow(3, float64(nbValues-1)))
		for j := 0; j < allCombinaisons; j++ {
			combin := leftPad(strconv.FormatInt(int64(j), 3), "0", nbValues-1)
			res := calc(values[i], combin)

			if res == key[i] {
				//fmt.Printf("solve %v=%v\n", key[i], combin)
				sumAll += key[i]
				break
			}
		}
	}
	println(sumAll)
}

func leftPad(s string, padStr string, pLen int) string {
	return strings.Repeat(padStr, pLen-len(s)) + s
}

func calc(values []int, combin string) int {
	var c = values[0]
	for i := 0; i < len(combin); i++ {
		if combin[i] == 48 {
			c += values[i+1]
		} else if combin[i] == 49 {
			c *= values[i+1]
		} else {
			c, _ = strconv.Atoi(fmt.Sprintf("%d%d", c, values[i+1]))
		}
	}
	return c
}

func read() ([]int, [][]int) {
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

	var allTotal = make([]int, 0)
	var all = make([][]int, 0)

	for _, line := range fileLines {
		var split1 = strings.Split(line, ": ")
		t, _ := strconv.Atoi(split1[0])
		allTotal = append(allTotal, t)

		var split2 = strings.Split(split1[1], " ")
		var split2int = make([]int, 0)

		for i := 0; i < len(split2); i++ {
			t2, _ := strconv.Atoi(split2[i])
			split2int = append(split2int, t2)
		}
		all = append(all, split2int)
	}

	return allTotal, all
}
