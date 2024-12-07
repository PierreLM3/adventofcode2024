package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("advent1.csv")
	var l1 []int
	var l2 []int

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '

	records, err := reader.ReadAll()

	for _, eachrecord := range records {
		i, _ := strconv.Atoi(eachrecord[0])
		j, _ := strconv.Atoi(eachrecord[3])

		l1 = append(l1, i)
		l2 = append(l2, j)
	}

	sort.Sort(sort.IntSlice(l1))
	sort.Sort(sort.IntSlice(l2))

	sum := 0

	for i := 0; i < len(l1); i++ {
		sum += l1[i] * countInSlice(l2, l1[i])
	}

	println(sum)
}

func countInSlice(slice []int, n int) int {
    count := 0
    for _, s := range slice {
        if s == n {
            count++
        }
    }
    return count
}
