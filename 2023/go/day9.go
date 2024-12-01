package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func ProcessInputD9() ([][]int, error) {
	const inputFile string = "../inputs/day9_1.txt"
	file, err := os.Open(inputFile)
	inputs := make([][]int, 0)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			var tokens []string = strings.Split(line, " ")
			tokensInt := make([]int, len(tokens))
			for i, t := range tokens {
				val, ok := strconv.Atoi(t)
				if ok != nil {
					return inputs, ok
				}
				tokensInt[i] = val
			}
			inputs = append(inputs, tokensInt)
		}	
	}
	return inputs, nil
}

func PredictNextValue(vals []int, is_hist bool) int {
	lastVals := make([]int, 0)
	currSeq := make([]int, len(vals))
	copy(currSeq, vals)
	v := vals[0]
	if !is_hist {
		v = vals[len(vals) - 1]
	}
	lastVals = append(lastVals, v)
	run := true
	for run {
		tempSeq := make([]int, 0)
		target := currSeq[1] - currSeq[0]
		run = false
		for i := 1; i < len(currSeq); i++ {
			diff := currSeq[i] - currSeq[i-1]
			tempSeq = append(tempSeq, diff)
			if diff != target {
				run = true
			}
		}
		appendVal := tempSeq[0]
		if !is_hist {
			appendVal = tempSeq[len(tempSeq) - 1]
		}
		lastVals = append(lastVals, appendVal)
		currSeq = tempSeq
	}
	ans := lastVals[len(lastVals) - 1]
	for j := len(lastVals) - 2; j >= 0; j-- {
		if !is_hist {
			ans += lastVals[j]
		} else {
			ans = lastVals[j] - ans
		}
	}

	return ans
}

func SolveDay9Part1(arr [][]int) int {
	ans := 0
	for i := 0; i < len(arr); i++ {
		ans += PredictNextValue(arr[i], false)
	}
	return ans
}

func SolveDay9Part2(arr [][]int) int {
	ans := 0
	for i := 0; i < len(arr); i++ {
		ans += PredictNextValue(arr[i], true)
	}
	return ans
}

func SolveDay9() {
	inputs, err := ProcessInputD9()
	if err != nil {
		return
	}
	ans := SolveDay9Part1(inputs)
	fmt.Println(ans)
	ans = SolveDay9Part2(inputs)
	fmt.Println(ans)
}