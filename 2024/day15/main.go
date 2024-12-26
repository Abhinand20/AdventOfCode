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

// SimulateMove2 - if found a box, check if
// recursively boxes can be moved?
// 


func parseInput2(input string) ([][]rune, string) {
	g, m, _ := strings.Cut(input, "\n\n")
	rows := strings.Split(g, "\n") 
	grid := make([][]rune, len(rows))
	for i, s := range rows {
		grid[i] = make([]rune, len(s) * 2)
		k := 0
		for j := range s {
			grid[i][k] = rune(s[j])
			grid[i][k+1] = rune(s[j])
			if s[j] == '@' {
				grid[i][k+1] = '.'
			}
			if s[j] == 'O' {
				grid[i][k] = '['
				grid[i][k+1] = ']'
			}
			k += 2
		}
	}
	return grid, strings.Join(strings.Split(m, "\n"), "")
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


func horizonalMove(grid [][]rune, p Pos, m rune) Pos {
	np := nextPos(p, m)
	end := np
	for grid[end.x][end.y] == '[' || grid[end.x][end.y] == ']' {
		// Jump two times
		end = nextPos(end, m)
		end = nextPos(end, m)
	}
	if grid[end.x][end.y] == '#' {
		return p
	}
	// Shift left
	prev := grid[p.x][p.y]
	if prev == '@' {
		prev = m
	}
	curr := np
	for grid[curr.x][curr.y] != '.' {
		t := grid[curr.x][curr.y]
		grid[curr.x][curr.y] = prev
		prev = t
		curr = nextPos(curr, m)
	}
	grid[curr.x][curr.y] = prev
	grid[p.x][p.y] = '.'
	return np
}

// Only checks up/down movements
func canMove(grid [][]rune, p Pos, m rune, locs []State) (bool, []State) {
	// recursively check both positions
	if grid[p.x][p.y] == '#' {
		return false, locs
	}
	var rep rune
	if m == '^' {
		rep = grid[p.x+1][p.y]
	} else {
		rep = grid[p.x-1][p.y]
	}
	locs = append(locs, State{rep, p})
	if grid[p.x][p.y] == '.' {
		return true, locs
	}
	np := nextPos(p, m)
	curr := Pos{p.x, p.y}
	move1, l := canMove(grid, np, m, locs)
	if grid[p.x][p.y] == '[' {
		np.y++
		curr.y++
		if m == '^' {
			rep = grid[p.x+1][p.y+1]
		} else {
			rep = grid[p.x-1][p.y+1]
		}
	} else {
		np.y--
		curr.y--
		if m == '^' {
			rep = grid[p.x+1][p.y-1]
		} else {
			rep = grid[p.x-1][p.y-1]
		}
	}
	l = append(l, State{rep, curr})
	move2, l := canMove(grid, np, m, l)
	return (move1 && move2), l
}

func moveVertical(grid [][]rune, locs []State) {
	for _, s := range locs {
		grid[s.pos.x][s.pos.y] = s.dir
	}
	// Very hacky because I was too lazy to fix recursion logic
	for _, s := range locs {
		if grid[s.pos.x][s.pos.y] == '[' && grid[s.pos.x][s.pos.y+1] != ']' {
			grid[s.pos.x][s.pos.y] = '.'
		}
		if grid[s.pos.x][s.pos.y] == ']' && grid[s.pos.x][s.pos.y-1] != '[' {
			grid[s.pos.x][s.pos.y] = '.'
		}
	}
}

func simulateMove2(grid [][]rune, p Pos, m rune) Pos {
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
	// we hit either '[' or ']'
	if m == '<' || m == '>' {
		// Simple checks no DFS required
		return horizonalMove(grid, p, m)
	}
	// We need DFS
	grid[p.x][p.y] = m
	l := make([]State, 0)
	canMove, locs := canMove(grid, np, m, l)
	if canMove {
		grid[p.x][p.y] = '.'
		moveVertical(grid, locs)
		return np
	}
	return p
}

// Part2: recursive implementation to handle cases
// where boxes have other boxes aligned to adges
// ....
// [][]
//  []
//   @
func solvePart2(input string) int {
	grid, moves := parseInput2(input)
	p := Pos{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '@' {
				p.x, p.y = i, j
				break
			}
		}
	}
	RenderGrid(grid)
	for _, m := range moves {
		p = simulateMove2(grid, p, m)
		// RenderGrid(grid)
	}
	RenderGrid(grid)
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '[' {
				ans += 100 * i + j
			}
		}
	}
	return ans
}

func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
	ans2 := solvePart2(input)
	fmt.Println(ans2)
}