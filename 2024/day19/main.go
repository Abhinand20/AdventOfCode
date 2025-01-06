package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type lookup map[rune][]string
type cacheKey struct {
	s string
	idx int
}

func parseInput(input string) (lookup, []string) {
	substr := make(lookup)
	s, t, _ := strings.Cut(input, "\n\n")
	for _, tok := range strings.Split(s, ", ") {
		ch := rune(tok[0])
		substr[ch] = append(substr[ch], tok)
	}
	return substr, strings.Split(t, "\n")
}

func canMake(s lookup, t string, i int) bool {
	if i >= len(t) {
		return true
	}
	ch := rune(t[i])
	candidates, found := s[ch]
	if  !found {
		return false
	}
	for _, c := range candidates {
		n := len(c)
		if i+n <= len(t) && t[i:i+n] == c {
			if valid := canMake(s, t, i+n); valid {
				return true
			}
		}
	}
	return false
}

func numWays(s lookup, t string, i int, cache map[cacheKey]int) int {
	if i >= len(t) {
		return 1
	}
	ck := cacheKey{t, i}
	if val, found := cache[ck]; found {
		return val
	}
	ch := rune(t[i])
	candidates, found := s[ch]
	if  !found {
		return 0
	}
	ways := 0
	for _, c := range candidates {
		n := len(c)
		if i+n <= len(t) && t[i:i+n] == c {
			ways += numWays(s, t, i+n, cache)
		}
	}
	cache[ck] = ways
	return ways
}

func solvePart1(input string) int {
	substr, targets := parseInput(input)
	ans := 0
	for _, t := range targets {
		if canMake(substr, t, 0) {
			ans++
		}
	}
	return ans
}

func solvePart2(input string) int {
	substr, targets := parseInput(input)
	ans := 0
	cache := make(map[cacheKey]int)
	for _, t := range targets {
		ans += numWays(substr, t, 0, cache)
	}
	return ans
}

func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
	ans2 := solvePart2(input)
	fmt.Println(ans2)
}