package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("advent4.txt")

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	readFile, err := os.Open("advent4.txt")

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

	var all [][]rune

	all = make([][]rune, 140)

	for index, line := range fileLines {
		all[index] = make([]rune, len(line))
		all[index] = []rune(line)
	}

	var sumAll = 0

	for i := 1; i < 139; i++ {
		for j := 1; j < 139; j++ {
			if all[j][i] == 'A' {
				sumAll += checkX1(all, i, j)
				sumAll += checkX2(all, i, j)
				sumAll += checkX3(all, i, j)
				sumAll += checkX4(all, i, j)
			}
		}
	}

	fmt.Println(sumAll)
}

func checkX1(all [][]rune, x int, y int) int {
		if all[y+1][x+1] == 'M' && all[y-1][x+1] == 'M' && all[y-1][x-1] == 'S' && all[y+1][x-1] == 'S' {
			return 1
		} else {
			return 0
		}
}

func checkX2(all [][]rune, x int, y int) int {
		if all[y+1][x+1] == 'S' && all[y-1][x+1] == 'M' && all[y-1][x-1] == 'M' && all[y+1][x-1] == 'S' {
			return 1
		} else {
			return 0
		}
}

func checkX3(all [][]rune, x int, y int) int {
		if all[y+1][x+1] == 'S' && all[y-1][x+1] == 'S' && all[y-1][x-1] == 'M' && all[y+1][x-1] == 'M' {
			return 1
		} else {
			return 0
		}
}

func checkX4(all [][]rune, x int, y int) int {
		if all[y+1][x+1] == 'M' && all[y-1][x+1] == 'S' && all[y-1][x-1] == 'S' && all[y+1][x-1] == 'M' {
			return 1
		} else {
			return 0
		}
}

func checkE(all [][]rune, x int, y int) int {
	if x+3 < 140 {
		if all[y][x] == 'X' && all[y][x+1] == 'M' && all[y][x+2] == 'A' && all[y][x+3] == 'S' {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func checkN(all [][]rune, x int, y int) int {
	if y-3 >= 0 {
		if all[y-3][x] == 'S' && all[y-2][x] == 'A' && all[y-1][x] == 'M' && all[y][x] == 'X' {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func checkS(all [][]rune, x int, y int) int {
	if y+3 < 140 {
		if all[y+3][x] == 'S' && all[y+2][x] == 'A' && all[y+1][x] == 'M' && all[y][x] == 'X' {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func checkW(all [][]rune, x int, y int) int {
	if x-3 >= 0 {
		if all[y][x] == 'X' && all[y][x-1] == 'M' && all[y][x-2] == 'A' && all[y][x-3] == 'S' {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func checkNE(all [][]rune, x int, y int) int {
	if x+3 < 140 && y-3 >= 0 {
		if all[y-3][x+3] == 'S' && all[y-2][x+2] == 'A' && all[y-1][x+1] == 'M' && all[y][x] == 'X' {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func checkNW(all [][]rune, x int, y int) int {
	if x-3 >= 0 && y-3 >= 0 {
		if all[y-3][x-3] == 'S' && all[y-2][x-2] == 'A' && all[y-1][x-1] == 'M' && all[y][x] == 'X' {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func checkSE(all [][]rune, x int, y int) int {
	if x+3 < 140 && y+3 < 140 {
		if all[y+3][x+3] == 'S' && all[y+2][x+2] == 'A' && all[y+1][x+1] == 'M' && all[y][x] == 'X' {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func checkSW(all [][]rune, x int, y int) int {
	if x-3 >= 0 && y+3 < 140 {
		if all[y+3][x-3] == 'S' && all[y+2][x-2] == 'A' && all[y+1][x-1] == 'M' && all[y][x] == 'X' {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}
