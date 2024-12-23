package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var directions = []int{0, 1, 0, -1, 0}

type Pos struct {
	x, y int
}


func parseInput(input string) []string {
	return strings.Split(input, "\n")
}


func isOutside(p Pos, n, m int) bool {
	return p.x < 0 || p.x >= n || p.y < 0 || p.y >= m
}

func computeArea(grid []string, p Pos, prev byte, visited map[Pos]bool) int {
	n := len(grid)
	m := len(grid[0])
	area := 1
	for i := 0; i < 4; i++ {
		np := Pos{p.x + directions[i], p.y + directions[i+1]}
		ok := visited[np]
		if !isOutside(np, n, m) && prev == grid[np.x][np.y] && !ok {
			visited[np] = true
			area += computeArea(grid, np, prev, visited)
		}
	}
	return area
}

func computePerimeter(grid []string, p Pos, prev byte, visited map[Pos]bool) int {
	perimeter := 4
	n := len(grid)
	m := len(grid[0])
	for i := 0; i < 4; i++ {
		np := Pos{p.x + directions[i], p.y + directions[i+1]}
		if !isOutside(np, n, m) && prev == grid[np.x][np.y] {
			perimeter -= 1
		}
	}
	for i := 0; i < 4; i++ {
		np := Pos{p.x + directions[i], p.y + directions[i+1]}
		ok := visited[np]
		if !isOutside(np, n, m) && prev == grid[np.x][np.y] && !ok {
			visited[np] = true
			perimeter += computePerimeter(grid, np, prev, visited)
		}
	}
	return perimeter
}

func isExternalCorner(grid []string, prev byte, p1, p2, d Pos) {
	n := len(grid)
	m := len(grid[0])
	c1 := false
	if isOutside()
}


func computeSides(grid []string, p Pos, prev byte, visited map[Pos]bool) int {
	corners := 0
	n := len(grid)
	m := len(grid[0])
	u := Pos{p.x - 1, p.y}
	d := Pos{p.x + 1, p.y}
	l := Pos{p.x, p.y - 1}
	r := Pos{p.x, p.y + 1}
	// Check external corner in all 4 directions
	if isExternalCorner(grid, prev, u, l, Pos{p.x - 1, p.y - 1}) {
		corners++
	}
	if isExternalCorner(grid, prev, u, r, Pos{p.x - 1, p.y + 1}) {
		corners++
	}
	if isExternalCorner(grid, prev, d, l, Pos{p.x + 1, p.y - 1}) {
		corners++
	}
	if isExternalCorner(grid, prev, d, r, Pos{p.x + 1, p.y + 1}) {
		corners++
	}
	// Check internal corner
	if isInternalCorner(u, l, Pos{p.x - 1, p.y - 1}) {
		corners++
	}
	if isInternalCorner(u, r, Pos{p.x - 1, p.y + 1}) {
		corners++
	}
	if isInternalCorner(d, l, Pos{p.x + 1, p.y - 1}) {
		corners++
	}
	if isInternalCorner(d, r, Pos{p.x + 1, p.y + 1}) {
		corners++
	}
	// Do it for all others

	return corners
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
				visited[cp] = true
				temp := make(map[Pos]bool)
				for k, v := range visited {
					temp[k] = v
				}
				a := computeArea(grid, cp, grid[i][j], temp)
				fmt.Println(string(grid[i][j]), a)
				p := computePerimeter(grid, cp, grid[i][j], visited)
				fmt.Println(string(grid[i][j]), p)
				ans += a * p
			}
		}
	}
	return ans
}

func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
}