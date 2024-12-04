package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var diagonalDirections = [4][2]int{
	{-1, -1}, // North-West
	{-1, 1},  // North-East
	{1, 1},   // South-East
	{1, -1},  // South-West
}


func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func getDiagonal(grid []string, row, col int, direction [2]int, length int) string {
	diagonal := ""
	for step := 0; step < length; step++ {
		newRow := row + step*direction[0]
		newCol := col + step*direction[1]

		// Check bounds
		if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[0]) {
			diagonal += string(grid[newRow][newCol])
		} else {
			break // Stop if out of bounds
		}
	}
	return diagonal
}

func countWord(grid []string, i, j int) int {
	ans := 0
	ref := "XMAS"
	revRef := "SAMX"
	m := len(grid)
	n := len(grid[0])
	// Check right and left
	if j + 4 <= n && grid[i][j:j+4] == ref {
		ans += 1
	}
	if j - 3 >= 0 && grid[i][j-3:j+1] == revRef {
		ans += 1
	}
	// Check up and down
	if i - 3 >= 0 {
		c := ""
		for k := 0; k < 4; k++ {
			c += string(grid[i-k][j])
		}
	}
	if i - 3 >= 0 {
		c := ""
		for k := 0; k < 4; k++ {
			c += string(grid[i-k][j])
		}
		if c == ref {
			ans++
		}
	}
	if i + 3 < m {
		c := ""
		for k := 0; k < 4; k++ {
			c += string(grid[i+k][j])
		}
		if c == ref {
			ans++
		}
	}
	for _, direction := range diagonalDirections {
		diagonal := getDiagonal(grid, i, j, direction, 4)
		if len(diagonal) == 4 && diagonal == ref {
			ans++
		}
	}
	return ans
}

func solvePart1(input string) int {
	grid := parseInput(input)
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'X' {
				ans += countWord(grid, i, j)
			}
		}
	}
	return ans
}
func solvePart2(input string) int {
	grid := parseInput(input)
	ans := 0
	for i := 1; i < len(grid) - 1; i++ {
		for j := 1; j < len(grid[0]) - 1; j++ {
			if grid[i][j] == 'A' {
				if (grid[i-1][j-1] == 'S' && grid[i+1][j+1] == 'M') || (grid[i-1][j-1] == 'M' && grid[i+1][j+1] == 'S') {
					if (grid[i+1][j-1] == 'S' && grid[i-1][j+1] == 'M') || (grid[i+1][j-1] == 'M' && grid[i-1][j+1] == 'S') {
						ans += 1
					}
				}
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


// This is wrong because of a misunderstanding about what the question was asking.
/*
var directions = [8][2]int{
	{-1, 0},  // North
	{-1, 1},  // North-East
	{0, 1},   // East
	{1, 1},   // South-East
	{1, 0},   // South
	{1, -1},  // South-West
	{0, -1},  // West
	{-1, -1}, // North-West
}

var cycle = map[byte]byte{
	'X': 'M',
	'M': 'A',
	'A': 'S',
	'S': '0',
}
func countWord(grid []string, curr byte, i, j int) int {
	m := len(grid)
	n := len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return 0
	}
	fmt.Println(string(curr), i, j)
	// Check if we should continue
	if grid[i][j] == curr {
		next, _ := cycle[curr]
		// If we found the last 'S' return 1
		if next == '0' {
			fmt.Println("Found it")
			return 1
		}
		// Else continue searching in all directions
		ans := 0
		for _,d := range directions {
			ans += countWord(grid, next, i+d[0], j+d[1])
		}
		return ans
	}
	return 0
}
*/