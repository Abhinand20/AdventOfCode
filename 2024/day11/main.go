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

func parseInput(input string) []int {
	s := strings.Split(input, " ")
	res := make([]int, len(s))
	for i := range s {
		n, _ := strconv.Atoi(s[i])
		res[i] = n
	}
	return res
}

// 2024 (1 -> 4) : 2 blinks
// 4048 (1 -> 4) : 4 blinks
// 40 48 -> 4 0 4 8 -> 4 * 2024 1 4 * 2024 8 * 2024

func doUpdate(arr []int) []int {
	// Follow rules and return new arr
	res := make([]int, 0)
	for _, v := range arr {
		if v == 0 {
			res = append(res, 1)
			continue
		}
		numDigits := int(math.Floor(math.Log10(float64(v)))) + 1
		if numDigits % 2 != 0 {
			res = append(res, v * 2024)
			continue
		}
		n := int(math.Pow10(numDigits / 2))
		l := int(math.Floor(float64(v / n)))
		r := v % n
		res = append(res, []int{l, r}...)
	}
	return res
}


func solvePart1(input string) int {
	candidates := parseInput(input)
	var ans []int = candidates
	for i := 0; i < 75; i++ {
		ans = doUpdate(ans)
	}
	return len(ans)
}

func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
}