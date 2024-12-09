package main

import (
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input string

func parseInput(input string) []int {
	expanded := make([]int, 0)
	count := 0
	for i := 0; i < len(input); i++ {
		n, _ := strconv.Atoi(string(input[i]))
		r := count
		if i % 2 != 0{
			r = -1
			count--
		}
		for k := 0; k < n; k++ {
			expanded = append(expanded, r)
		}
		count++
	}
	return expanded
}

func checksum(arr []int) int {
	ans := 0
	for i, v := range arr {
		if v != -1 {
			ans += i * v
		}
	}
	return ans
}

func doCompression(expanded []int) {
	j, i := -1, -1 
	for idx, v := range expanded {
		if v != -1 {
			j = idx
		} 
		if i == -1 && v == -1 {
			i = idx
		}
	}
	for i < j {
		// swap
		expanded[i], expanded[j] = expanded[j], expanded[i]
		j--
		i++
		for j >= 0 && expanded[j] == -1 {
			j--
		}
		for i < len(expanded) && expanded[i] != -1 {
			i++
		}
	}
}

func tryFillFile(arr []int, fs, j int) {
	for i := 0; i < len(arr); i++ {
		if i >= j {
			return 	
		}
		if arr[i] == -1 {
			if i + fs < len(arr) {
				fill := true
				for k := 0; k < fs; k++ {
					if arr[i+k] != -1 {
						fill = false
					}
				}
				if fill {
					for k := 0; k < fs; k++ {
						arr[i+k], arr[j+k] = arr[j+k], arr[i+k]
					}
					return
				}
			}
		}
	}
}


func doCompressionFile(arr []int) {
	j, i := -1, -1 
	for idx, v := range arr {
		if v != -1 {
			j = idx
		} 
		if i == -1 && v == -1 {
			i = idx
		}
	}
	for j > 0 {
		if i > j {
			break
		}
		for j >= 0 && arr[j] == -1 {
			j--
		}
		// Get file size
		f := arr[j]
		fs := 0
		for j >= 0 && arr[j] == f {
			fs++
			j--
		}
		// Now,
		tryFillFile(arr, fs, j+1)
	}
}

func solveParts(input string, part2 bool) int {
	expanded := parseInput(input)
	if !part2 {
		doCompression(expanded)
	} else
	{
		doCompressionFile(expanded)
	}
	return checksum(expanded)
}

func main() {
	// ans1 := solveParts(input, false)
	// fmt.Println(ans1)
	ans2 := solveParts(input, true)
	fmt.Println(ans2)
}