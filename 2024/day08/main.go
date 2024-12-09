package main

import (
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input string

func parseInput(input string) []int {
	expanded := make([]int, 0)
	for i := 0; i < len(input); i++ {
		n, _ := strconv.Atoi(string(input[i]))
		fmt.Println(n)
	}
	return expanded
}

func solvePart1(input string) int {
	expanded := parseInput(input)
	return 0
}

func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
}