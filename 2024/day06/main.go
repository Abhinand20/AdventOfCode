package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string


type Pos struct {
	x, y int
}

type PosDir struct {
	p Pos
	d int
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

func getPath(grid []string, currPos Pos) (map[Pos]struct{}, bool) {
	currDir := 0
	n, m := len(grid), len(grid[0])
	visited := make(map[Pos]struct{})
	states := make(map[PosDir]struct{})
	for !isOutside(currPos.x, currPos.y, n, m) {
		currState := PosDir{currPos, currDir}
		// Found a cycle
		if _, found := states[currState]; found {
			return nil, true
		}
		if grid[currPos.x][currPos.y] != '#' {
			visited[currPos] = struct{}{}
		}
		if grid[currPos.x][currPos.y] == '#' {
			states[currState] = struct{}{}
			prevDir := currDir
			currDir = (currDir + 1) % 4
			currPos.x -= dirs[prevDir].x
			currPos.y -= dirs[prevDir].y
		}
		currPos.x += dirs[currDir].x
		currPos.y += dirs[currDir].y
	}
	return visited, false
}


// track curr dir and decide next idx based on
// that - simulate 
func solvePart1(input string) int {
	defer timeIt("solvePart1")()
	grid := parseInput(input)
	currPos := Pos{}
	for i := 0; i < len(grid); i++ {
		j := strings.IndexByte(grid[i], '^')
		if j != -1 {
			currPos.x, currPos.y = i, j
			break
		}	
	}
	visited, _ := getPath(grid, currPos)
	return len(visited)
}

// Get all distinct positions in the path
// Place obstacle in each and then check for loop 
func solvePart2(input string) int {
	defer timeIt("solvePart2")()
	ans := 0
	grid := parseInput(input)
	currPos := Pos{}
	for i := 0; i < len(grid); i++ {
		j := strings.IndexByte(grid[i], '^')
		if j != -1 {
			currPos.x, currPos.y = i, j
			break
		}	
	}
	path, _ := getPath(grid, currPos)
	for p := range path {
		if p == currPos {
			continue
		}
		b := []byte(grid[p.x])
		c := b[p.y]
		b[p.y] = '#'
		grid[p.x] = string(b)
		if _, cycle := getPath(grid, currPos); cycle {
			ans++
		}
		b[p.y] = c
		grid[p.x] = string(b)
	}

	return ans
}


func timeIt(name string) func() {
	start := time.Now()
	return func(){ 
		fmt.Printf("[%s] Time elapsed: %dus\n", name, time.Since(start).Microseconds())
	}
}

func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
	ans2 := solvePart2(input)
	fmt.Println(ans2)
}