package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	x int
	y int
}

func (p position) equals(other position) bool {
	return p.x == other.x && p.y == other.y
}

var posLookup = map[string][2]position{
	"F": {position{0,1}, position{1,0}},
	"|": {position{0,1}, position{0,-1}},
	"-": {position{1,0}, position{-1,0}},
	"L": {position{0,-1}, position{1,0}},
	"J": {position{0,-1}, position{-1,0}},
	"7": {position{0,1}, position{-1,0}},
	"S": {position{0,0}, position{0,0}},
}

func ProcessInputD10() ([]string, error) {
	const inputFile string = "../tests/day10_1.txt"
	file, err := os.Open(inputFile)
	inputs := make([]string, 0)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			inputs = append(inputs, line)
		}	
	}
	return inputs, nil
}

func IsValid(pos position, arr []string) bool {
	m, n := len(arr), len(arr[0])
	bounded := pos.x >= 0 && pos.x < m && pos.y >= 0 && pos.y < n
	if bounded {
		_, exist := posLookup[string(arr[pos.x][pos.y])]
		return exist
	}
	return false
}

func FindLoopSize(prev position, pos position, arr []string) (int, bool) {
	len := 1
	fmt.Println("Finding loop at: ", pos)
	for {
		for i := 0; i < 2; i++ {
			val := string(arr[pos.x][pos.y])
			dirs := posLookup[val]
			nextPos := position{pos.x + dirs[i].y, pos.y + dirs[i].x} // TODO: check
			if prev.equals(nextPos) {
				continue
			}
			// fmt.Println(val)
			if !IsValid(nextPos, arr) {
				continue
			}
			// DEBUG
			curr := arr[pos.x]
			arr[pos.x] = curr[:pos.y] + "#" + curr[pos.y+1:]

			if string(arr[nextPos.x][nextPos.y]) == "S" {
				return len + 1, true
			}
			len++
			prev = pos
			pos = nextPos
		}
	}
}

func StartTraversal(pos position, arr []string) int {
	fmt.Println("Starting at: ", pos)
	allDirs := []int{1, 0, -1, 0, 1}

	for i := 0; i < 4; i++ {
		nextPos := position{pos.x + allDirs[i], pos.y + allDirs[i+1]}
		if IsValid(nextPos, arr) {
			l, valid := FindLoopSize(pos, nextPos, arr)
			if valid {
				for i := 0; i < len(arr); i++ {
					fmt.Println(arr[i])
				}
				return l / 2 // odd size?
			}
		}
	}
	return -1
}

func SolveDay10Part1(arr []string) int {
	ans := 0
	// 1. Get starting positions
	for i, s := range arr {
		for j := 0; j < len(s); j++ {
			val := string(s[j])
			if val == "S" {
				ans = StartTraversal(position{i,j}, arr)
			}
		}
	}
	// 2. Traverse at once and find the point both meet
	fmt.Println(ans)
	return ans
}
func SolveDay10() {
	inputs, _ := ProcessInputD10()
	// fmt.Println(inputs)
	SolveDay10Part1(inputs)
	
}