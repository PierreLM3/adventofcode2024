package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	rules1 := read1()
	rules2 := read2()

	sumAll := 0

	for _, r2 := range rules2 {
		var res = checkLine(r2, rules1)
		if res != nil {
			var r2sort = r2
			for res != nil {
				r2sort = Revert(r2sort, res[0], res[1])
				res = checkLine(r2sort, rules1)
			}
			sumAll += middle(r2sort)
		}
	}

	println(sumAll)
}

func middle(r2 []int) int {
	return r2[len(r2)/2]
}

func checkLine(r2 []int, rules1 [][]int) []int {
	for i := 0; i < len(r2); i++ {
		res := check(r2[i], RemoveIndex(r2, i), rules1)

		if res != nil {
			return res
		}
	}

	return nil
}

func check(n int, rights []int, rules [][]int) []int {
	for _, rule := range rules {
		if rule[1] == n {
			for _, right := range rights {
				if right == rule[0] {
					return rule
				}
			}
		}
	}
	return nil
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	return append(ret, s[index:]...)
}

func Revert(s []int, v1 int, v2 int) []int {
	ret := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == v1 {
			ret[i] = v2
		} else if s[i] == v2 {
			ret[i] = v1
		} else {
			ret[i] = s[i]
		}
	}

	return ret
}

func read1() [][]int {

	readFile, err := os.Open("a5.txt")

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

	var rules [][]int

	rules = make([][]int, 1176)
	rGroup, _ := regexp.Compile("([0-9]{2})\\|([0-9]{2})")

	for index, line := range fileLines {
		rules[index] = make([]int, 2)
		res := rGroup.FindAllStringSubmatch(line, -1)
		i0, _ := strconv.Atoi(res[0][1])
		i1, _ := strconv.Atoi(res[0][2])
		rules[index] = []int{i0, i1}
	}

	return rules
}

func read2() [][]int {
	readFile, err := os.Open("a5bis.txt")

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

	var rules [][]string
	rules = make([][]string, 196)

	for index, line := range fileLines {
		rules[index] = strings.Split(line, ",")
	}

	var rulesInt [][]int
	rulesInt = make([][]int, 196)

	for index, line := range rules {
		rulesInt[index] = make([]int, len(line))

		for i := 0; i < len(line); i++ {
			j, _ := strconv.Atoi(line[i])
			rulesInt[index][i] = j
		}
	}

	return rulesInt
}
