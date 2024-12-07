package main

import (
	"bufio"
	"fmt"
	"os"
)

var squareSize = 130
var fileName = "advent6.txt"

func main() {
	var sumAll = 0

	var initialMap = read()
	var _, vis0 = stepAll(initialMap)
	var vis1 = removeDuplicates(vis0)

	fmt.Printf("%v %v\n", len(vis0), len(vis1))

	for index, visit := range vis1 {
		var i = visit[0]
		var j = visit[1]

		var startMap, err = newMap(i, j)
		if err == nil {
			var add, _ = stepAll(startMap)
			sumAll += add
			if(add == 1) {
				fmt.Printf("%v/%v sum:%v\n", index+1, len(vis1), sumAll)
			}
		}
	}

	println(sumAll)
}

func removeDuplicates(slice [][3]int) [][2]int {
	seen := make(map[string]bool)
	result := [][2]int{}

	for _, val := range slice {
		var key = fmt.Sprintf("%d%s%d", val[0], ",", val[1])
		if _, ok := seen[key]; !ok {
			seen[key] = true
			result = append(result, [2]int{val[0], val[1]})
		}
	}
	return result
}

func newMap(i int, j int) ([][]rune, error) {
	var startMap = read()

	if startMap[i][j] != '^' && startMap[i][j] != '#' {
		startMap[i][j] = '#'
		return startMap, nil
	} else {
		return startMap, fmt.Errorf("")
	}
}

func stepAll(oldMap [][]rune) (int, [][3]int) {
	var oldVisited [][3]int
	var oldGuard = findGuard(oldMap)

	for true {
		var posY = oldGuard[0]
		var posX = oldGuard[1]
		var direction = oldGuard[2]

		newMap, newGuard, err := step(oldMap, direction, posY, posX)
		if err != nil {
			// END
			// println("END")
			newVisitedEnd, _ := visited(oldVisited, newGuard)
			return 0, newVisitedEnd
		}

		oldMap = newMap
		newVisited, errVis := visited(oldVisited, newGuard)
		if errVis != nil {
			// LOOP
			// println("LOOP")
			return 1, newVisited
		}

		oldVisited = newVisited
		oldGuard = newGuard
	}

	return 0, oldVisited
}

func visited(visited [][3]int, guard [3]int) ([][3]int, error) {
	for _, vis := range visited {
		if vis[0] == guard[0] && vis[1] == guard[1] && vis[2] == guard[2] {
			return visited, fmt.Errorf("Already visited")
		}
	}

	visited = append(visited, guard)
	return visited, nil
}

func step(mmap [][]rune, direction int, posY int, posX int) ([][]rune, [3]int, error) {
	var guard [3]int

	if direction == 0 && posY-1 >= 0 && mmap[posY-1][posX] != '#' { // N
		mmap[posY][posX] = 'X'
		mmap[posY-1][posX] = '^'
		guard = [3]int{posY - 1, posX, 0}
	} else if direction == 0 && posY-1 >= 0 && mmap[posY-1][posX] == '#' {
		mmap[posY][posX] = 'X'
		mmap[posY][posX+1] = 'E'
		guard = [3]int{posY, posX + 1, 1}
	} else if direction == 1 && posX+1 < squareSize && mmap[posY][posX+1] != '#' { // E
		mmap[posY][posX] = 'X'
		mmap[posY][posX+1] = 'E'
		guard = [3]int{posY, posX + 1, 1}
	} else if direction == 1 && posX+1 < squareSize && mmap[posY][posX+1] == '#' {
		mmap[posY][posX] = 'X'
		mmap[posY+1][posX] = 'S'
		guard = [3]int{posY + 1, posX, 2}
	} else if direction == 2 && posY+1 < squareSize && mmap[posY+1][posX] != '#' { // S
		mmap[posY][posX] = 'X'
		mmap[posY+1][posX] = 'S'
		guard = [3]int{posY + 1, posX, 2}
	} else if direction == 2 && posY+1 < squareSize && mmap[posY+1][posX] == '#' {
		mmap[posY][posX] = 'X'
		mmap[posY][posX-1] = 'W'
		guard = [3]int{posY, posX - 1, 3}
	} else if direction == 3 && posX-1 >= 0 && mmap[posY][posX-1] != '#' { // W
		mmap[posY][posX] = 'X'
		mmap[posY][posX-1] = 'W'
		guard = [3]int{posY, posX - 1, 3}
	} else if direction == 3 && posX-1 >= 0 && mmap[posY][posX-1] == '#' {
		mmap[posY][posX] = 'X'
		mmap[posY-1][posX] = '^'
		guard = [3]int{posY - 1, posX, 0}
	} else {
		guard = [3]int{posY, posX, 0}
		return mmap, guard, fmt.Errorf("Out of bound")
	}

	return mmap, guard, nil
}

func findGuard(mmap [][]rune) [3]int {
	for i, row := range mmap {
		for j, val := range row {
			if val == '^' {
				return [3]int{i, j, 0}
			}
		}
	}
	return [3]int{-1, -1, 0}
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

	var mmap [][]rune

	mmap = make([][]rune, squareSize)

	for index, line := range fileLines {
		mmap[index] = make([]rune, len(line))
		mmap[index] = []rune(line)
	}

	return mmap
}
