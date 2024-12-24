package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const H = 103
const W = 101
const T = 100


type Pos struct {
	x, y int
}

type Robot struct {
	p, v Pos
}

func AbsInt(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func parseRow(row string) Robot {
	pos, vel, _ := strings.Cut(row, " ")
	reg := regexp.MustCompile("-*\\d+")
	p := reg.FindAllString(pos, -1)
	v := reg.FindAllString(vel, -1)
	initPos := Pos{}
	for idx, i := range p {
		s, _ := strconv.Atoi(i)
		if idx == 0 {
			initPos.x = s
		} else {
			initPos.y = s
		}
	}
	fixedVel := Pos{}
	for idx, i := range v {
		s, _ := strconv.Atoi(i)
		if idx == 0 {
			fixedVel.x = s
		} else {
			fixedVel.y = s
		}
	}
	return Robot{initPos, fixedVel}
}


func parseInput(input string) []Robot {
	res := make([]Robot, 0)
	rows := strings.Split(input, "\n")
	for _, r := range rows {
		r := parseRow(r)
		res = append(res, r)
	}
	return res
}

func findQuad(p Pos) int {
	mid := Pos{W / 2, H / 2}
	if p.x == mid.x || p.y == mid.y {
		return -1
	}
	if p.x < mid.x {
		if p.y < mid.y {
			return 0
		}
		return 2
	}
	if p.y < mid.y {
		return 1
	}
	return 3
}

func solvePart1(input string) int {
	robots := parseInput(input)
	// Go through each and update the final positions
	finalRobots := make(map[Pos]int)
	for _, r := range robots {
		fx := ((r.v.x * T) % W + r.p.x) % W
		fy := ((r.v.y * T) % H + r.p.y) % H
		fp := Pos{(fx + W) % W, (fy + H) % H}
		finalRobots[fp]++
	}
	quads := make([]int, 4)
	for p, c := range finalRobots {
		quad := findQuad(p)
		if quad != -1 {
			quads[quad] += c
		}
	}
	ans := 1
	for _, q := range quads {
		ans *= q
	}
	return ans
}

// Render the position of robots in a grid
func RenderGrid(rp map[Pos]int) {
	grid := make([][]string, H)
	for i := 0; i < H; i++ {
		grid[i] = make([]string,W)
		for j := 0; j < W; j++ {
			grid[i][j] = "."
		}
	}
	for p := range rp {
		grid[p.y][p.x] = strconv.Itoa(rp[p])
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Printf("%s", grid[i][j])
		}
		fmt.Printf("\n")
	}
}

func Debug(robots []Robot, iter int) {
	finalRobots := make(map[Pos]int)
	for _, r := range robots {
		fx := ((r.v.x * iter) % W + r.p.x) % W
		fy := ((r.v.y * iter) % H + r.p.y) % H
		fp := Pos{(fx + W) % W, (fy + H) % H}
		finalRobots[fp]++
	}
	// Look for a horizontal line of robots (say atleast 10 together)
	maxLine := 0.0
	for i := 0; i < H; i++ {
		streak := 0
		for j := 0; j < W; j++ {
			if _, ok := finalRobots[Pos{j, i}]; ok {
				streak++
			} else {
				streak = 0
			}
			maxLine = math.Max(float64(maxLine), float64(streak))
		}
	}
	if maxLine >= 10 {
		RenderGrid(finalRobots)
		fmt.Println("Iter: ", iter)
	}
}

func main() {
	// ans1 := solvePart1(input)
	// fmt.Println(ans1)
	robots := parseInput(input)
	for it := 0; it < 800000; it++ {
		Debug(robots, it)
	}
}