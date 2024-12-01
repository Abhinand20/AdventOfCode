package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func PrintGrid(arr []string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func ProcessInputD11() ([]string, error) {
	const inputFile string = "../inputs/day11_1.txt"
	file, err := os.Open(inputFile)
	inputs := make([]string, 0)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			inputs = append(inputs, line)
		}
	}
	return inputs, nil
}

func expandGrid(arr []string) ([]int, []int) {
	m, n := len(arr), len(arr[0])
	offsetRows := make([]int, m)
	offsetCols := make([]int, n)

	for i := 0; i < m; i++ {
		repeat := true
		for j := 0; j < n; j++ {
			if arr[i][j] == '#' {
				repeat = false
			}
		}
		if repeat {
			offsetRows[i+1] = 1 // Borders are not handled for simplicity
		} 
	}

	for j := 0; j < n; j++ {
		repeat := true
		for i := 0; i < m; i++ {
			if arr[i][j] == '#' {
				repeat = false
			}
		}
		if repeat {
			offsetCols[j+1] = 1
		}
	}
	for i := 1; i < m; i++ {
		offsetRows[i] += offsetRows[i-1]
	}
	for i := 1; i < n; i++ {
		offsetCols[i] += offsetCols[i-1]
	}

	return offsetRows, offsetCols
}

func getCordinates(arr []string, or, oc []int, scale int) []position {
	m, n := len(arr), len(arr[0])
	points := make([]position, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] == '#' {
				offset:= position{i + scale * or[i], j + scale * oc[j]}
				points = append(points, offset)
			}
		}
	}
	return points
}

func getDistance(allPoints []position) int {
	ans := 0.0
	for i := 0; i < len(allPoints) - 1; i++ {
		for j := i+1; j < len(allPoints); j++ {
			p := allPoints[i]
			q := allPoints[j]
			ans += math.Abs(float64(p.x) - float64(q.x)) + math.Abs(float64(p.y) - float64(q.y))
		}
	}
	return int(ans)
}

func SolveDay11() {
	inputs, _ := ProcessInputD11()
	scale := 1000000 - 1
	offsetRows, offsetCols := expandGrid(inputs)
	allPoints := getCordinates(inputs, offsetRows, offsetCols, scale)
	ans := getDistance(allPoints) 
	fmt.Println(ans)
}