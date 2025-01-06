package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type lookup map[rune][]string

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

func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
}