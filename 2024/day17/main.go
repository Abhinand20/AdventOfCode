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

type Reg struct {
	A, B, C int
}

type CPU struct {
	reg Reg
	PC int
	code []int
}

func (cpu *CPU) combo(o int) int {
	switch {
		case o < 4: {
			return o
		}
		case o == 4: {
			return cpu.reg.A
		}
		case o == 5: {
			return cpu.reg.B
		}
		case o == 6: {
			return cpu.reg.C
		}
		default: {
			return 0
		}
	}
}

func (cpu *CPU) runCode() string {
	ans := ""
	for cpu.PC < len(cpu.code) {
		opcode := cpu.code[cpu.PC]
		operand := cpu.code[cpu.PC + 1]
		fmt.Println(opcode, operand)
		cpu.PC += 2
		switch opcode {
		case 0: {
			res := int(float64(cpu.reg.A) / math.Exp2(float64(cpu.combo(operand))))
			cpu.reg.A = res
		}
		case 1: {
			cpu.reg.B = cpu.reg.B ^ operand
		}
		case 2: {
			cpu.reg.B = cpu.combo(operand) % 8
		}
		case 3: {
			if cpu.reg.A != 0 {
				cpu.PC = operand
			}
		}
		case 4: {
			cpu.reg.B = cpu.reg.C ^ cpu.reg.B
		}
		case 5: {
			ans += strconv.Itoa(cpu.combo(operand) % 8) + ","
		}
		case 6: {
			res := int(float64(cpu.reg.A) / math.Exp2(float64(cpu.combo(operand))))
			cpu.reg.B = res
		}
		case 7: {
			res := int(float64(cpu.reg.A) / math.Exp2(float64(cpu.combo(operand))))
			cpu.reg.C = res
		}
		}
		fmt.Println(cpu.reg)
	}
	return strings.TrimRight(ans, ",")
}

func NewCPU(reg Reg, program []int) *CPU {
	return &CPU{
		reg: reg,
		PC: 0,
		code: program,
	}
}

func parseInput(input string) (Reg, []int) {
	r, c, _:= strings.Cut(input, "\n\n")
	regex := regexp.MustCompile("\\d+")
	vals := make([]int, 3)
	for i, row := range strings.Split(r, "\n") {
		v, _ := strconv.Atoi(regex.FindString(row))
		vals[i] = v
	}
	reg := Reg{vals[0], vals[1], vals[2]}
	code := make([]int, 0)
	for _, ch := range strings.Split(c, ",") {
		n, _ := strconv.Atoi(ch)
		code = append(code, n)
	}
	return reg, code
}

func solvePart1(input string) string {
	r, c := parseInput(input)
	cpu := NewCPU(r, c)
	return cpu.runCode()
}

func solvePart2(input string) string {
	// 100000000000000
	return ""
}


func main() {
	ans1 := solvePart1(input)
	fmt.Println(ans1)
	// ans2 := solvePart2(input)
}