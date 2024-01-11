package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const inputFile string = "../inputs/day8_1.txt"
// const inputFile string = "../tests/day8_1.txt"
var dirMap = map[string]int{ "L" : 0, "R" : 1}

func ProcessInput() (string, map[string][2]string, error) {
	file, err := os.Open(inputFile)
	graph := make(map[string][2]string)
	if err != nil {
		fmt.Println(err)
		return "", graph, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Scan()
	directions := fileScanner.Text()

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			var tokens []string = strings.Split(line, " = ")
			src := tokens[0]
			nodes := strings.Split(tokens[1], ", ")
			neighbors := [2]string{nodes[0][1:], nodes[1][:3]}
			graph[src] = neighbors
		}	
	}
	return directions, graph, nil
}

func solve(dirs string, graph map[string][2]string) int {
	valid := true
	ans := 0
	i := 0
	node := "AAA"
	dest := "ZZZ"
	for valid {
		if i >= len(dirs) {
			i = i % len(dirs)
		}
		node = graph[node][dirMap[string(dirs[i])]]
		ans++
		if node == dest {
			valid = false
		}
		i++
	}
	return ans
}

func GetMinSteps(src, dirs string, graph map[string][2]string) int {
	i := 0
	steps := 0
	for {
		if i >= len(dirs) {
			i %= len(dirs)
		}
		steps++
		src = graph[src][dirMap[string(dirs[i])]]
		if src[len(src) - 1] == 'Z' {
			break
		}
		i++
	}
	return steps
}

func GetGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GetGCD(b, a % b)
}


func GetLCM(arr []int) int {
	ans := arr[0]

	for i := 1; i < len(arr); i++ {
		ans = (ans * arr[i]) / GetGCD(ans, arr[i])
	}

	return ans
}

func solve2(dirs string, graph map[string][2]string) int {
	start := make([]string, 0)
	for k := range graph {
		if k[len(k) - 1] == 'A' {
			start = append(start, k)
		}
	}
	minSteps := make([]int, len(start))
	for idx, src := range start {
		minSteps[idx] = GetMinSteps(src, dirs, graph)
	}
	fmt.Println(minSteps)
	return GetLCM(minSteps)
}


func main() {
	dirs, graph, err := ProcessInput()
	if err != nil {
		return
	}
	ans := solve2(dirs, graph)
	// ans := solve(dirs, graph)
	fmt.Println(ans)
}