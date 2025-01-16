package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string


func parseInput(input string) []int {
	var ret []int
	for _, s := range strings.Split(input, "\n") {
		n, _ := strconv.Atoi(s)
		ret = append(ret, n)
	}
	return ret
}

func mixPrune(s, n int) int {
	m := 16777216
	s = s ^ n
	return s & (m - 1)
}

func transform(n int) int {
	n1 := n << 6
	n1 = mixPrune(n, n1)
	n2 := n1 >> 5
	n2 = mixPrune(n1, n2)
	n3 := n2 << 11
	return mixPrune(n2, n3)
}

func solvePart1(input string) int {
	arr := parseInput(input)
	ans := 0
	for _, a := range arr {
		n := a
		for _ = range 2000 {
			n = transform(n)
		}
		fmt.Println(a, n)
		ans += n
	}
	return ans
}

func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
} 