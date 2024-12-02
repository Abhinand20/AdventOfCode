package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseInput(text string) [][]int {
	tokens := strings.Split(text, "\n")
	list := make([][]int, 0)
	for _, token := range tokens {
		strNums := strings.Split(token, " ")
		curr := make([]int, 0)
		for _, s := range strNums {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic("could not parse input")
			}
			curr = append(curr, n)
		}
		list = append(list, curr)
	}
	return list
}

func isRowSorted(row []int) bool {
	sortedRow := slices.Clone(row)
	slices.Sort(sortedRow)
	if !slices.Equal(row, sortedRow) {
		slices.Reverse(sortedRow)
		if !slices.Equal(row, sortedRow) {
			return false
		}
	}
	return true
}

func validateRow(row []int) bool {
	if !isRowSorted(row) {
		return false
	}
	lower := 1.0
	upper := 3.0
	for i := 0; i < len(row) - 1; i++ {
		diff := math.Abs(float64(row[i] - row[i+1]))
		if diff < lower || diff > upper {
			return false
		}
	}
	return true
}

func solvePart1(text string) int {
	l := parseInput(text)
	c := 0
	for i := range l {
		if valid := validateRow(l[i]); valid {
			c++
		}
	}
	return c
}

func validateRowWithSkips(row []int) bool {
	// start skipping columns
	for i := 0; i < len(row); i++ {
		candidate := append(append([]int{}, row[:i]...), row[i+1:]...)
		if validateRow(candidate) {
			return true
		}
	}
	return false
}

func solvePart2(text string) int {
	l := parseInput(text)
	c := 0
	for i := range l {
		if valid := validateRowWithSkips(l[i]); valid {
			c++
		}
	}
	return c
}

func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
	ans2 := solvePart2(input)
	fmt.Println(ans2)
}
