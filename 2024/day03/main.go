package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func solvePart2(input string) int {
	patMul := `mul\(([0-9]+,[0-9]+)\)`
	patDo := `do\(\)`
	patDont := `don't\(\)`
	
	reCombined := regexp.MustCompile(patMul + "|" + patDo + "|" + patDont)
	matches := reCombined.FindAllStringSubmatch(input, -1)
	flag := 0 
	ans := 0
	for _, match := range matches {
		if match[0] == "do()" {
			flag = 0
			continue
		}
		if match[0] == "don't()" {
			flag = 1
			continue
		}
		if flag == 0 {
			ans += doMul(match[1])
		}
		
	}
	return ans
}

func doMul(candidate string) int {
	s1, s2, _ := strings.Cut(candidate, ",")
	n1, err := strconv.Atoi(s1)
	if err != nil {
		panic("Invalid input")
	}
	n2, err := strconv.Atoi(s2)
	if err != nil {
		panic("Invalid input")
	}
	return n1 * n2
}

func solvePart1(input string) int {
	re := regexp.MustCompile(`mul\(([0-9]+,[0-9]+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	ans := 0
	for _, match := range matches {
		candidate := match[1]
		ans += doMul(candidate)
	}
	return ans
}


func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
	ans2 := solvePart2(input)
	fmt.Println(ans2)
}
