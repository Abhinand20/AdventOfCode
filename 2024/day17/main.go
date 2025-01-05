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

func (cpu *CPU) printASM() {
	strCode := ""
	for i := 0; i < len(cpu.code) - 1; i += 2{
		opcode := cpu.code[i]
		operand := cpu.code[i+1]
		switch opcode {
		case 0: {
			strCode = fmt.Sprintf("%s\n[%d] %s %v", strCode, i, "ADV", cpu.comboDebug(operand))
		}
		case 1: {
			strCode = fmt.Sprintf("%s\n[%d] %s %s %v", strCode, i, "BXL", "B", operand)
		}
		case 2: {
			strCode = fmt.Sprintf("%s\n[%d] %s %s %v", strCode, i, "BST", "B", cpu.comboDebug(operand))
		}
		case 3: {
			strCode = fmt.Sprintf("%s\n[%d] %s %v", strCode, i, "JNZ", cpu.comboDebug(operand))
		}
		case 4: {
			strCode = fmt.Sprintf("%s\n[%d] %s", strCode, i, "BXC")
		}
		case 5: {
			strCode = fmt.Sprintf("%s\n[%d] %s %v", strCode, i, "OUT", cpu.comboDebug(operand)) 
		}
		case 6: {
			strCode = fmt.Sprintf("%s\n[%d] %s %v", strCode, i, "BDV", cpu.comboDebug(operand))
		}
		case 7: {
			strCode = fmt.Sprintf("%s\n[%d] %s %v", strCode, i, "CDV", cpu.comboDebug(operand))
		}
		}
	}
	fmt.Println(strCode)
}

func (cpu *CPU) comboDebug(o int) string {
	switch {
		case o < 4: {
			return strconv.Itoa(o)
		}
		case o == 4: {
			return "A"
		}
		case o == 5: {
			return "B"
		}
		case o == 6: {
			return "C"
		}
		default: {
			return ""
		}
	}
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
		// fmt.Println(opcode, operand)
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
		// fmt.Println(cpu.reg)
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
	cpu.printASM()
	return cpu.runCode()
}

func incrementOctalDigitAtIndex(octal int, index int) int {
	numDigits := 16
	rightIndex := numDigits - index - 1

	power := int(1)
	for i := 0; i < rightIndex; i++ {
		power *= 8
	}

	digit := (octal / power) % 8

	if digit < 7 {
		octal += power
	} else {
		panic("Overflow octal...")
	}

	return octal
}


func solvePart2(input string) int {
	// 100000000000000
	// 555555555555555
	// Start with octal 10 ** 16
	// Keep fixing digits starting from left side (equal to last values)
	_, code := parseInput(input)
	var res int = 0o1000000000000000
	for i := 0; i < len(code); i++ {
		reg := Reg{res, 0, 0}
		cpu := NewCPU(reg, code)
		currRes := cpu.runCode()
		want := code[len(code) - 1 - i]
		got, _ := strconv.Atoi(strings.Split(currRes, ",")[len(code) - 1 - i])
		for got != want {
			res = incrementOctalDigitAtIndex(res, i)
			// calc new val
			reg = Reg{res, 0, 0}
			cpu := NewCPU(reg, code)
			out := cpu.runCode()
			fmt.Println(out)
			got, _ = strconv.Atoi(strings.Split(out, ",")[len(code) - 1 - i])
		}

	}
	return res
}


func main() {
	ans1 := solvePart1(input)
	fmt.Println("A1: ", ans1)
	ans2 := solvePart2(input)
	fmt.Println(ans2)
}