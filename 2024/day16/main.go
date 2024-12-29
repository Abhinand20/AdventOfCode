package main

import (
	"container/heap"
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
	inf = math.MaxFloat64
)


/* Priority Queue */
type Node struct {
	p Pos
	score float64
	dir int
}

type PriorityQueue []Node

// Implement heap.Interface
func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(val any) {
	*pq = append(*pq, val.(Node))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	ret := old[n-1]
	*pq = old[0:n-1]
	return ret 
}


// E, S, W, N
var dirs = [4][2]int{{0,1}, {1,0}, {0,-1}, {-1,0}}
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

func findMinScore(grid [][]rune, visited [][]bool, s Pos, d int, score float64) float64 {
	if grid[s.x][s.y] == 'E' {
		return score
	}
	if grid[s.x][s.y] == '#' || visited[s.x][s.y] || score >= answer {
		return math.MaxFloat64
	}
	visited[s.x][s.y] = true
	// Check other options
	np := nextPos(s, d)
	npc, dc := turnClockwise(s, d)
	npa, da := turnAntiClockwise(s, d)
	// Did we reach the dest?
	ans := findMinScore(grid, visited, np, d, 1 + score)
	ans = math.Min(ans, findMinScore(grid, visited, npc, dc, 1001 + score))
	ans = math.Min(ans, findMinScore(grid, visited, npa, da, 1001 + score))
	visited[s.x][s.y] = false
	answer = math.Min(answer, ans)
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
	return int(findMinScore(grid, visited, startPos, EAST, 0.0))
}

// Djikstra's - cost is based on direction?
func solvePart1Djisktra(input string) float64 {
	grid := parseInput(input)
	startPos := Pos{}
	visited := make(map[cacheKey]bool)
	for i := range grid {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'S' {
				startPos.x, startPos.y = i, j
			}
		}
	}
	pq := &PriorityQueue{Node{p: startPos, dir: EAST, score: 0}}
	var ans float64
	for pq.Len() > 0 {
		cn := heap.Pop(pq).(Node)
		// We reached the final node
		if grid[cn.p.x][cn.p.y] == 'E' {
			ans = cn.score
			break
		}
		if grid[cn.p.x][cn.p.y] == '#' {
			continue
		}
		ck := cacheKey{cn.dir, cn.p}
		if _, ok := visited[ck]; ok {
			continue
		}
		visited[ck] = true
		np := nextPos(cn.p, cn.dir)
		npc, dc := turnClockwise(cn.p, cn.dir)
		npa, da := turnAntiClockwise(cn.p, cn.dir)

		heap.Push(pq, Node{np, cn.score + 1, cn.dir})
		heap.Push(pq, Node{npc, cn.score + 1001, dc})
		heap.Push(pq, Node{npa, cn.score + 1001, da})
	}
	return ans
}

func main() {
	ans1 := solvePart1Djisktra(input)
	fmt.Println(ans1)
}