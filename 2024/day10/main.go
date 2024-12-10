package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var directions = []int{0, 1, 0, -1, 0}

type pos struct {
	x, y int
}

func parseInput(input string) [][]int8 {
	rows := strings.Split(input, "\n")
	grid := make([][]int8 , len(rows))
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			s, _ := strconv.Atoi(string(rows[i][j]))
			grid[i] = append(grid[i], int8(s))
		}
	}
	return grid
}

func isOutside(i, j, n, m int) bool {
	return i < 0 || i >= n || j < 0 || j >= m
}


func explorePaths(grid [][]int8, prev int8, seen map[pos]bool, part2 bool, i, j int) int {
	n := len(grid)
	m := len(grid[0])
	if grid[i][j] == 9 {
		return 1
	}
	numPaths := 0
	for idx := 0; idx < 4; idx++ {
		x := i + directions[idx]
		y := j + directions[idx + 1]
		if !isOutside(x, y, n, m) && grid[x][y] == prev + 1 { 
			if !part2 {
				nextPos := pos{x,y}
				if ok := seen[nextPos]; !ok {
					seen[nextPos] = true
					numPaths += explorePaths(grid, prev+1, seen, part2, x, y)
				}
			} else {
				numPaths += explorePaths(grid, prev+1, seen, part2, x, y)
			}
			
		}
	}
	return numPaths
}

func solveParts(input string, part2 bool) int {
	grid := parseInput(input)
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				seen := make(map[pos]bool)
				ans += explorePaths(grid, 0, seen, part2, i, j)
			}
		}
	}
	return ans
}

func main() {
	ans1 := solveParts(input, false)
	fmt.Println(ans1)
	ans2 := solveParts(input, true)
	fmt.Println(ans2)
}