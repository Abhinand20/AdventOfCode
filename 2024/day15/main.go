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

type State struct {
	dir rune
	pos Pos
}

var dirs = map[rune]Pos{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

func parseInput(input string) ([][]rune, string) {
	g, m, _ := strings.Cut(input, "\n\n")
	rows := strings.Split(g, "\n") 
	grid := make([][]rune, len(rows))
	for i, s := range rows {
		grid[i] = make([]rune, len(s))
		for j := range s {
			grid[i][j] = rune(s[j])
		}
	}
	return grid, strings.Join(strings.Split(m, "\n"), "")
}

// Render the position of robots in a grid
func RenderGrid(grid [][]rune) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println()
	}
}

func nextPos(p Pos, m rune) Pos {
	return Pos{p.x + dirs[m].x, p.y + dirs[m].y}
}

func simulateMove(grid [][]rune, p Pos, m rune) Pos {
	np := nextPos(p, m)
	if grid[np.x][np.y] == '#' {
		return p
	}
	if grid[np.x][np.y] == '.' {
		grid[p.x][p.y] = '.'
		grid[np.x][np.y] = m
		return np
	}
	// Handle moving boxes
	end := np
	for grid[end.x][end.y] == 'O' {
		end = nextPos(end, m)
	}
	if grid[end.x][end.y] == '#' {
		return p
	}
	grid[end.x][end.y] = 'O'
	grid[p.x][p.y] = '.'
	grid[np.x][np.y] = m
	return np
}


func solvePart1(input string) int {
	grid, moves := parseInput(input)
	p := Pos{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '@' {
				p.x, p.y = i, j
				break
			}
		}
	}
	for _, m := range moves {
		p = simulateMove(grid, p, m)
	}
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'O' {
				ans += 100 * i + j
			}
		}
	}
	return ans
}

func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
}