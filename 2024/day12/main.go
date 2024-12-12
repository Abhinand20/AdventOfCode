package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Pos struct {
	x, y int
}


func parseInput(input string) []string {
	return strings.Split(input, "\n")
}


func computeAreaPerimeter(grid []string, p Pos, prev byte, visited map[Pos]bool) (int, int) {
	return 0, 0
}


func solvePart1(input string) int {
	grid := parseInput(input)
	visited := make(map[Pos]bool)
	n, m := len(grid), len(grid[0])
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cp := Pos{i, j}
			if ok := visited[cp]; !ok {
				a, p := computeAreaPerimeter(grid, cp, grid[i][j], visited)
				ans += a * p
			}
		}
	}
}


func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
}