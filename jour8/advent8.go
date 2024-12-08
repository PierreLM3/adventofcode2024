package main

import (
	"bufio"
	"fmt"
	"os"
)

var fileName = "advent8.txt"
var lineTotal = 50 // 12

func main() {

	var antennaMaps = read()
	printMap(antennaMaps)

	antinodesMap := make([][]rune, lineTotal)
	

	for i := range lineTotal {
		antinodesMap[i] = make([]rune, lineTotal)
		for j := range lineTotal {
			antinodesMap[i][j] = '.'
		}
	}

	for _, n := range searchNext(antennaMaps, '0', 1, 1) {
		fmt.Printf("%v:%v\n", n[0], n[1])
	}

	for i := range lineTotal {
		for j := range lineTotal {
			if antennaMaps[i][j] != '.' {
				nexts := searchNext(antennaMaps, antennaMaps[i][j], i, j)
				putAntinodes(antinodesMap, i, j, nexts)
			}
		}
	}

	printMap(antinodesMap)
	println(countAntinodes(antinodesMap))
}

func countAntinodes(antinodesMap [][]rune) int {
	sumAll := 0
	for i := range lineTotal {
		for j := range lineTotal {
			if antinodesMap[i][j] == '#' {
				sumAll += 1
			}
		}
	}
	return sumAll
}

func putAntinodes(antinodesMap [][]rune, inputI int, inputJ int, nexts [][2]int) {
	antinodesMap[inputI][inputJ] = '#'
	
	for _, next := range nexts {
		dist1i := next[0] - inputI
		antinode1i := inputI - dist1i
		dist1j := next[1] - inputJ
		antinode1j := inputJ - dist1j

		for antinode1i >= 0 && antinode1j >= 0 && antinode1j < lineTotal {
			//fmt.Printf("1: letter:%v antinode1:(%v:%v)\n", string(letter), antinode1i, antinode1j)
			antinodesMap[antinode1i][antinode1j] = '#'

			antinode1i = antinode1i - dist1i
			antinode1j = antinode1j - dist1j
		}

		antinode2i := next[0] + (next[0] - inputI)
		antinode2j := next[1] + (next[1] - inputJ)

		for antinode2i < lineTotal && antinode2j >= 0 && antinode2j < lineTotal {
			//fmt.Printf("2: letter:%v antinode2:(%v,%v)\n", string(letter), antinode2i, antinode2j)
			antinodesMap[antinode2i][antinode2j] = '#'

			antinode2i = antinode2i + dist1i
			antinode2j = antinode2j + dist1j
		}
	}
}

func searchNext(mmap [][]rune, letter rune, inputI int, inputJ int) [][2]int {
	nexts := make([][2]int, 0)

	for j := inputJ + 1; j < lineTotal; j++ {
		if mmap[inputI][j] == letter {
			nexts = append(nexts, [2]int{inputI, j})
		}
	}

	for i := inputI + 1; i < lineTotal; i++ {
		for j := 0; j < lineTotal; j++ {
			if mmap[i][j] == letter {
				nexts = append(nexts, [2]int{i, j})
			}
		}
	}

	return nexts
}

func printMap(mmap [][]rune) {
	for _, runes := range mmap {
		println(string(runes))
	}
}

func read() [][]rune {
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

	var all = make([][]rune, lineTotal)

	for index, line := range fileLines {
		all[index] = make([]rune, lineTotal)
		for j, r := range line {
			all[index][j] = r
		}
	}

	return all
}
