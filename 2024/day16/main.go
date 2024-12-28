package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

type Pos struct {
	x, y int
}

type cacheKey struct {
	dir int
	pos Pos
}

const (
	EAST = iota
	SOUTH
	WEST
	NORTH
)

// E, S, W, N
var dirs = [4][2]int{{0,1}, {1,0}, {0,-1}, {-1,0}}
var cache = make(map[cacheKey]float64)
var answer = math.MaxFloat64


func RenderGrid(grid [][]rune) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println()
	}
}

func parseInput(input string) [][]rune {
	rows := strings.Split(input, "\n")
	grid := make([][]rune, len(rows))
	for i, row := range rows {
		grid[i] = make([]rune, len(row))
		for j, c := range row {
			grid[i][j] = c
		}
	}
	return grid
}

func nextPos(p Pos, d int) Pos {
	nextDir := dirs[d]
	return Pos{
		p.x + nextDir[0], 
		p.y + nextDir[1],
	}
}

func turnAntiClockwise( p Pos, currDir int) (Pos, int) {
	i := ((currDir - 1) + 4) % 4 
	return nextPos(p, i), i
}

func turnClockwise(p Pos, currDir int) (Pos, int) {
	i := (currDir + 1) % 4
	return nextPos(p, i), i
}


func findMinScore(grid [][]rune, visited [][]bool, s Pos, d int, score float64, memo map[cacheKey]float64) float64 {
	if grid[s.x][s.y] == 'E' {
		return score
	}
	if grid[s.x][s.y] == '#' || visited[s.x][s.y] || score >= answer {
		return math.MaxFloat64
	}
	state := cacheKey{d, s}
    if val, found := memo[state]; found {
        return val
    }
	// Add it to visited
	visited[s.x][s.y] = true
	// Check other options
	np := nextPos(s, d)
	npc, dc := turnClockwise(s, d)
	npa, da := turnAntiClockwise(s, d)
	// Did we reach the dest?
	ans := findMinScore(grid, visited, np, d, 1 + score, memo)
	ans = math.Min(ans, findMinScore(grid, visited, npc, dc, 1001 + score, memo))
	ans = math.Min(ans, findMinScore(grid, visited, npa, da, 1001 + score, memo))
	visited[s.x][s.y] = false
	answer = math.Min(answer, ans)
	// memo[state] = ans
	return ans
}

// 3 choices - straight (+1), turn clockwise (+1000), turn anticlockwise (+1000) 
// DFS and find the minimum cost
func solvePart1(input string) int {
	grid := parseInput(input)
	startPos := Pos{}
	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'S' {
				startPos.x, startPos.y = i, j
			}
		}
	}
	memo := make(map[cacheKey]float64)
	return int(findMinScore(grid, visited, startPos, EAST, 0.0, memo))
}

func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
}