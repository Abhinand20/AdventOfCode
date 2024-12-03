package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func solvePart2(input string) int {
	reMul := regexp.MustCompile(`mul\(([0-9]+,[0-9]+)\)`)
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)
	matchesMul := reMul.FindAllStringSubmatch(input, -1)
	indexesMul := reMul.FindAllStringIndex(input, -1) // [ [start1, end1], ... ]
	indexesDo := reDo.FindAllStringIndex(input, -1)
	indexesDont := reDont.FindAllStringIndex(input, -1)

	combined := make([]int, 0)
	lookupDo := make(map[int]bool)
	for _,idx := range indexesDo {
		lookupDo[idx[0]] = true
		combined = append(combined, idx[0])
	}
	lookupDont := make(map[int]bool)
	for _,idx := range indexesDont {
		lookupDont[idx[0]] = true
		combined = append(combined, idx[0])
	}
	lookupMul := make(map[int]int)
	for i,idx := range indexesMul {
		// Map start index to list index
		lookupMul[idx[0]] = i
		combined = append(combined, idx[0])
	}
	slices.Sort(combined)
	flag := 0 // 0 - Do
	ans := 0
	for _, currIdx := range combined {
		if ok := lookupDont[currIdx]; ok {
			flag = 1
			continue
		}
		if ok := lookupDo[currIdx]; ok {
			flag = 0
			continue
		}
		if flag == 0 {
			candidate := matchesMul[lookupMul[currIdx]][1]
			ans += doMul(candidate)
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
