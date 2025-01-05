package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const (
	N = 71
	LIMIT = 1024
)


type Pos struct {
	x, y int
}

var dirs = [5]int{0, 1, 0, -1, 0}

func RenderGrid(grid [][]rune) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println()
	}
}

func parseInput(input string, limit int) [][]rune {
	rows := strings.Split(input, "\n")
	grid := make([][]rune, N)
	for j := 0; j < N; j++ {
		grid[j] = make([]rune, N)
		for k := 0; k < N; k++ {
			grid[j][k] = '.'
		}
	}
	for i, row := range rows {
		if i == limit {
			break
		}
		sy, sx, _ := strings.Cut(row, ",")
		x, _ := strconv.Atoi(sx)
		y, _ := strconv.Atoi(sy)
		grid[x][y] = '#'
	}
	return grid
}


func isValid(grid [][]rune, p Pos) bool {
	return p.x >= 0 && p.x < N && p.y >= 0 && p.y < N && grid[p.x][p.y] != '#'
}

func BFS(grid [][]rune) int {
	start := Pos{0, 0}
	end := Pos{N-1, N-1}
	queue := make([][]Pos, 0)
	visited := make(map[Pos]bool)
	queue = append(queue, []Pos{start})
	visited[start] = true
	ans := 0
	stop := false
	for len(queue) > 0 && !stop {
		// Pop from queue
		curr := queue[0]
		queue = queue[1:]
		ngbs := make([]Pos, 0)
		for _, c := range curr {
			if c == end {
				stop = true
				return ans
			}
			for i := 0; i < 4; i++ {
				np := Pos{c.x + dirs[i], c.y + dirs[i + 1]}
				if isValid(grid, np) && !visited[np] {
					ngbs = append(ngbs, np)
					visited[np] = true
				}
			}
		}
		if len(ngbs) > 0 {
			queue = append(queue, ngbs)
		}
		ans++
	}
	return -1
}


func solvePart1(input string) int {
	grid := parseInput(input, LIMIT)
	return BFS(grid)
}

func solvePart2(input string) string {
	indices := strings.Split(input, "\n")
	for i := LIMIT; i < len(indices); i++ {
		grid := parseInput(input, i)
		res := BFS(grid)
		if res == -1 {
			return indices[i - 1]
		}
	}
	return "UNK"
}

func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
	ans2 := solvePart2(input)
	fmt.Println(ans2)
}