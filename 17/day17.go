package day17

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Cpu struct represents the state of the CPU with registers A, B, C,
// an instruction pointer (IP), the program itself, and output.
type Cpu struct {
	A, B, C, IP int
	Program     []int
	Output      []int
}

func Run() {
	file, err := os.ReadFile("17/input.txt")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	lines := strings.Split(string(file), "\n")

	cpu := Cpu{}
	cpu.A, _ = strconv.Atoi(strings.Split(lines[0], ": ")[1])
	cpu.B, _ = strconv.Atoi(strings.Split(lines[1], ": ")[1])
	cpu.C, _ = strconv.Atoi(strings.Split(lines[2], ": ")[1])

	programStr := strings.Split(strings.Split(lines[4], ": ")[1], ",")
	for _, s := range programStr {
		val, _ := strconv.Atoi(strings.TrimSpace(s))
		cpu.Program = append(cpu.Program, val)
	}

	for cpu.IP < len(cpu.Program) {
		cpu.executeInstruction()
	}

	fmt.Println("Part 1:", listToString(cpu.Output))
	fmt.Println("Part 2:", cpu.findMatchingA(cpu.Program, 0))
}

func (cpu *Cpu) executeInstruction() {
	opcode := cpu.Program[cpu.IP]
	operand := cpu.Program[cpu.IP+1]

	switch opcode {
	case 0:
		cpu.A = cpu.A / (1 << cpu.getOperandValue(operand))
	case 1:
		cpu.B ^= operand
	case 2:
		cpu.B = cpu.getOperandValue(operand) % 8
	case 3:
		if cpu.A != 0 {
			cpu.IP = operand - 2
		}
	case 4:
		cpu.B ^= cpu.C
	case 5:
		cpu.Output = append(cpu.Output, cpu.getOperandValue(operand)%8)
	case 6:
		cpu.B = cpu.A / (1 << cpu.getOperandValue(operand))
	case 7:
		cpu.C = cpu.A / (1 << cpu.getOperandValue(operand))
	default:
		panic("Invalid opcode")
	}

	cpu.IP += 2
}

func (cpu *Cpu) getOperandValue(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return cpu.A
	case 5:
		return cpu.B
	case 6:
		return cpu.C
	default:
		panic("Invalid operand")
	}
}

func listToString(list []int) string {
	var sb strings.Builder
	for i, val := range list {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(fmt.Sprintf("%d", val))
	}
	return sb.String()
}

func (cpu *Cpu) findMatchingA(targetOutput []int, targetA int) int {
	if len(targetOutput) == 0 {
		return targetA
	}
	results := []int{}
	for i := targetA * 8; i < (targetA+1)*8; i++ {
		b, newA := cpu.simulateProgram(i)
		if newA != targetA {
			panic("newA != targetA")
		}
		if b == targetOutput[len(targetOutput)-1] {
			result := cpu.findMatchingA(targetOutput[:len(targetOutput)-1], i)
			if result != -1 {
				results = append(results, result)
			}
		}
	}
	if len(results) == 0 {
		return -1
	}
	return slices.Min(results)
}

func (cpu *Cpu) simulateProgram(a int) (int, int) {
	cpu.IP = 0
	cpu.A = a
	for cpu.Program[cpu.IP] != 5 {
		cpu.executeInstruction()
	}
	output := cpu.getOperandValue(cpu.Program[cpu.IP+1]) % 8
	cpu.IP += 2
	for cpu.Program[cpu.IP] != 3 {
		cpu.executeInstruction()
	}
	return output, cpu.A
}
