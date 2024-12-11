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
const numIter = 75

type State struct {
	num, iter int
}

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

// Store distinct {n, iter} : blinks map as a cache
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


func doSingleUpdate(cState State, cache map[State]int) int {
	if cState.iter >= numIter {
		return 0
	}
	if _, ok := cache[cState]; ok {
		return cache[cState]
	}
	num := cState.num
	if num == 0 {
		nState := State{1, cState.iter + 1}
		cache[cState] = doSingleUpdate(nState, cache)
		return cache[cState]
	}
	
	numDigits := int(math.Floor(math.Log10(float64(num)))) + 1
	if numDigits % 2 != 0 {
		nState := State{num * 2024, cState.iter + 1}
		cache[cState] = doSingleUpdate(nState, cache)
		return cache[cState]
	}
	n := int(math.Pow10(numDigits / 2))
	l := int(math.Floor(float64(num / n)))
	r := num % n
	nState1 := State{l, cState.iter + 1}
	nState2 := State{r, cState.iter + 1}
	cache[cState] = 1 + doSingleUpdate(nState1, cache) + doSingleUpdate(nState2, cache)
	return cache[cState]
}



func solvePart1(input string) int {
	candidates := parseInput(input)
	var ans []int = candidates
	for i := 0; i < numIter; i++ {
		ans = doUpdate(ans)
	}
	return len(ans)
}

func solvePart2(input string) int {
	candidates := parseInput(input)
	cache := make(map[State]int)
	ans := 0
	for i := 0; i < len(candidates); i++ {
		s := State{candidates[i], 0}
		ans += (1 + doSingleUpdate(s, cache))
	}
	return ans
}

func main() {
	ans1 := solvePart2(input)
	fmt.Println(ans1)
}