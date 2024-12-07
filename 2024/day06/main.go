package main

import (
	_ "embed"
	"fmt"
	"strings"
	"sync"
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

func findCycle(currPos Pos, grids <-chan []string, results chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for grid := range grids {
		_, cycle := getPath(grid, currPos)
		results <- cycle
	}
}


func solvePart2Concurrent(input string) int {
	defer timeIt("solvePart2Concurrent")()
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
	
	const numWorkers = 64
	var numTasks int = len(path) - 1  // For each possible obstacle
	
	// Setup go routines and channels
	wg := sync.WaitGroup{}
	results := make(chan bool, numTasks)
	grids := make(chan []string, numTasks)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		// Try to find a cycle for each updated grid
		go findCycle(currPos, grids ,results, &wg)
	}

	// Send a new grid to each worker
	for p := range path {
		if p == currPos {
			continue
		}
		b := []byte(grid[p.x])
		c := b[p.y]
		b[p.y] = '#'
		grid[p.x] = string(b)
		grids <- append([]string{}, grid...)
		// Do we need to reset it again?
		b[p.y] = c
		grid[p.x] = string(b)
	}
	close(grids)
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Check how many grids led to a cycle
	for r := range results {
		if r {
			ans++
		}
	}
	return ans
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
	ans2 := solvePart2Concurrent(input)
	fmt.Println(ans2)
}