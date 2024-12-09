package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	x, y int
}


// Find (dx,dy) between two pairs
// Calc. the left and right most points and compare with bounds

func parseInput(input string) (map[rune][]Point, int, int) {
	candidates := make(map[rune][]Point)
	tokens := strings.Split(input, "\n")
	n := len(tokens)
	m := 0
	for i, t := range tokens {
		m = len(t)
		for j, c := range t {
			if c != '.' {
				candidates[c] = append(candidates[c], Point{i,j})
			}
		}
	}
	return candidates, n, m
}


func getDistance(p1, p2 Point) Point {
	return Point{p2.x - p1.x, p2.y - p1.y}
}

func isOutside(p Point, n, m int) bool {
	return p.x < 0 || p.x >= n || p.y < 0 || p.y >= m
}


func updateAntidoes(p []Point, a map[Point]struct{}, n, m int) {
	for i := 0; i < len(p) - 1; i++ {
		for j := i + 1; j < len(p); j++ {
			pd := getDistance(p[i], p[j])
			p1 := Point{p[i].x - pd.x, p[i].y - pd.y}
			p2 := Point{p[j].x + pd.x, p[j].y + pd.y}
			if !isOutside(p1, n, m) {
				a[p1] = struct{}{}
			}
			if !isOutside(p2, n, m) {
				a[p2] = struct{}{}
			}
		}
	}
}

func updateAntidoes2(p []Point, a map[Point]struct{}, n, m int) {
	for i := 0; i < len(p) - 1; i++ {
		for j := i + 1; j < len(p); j++ {
			pd := getDistance(p[i], p[j])
			// How much can we go to the side of p1
			newP := p[i]
			for !isOutside(newP,n, m) {
				a[newP] = struct{}{}
				newP.x -= pd.x
				newP.y -= pd.y
			}
			newP = p[j]
			for !isOutside(newP,n, m) {
				a[newP] = struct{}{}
				newP.x += pd.x
				newP.y += pd.y
			}
		}
	}
}

func solveParts(input string, part2 bool) int {
	candidates, n, m:= parseInput(input)
	// Go through each pair and process
	antidotes := make(map[Point]struct{})
	for _, points := range candidates {
		updateAntidoes(points, antidotes, n, m)
		if part2 {
			updateAntidoes2(points, antidotes, n, m)
		}
	}
	return len(antidotes)
}

// Part 2 - just add max num of points
// possible on both sides

func main() {
	ans1 := solveParts(input, false)
	fmt.Println(ans1)
	ans2 := solveParts(input, true)
	fmt.Println(ans2)
}