package main

import (
	"container/list"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

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

func parseInput(input string) [][]rune {
	rows := strings.Split(input, "\n")
	grid := make([][]rune, len(rows))

	for i, r := range rows {
		grid[i] = make([]rune, len(r))
		for j := range r {
			grid[i][j] = rune(r[j])
		}
	}
	return grid
}


func isValid(grid [][]rune, p Pos) bool {
	n, m := len(grid), len(grid[0])
	return p.x >= 0 && p.x < n && p.y >= 0 && p.y < m && grid[p.x][p.y] != '#'
}

func getRemovals(grid [][]rune) map[Pos]bool {
	r := make(map[Pos]bool)
	n, m := len(grid), len(grid[0])
	for i := 1; i < n - 1; i++ {
		for j := 1; j < m - 1; j++ {
			if grid[i][j] != '#' {
				continue
			}
			// can we remove this wall?
			numPaths := 0
			for k := 0; k < 4; k++ {
				x, y := i + dirs[k], j + dirs[k+1]
				if grid[x][y] != '#' {
					numPaths++
				}
			}
			if numPaths > 1 {
				r[Pos{i,j}] = true
			}
		}
	}
	return r
}

func BFS(grid [][]rune) int {
	start := Pos{67, 76}
	end := Pos{43, 72}
	queue := list.New()
	queue.PushBack([]int{start.x, start.y, 0})
	visited := make(map[Pos]bool)
	visited[start] = true
	for queue.Len() > 0 {
		top := queue.Remove(queue.Front()).([]int)
		curr, depth := Pos{top[0], top[1]}, top[2]
		if curr == end {
			return depth
		}
		for i := 0; i < 4; i++ {
			np := Pos{curr.x + dirs[i], curr.y + dirs[i+1]}
			if isValid(grid, np) && !visited[np] {
				queue.PushBack([]int{np.x, np.y, depth + 1})
				visited[np] = true
			}
		}
	}
	return -1
}

func solvePart1(input string) int {
	grid := parseInput(input)
	removals := getRemovals(grid)
	fmt.Println(len(removals))
	savings := make(map[int]int)
	maxCost := BFS(grid)
	fmt.Println("Max: ", maxCost)
	for r := range removals {
		grid[r.x][r.y] = '.'
		currCost := BFS(grid)
		if currCost > -1 && currCost < maxCost {
			savings[maxCost - currCost]++
		}
		grid[r.x][r.y] = '#'
	}
	ans := 0
	for k,v := range savings {
		if k >= 100 {
			ans += v
		}
	}
	return ans
}

func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
}