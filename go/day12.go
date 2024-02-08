package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ProcessInputD12() ([]string, [][]int) {	
	const inputFile string = "../inputs/day12_1.txt"
	file, _ := os.Open(inputFile)
	inputs := make([]string, 0)
	reference := make([][]int, 0)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			tokens := strings.Split(line, " ")
			inputs = append(inputs, tokens[0])
			refs := make([]int, 0)
			for _, s := range strings.Split(tokens[1], ",") {
				val, _ := strconv.Atoi(s) 
				refs = append(refs, val)
			}
			reference = append(reference, refs)
		}
	}
	return inputs, reference
}

func valid(arr string, ref []int) bool {
	seen := make([]int, 0)
	curr := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == '.' {
			if curr > 0 {
				seen = append(seen, curr)
			}
			curr = 0
		}
		if arr[i] == '#' {
			curr++
		}
	}
	if curr > 0 {
		seen = append(seen, curr)
	}
	return slices.Equal(seen, ref)
}

func allCombs(arr string, ref []int, i int) int {
	if i == len(arr) {
		if valid(arr, ref) {
			return 1
		}
		return 0
	}

	if arr[i] == '?' {
		return allCombs(arr[:i] + "#" + arr[i+1:], ref, i+1) + 
		allCombs(arr[:i] + "." + arr[i+1:], ref, i+1)
	}
	
	return allCombs(arr, ref, i+1)
}

func solve1(inputs []string, refs [][]int) int {
	ans := 0
	for i := 0; i < len(inputs); i++ {
		ans += allCombs(inputs[i], refs[i], 0)
	}
	return ans
}

func SolveDay12() {
	inputs, references := ProcessInputD12()
	ans := solve1(inputs, references)
	fmt.Println(ans)
}
