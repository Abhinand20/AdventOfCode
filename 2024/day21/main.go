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

// Set of unique sequences

type sequence struct {
	i, j, k, l int
}

func (s *sequence) updateIndex(i , val int) {
	switch i {
	case 0: s.i = val
	case 1: s.j = val
	case 2: s.k = val
	case 3: s.l = val
	default:
		return
	}
}
type delta map[sequence]int

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

func getDelta(price []int) delta {
	deltas := make(delta)
	diffs := make([]int, 0)
	for i := 1; i < len(price); i++ {
		diffs = append(diffs, price[i] - price[i-1])
	}
	for j := range len(diffs) - 3 {
		seq := &sequence{}
		for idx := range 4 {
			seq.updateIndex(idx, diffs[j + idx])
		}
		if _, ok := deltas[*seq]; !ok {
			deltas[*seq] = price[j + 4]
		}
		seq = &sequence{}
	}
	return deltas
}


func getMaxValue(deltas []delta) int {
	ans := 0.0
	allSeqs := make(map[sequence]bool)
	for _, d := range deltas {
		for k, _ := range d {
			allSeqs[k] = true
		}
	}
	for k, _ := range allSeqs {
		currAns := 0.0
		for j := 0; j < len(deltas); j++ {
			b := deltas[j]
			if val, found := b[k]; found {
				currAns += float64(val)
				continue
			}
		}
		ans = math.Max(ans, currAns)
	}
	return int(ans)
}

func solvePart2(input string) int {
	arr := parseInput(input)
	priceDeltas := make([]delta, 0)
	for _, a := range arr {
		prices := make([]int, 0)
		n := a
		prices = append(prices, a % 10)
		for _ = range 2000 {
			n = transform(n)
			lastDigit := n % 10
			prices = append(prices, lastDigit)
		}
		// For each price array get the changes
		priceDeltas = append(priceDeltas, getDelta(prices))
	}
	return getMaxValue(priceDeltas)
}


func main() {
	// ans1 := solvePart1(input)
	// fmt.Println(ans1)
	ans2 := solvePart2(input)
	fmt.Println(ans2)
} 