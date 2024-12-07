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
// 0,1,2,3 - N,E,S,W
var dirs = []Pos{{-1,0}, {0,1}, {1,0}, {0,-1}}


func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func isOutside(i,j,x,y int) bool {
	if i < 0 || i >= x || j < 0 || j >= y {
		return true
	}
	return false
}

// track curr dir and decide next idx based on
// that - simulate 
func solvePart1(input string) int {
	grid := parseInput(input)
	n, m := len(grid), len(grid[0])
	currDir := 0
	currPos := Pos{}
	
	for i := 0; i < n; i++ {
		j := strings.IndexByte(grid[i], '^')
		if j != -1 {
			currPos.x, currPos.y = i, j
			break
		}	
	}
	visited := make(map[Pos]struct{})
	for !isOutside(currPos.x, currPos.y, n, m) {
		if grid[currPos.x][currPos.y] != '#' {
			visited[currPos] = struct{}{}
		}
		if grid[currPos.x][currPos.y] == '#' {
			prevDir := currDir
			currDir = (currDir + 1) % 4
			currPos.x -= dirs[prevDir].x
			currPos.y -= dirs[prevDir].y
		}
		currPos.x += dirs[currDir].x
		currPos.y += dirs[currDir].y
	}
	return len(visited)
}

func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
}