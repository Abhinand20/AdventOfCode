package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string


func parseInput(input string) (map[int][]int, [][]int) {
	tokens := strings.Split(input, "\n\n")
	rules := make(map[int][]int)
	orderTokens := strings.Split(tokens[1], "\n")
	orders := make([][]int, len(orderTokens))
	for i := range orderTokens {
		currOrder := strings.Split(orderTokens[i], ",")
		for _, o := range currOrder {
			n, _ := strconv.Atoi(o)
			orders[i] = append(orders[i], n)
		}
	}
	ruleTokens := strings.Split(tokens[0], "\n")
	for _, tok := range ruleTokens {
		s1, s2, _ := strings.Cut(tok, "|")
		n1, _ := strconv.Atoi(s1)
		n2, _ := strconv.Atoi(s2)
		rules[n1] = append(rules[n1], n2)
	}
	return rules, orders
}


func isCorrect(r map[int][]int, order []int) bool {
	for i := 0; i < len(order) - 1; i++ {
		candidates := r[order[i]]
		for j := i + 1; j < len(order); j++ {
			if !slices.Contains(candidates, order[j]) {
				return false
			}
		}
	}
	return true
}

func matchRules(r map[int][]int, o [][]int) int {
	ans := 0
	for i := range o {
		if isCorrect(r, o[i]) {
			midIdx := int(len(o[i]) / 2)
			ans += o[i][midIdx]
		}
	}
	return ans
}

func swapInplace(r map[int][]int, elems []int) {
	l := len(elems)
	for i := 0; i < l - 1; i++ {
		candidates, ok := r[elems[i]]
		// If this is the greatest element, we swap with the last one
		if !ok {
			elems[i], elems[l-1] = elems[l-1], elems[i]
			return
		}
		for j := i + 1; j < l; j++ {
			if !slices.Contains(candidates, elems[j]) {
				elems[i], elems[j] = elems[j], elems[i]
				return
			}
		}
	}
}

func fixOrdering(r map[int][]int, order []int) []int {
	// Swap elements, and check
	fixed := slices.Clone(order)
	for !isCorrect(r, fixed) {
		swapInplace(r, fixed)
	}
	return fixed
}

func matchRules2(r map[int][]int, o [][]int) int {
	ans := 0
	for i := range o {
		if !isCorrect(r, o[i]) {
			midIdx := int(len(o[i]) / 2)
			fixedOrder := fixOrdering(r, o[i])
			ans += fixedOrder[midIdx]
		}
	}
	return ans
}

func solvePart1(input string) int {
	rules, order := parseInput(input)
	return matchRules(rules, order)
}

func solvePart2(input string) int {
	rules, order := parseInput(input)
	return matchRules2(rules, order)
}


func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
	ans2 := solvePart2(input)
	fmt.Println(ans2)
}