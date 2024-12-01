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

func parseInput(text string) ([]int, []int) {
	tokens := strings.Split(text, "\n")
	list1 := make([]int, len(tokens))
	list2 := make([]int, len(tokens))
	for _, token := range tokens {
		s1, s2, _ := strings.Cut(token, " ")
		i1, err := strconv.Atoi(strings.TrimSpace(s1))
		if err != nil {
			panic("Invalid input")
		}
		list1 = append(list1, i1)
		i2, err := strconv.Atoi(strings.TrimSpace(s2))
		if err != nil {
			panic("Invalid input")
		}
		list2 = append(list2, i2)
	}
	return list1, list2
}

func solvePart1(text string) int {
	l1, l2 := parseInput(text)
	slices.Sort(l1)
	slices.Sort(l2)
	ans := 0.0
	for i := range l1 {
		ans += (math.Abs(float64(l1[i] - l2[i])))
	}
	return int(ans)
}

func getCounts(l []int) map[int]int {
	res := make(map[int]int)
	for _, k := range l {
		res[k] += 1
	}
	return res
}

func solvePart2(text string) int {
	l1, l2 := parseInput(text)
	c1 := getCounts(l1)
	c2 := getCounts(l2)

	ans := 0
	for k, freq := range c1 {
		factor := c2[k]
		ans += k * freq * factor
	}
	return ans
}


func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
	ans2 := solvePart2(input)
	fmt.Println(ans2)
}
