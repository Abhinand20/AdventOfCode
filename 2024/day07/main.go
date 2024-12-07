package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseInput(input string) ([]int, [][]int) {
	tokens := strings.Split(input, "\n")
	tests := make([]int, len(tokens))
	candidates := make([][]int, len(tokens))
	for i, t := range tokens {
		l,r,_:= strings.Cut(t, ":")
		li, _ := strconv.Atoi(l)
		tests[i] = li
		for _, s := range strings.Split(strings.TrimSpace(r), " ") {
			si, _ := strconv.Atoi(s)
			candidates[i] = append(candidates[i], si)
		}
	}
	return tests, candidates
}

// Check all possible combinations if it matches t
func isValid(seq []int, part2 bool, sum, i, t int) bool {
	if i == len(seq) {
		return sum == t
	}
	if !part2 {
		return isValid(seq, false, sum + seq[i], i + 1, t) || isValid(seq, false, sum * seq[i], i + 1, t)
	}
	n := len(strconv.Itoa(seq[i]))
	combined := sum * int(math.Pow10(n)) + seq[i] 
	return isValid(seq, true, sum+seq[i], i+1, t) || isValid(seq, true, sum*seq[i], i+1, t) || isValid(seq, true, combined, i+1, t)
}

func solveParts(input string, part2 bool) int {
	tests, candidates := parseInput(input)
	ans := 0
	for i := range tests {
		if isValid(candidates[i], part2, candidates[i][0], 1, tests[i]) {
			ans += tests[i]
		}
	}
	return ans
}

func main() {
	ans1 := solveParts(input, false)
	fmt.Println(ans1)
	ans2 := solveParts(input, true)
	fmt.Println(ans2)
}