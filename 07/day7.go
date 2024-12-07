package day7

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Value   int
	Numbers []int
}

func part1(input []Input) {
	sum := 0
	for _, eq := range input {
		if matchTargetPart1(eq.Numbers, eq.Value, eq.Numbers[0], 1) {
			sum += eq.Value
		}
	}
	fmt.Println("Part 1:", sum)
}

func part2(input []Input) {
	sum := 0
	for _, eq := range input {
		if matchTargetPart2(eq.Numbers, eq.Value, eq.Numbers[0], 1) {
			sum += eq.Value
		}
	}
	fmt.Println("Part 2:", sum)
}

func Run() {
	file, err := os.ReadFile("07/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	data := strings.Split(string(file), "\n")
	input := make([]Input, 0, len(data))

	for _, line := range data {
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 3 {
			log.Fatalln("Invalid line")
		}

		valueStr := strings.TrimSuffix(fields[0], ":")
		value, _ := strconv.Atoi(valueStr)
		var numbers []int

		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			numbers = append(numbers, num)
		}

		input = append(input, Input{
			Value:   value,
			Numbers: numbers,
		})
	}

	part1(input)
	part2(input)
}

func matchTargetPart1(nums []int, target, current, index int) bool {
	if index == len(nums) {
		return current == target
	}

	if matchTargetPart1(nums, target, current+nums[index], index+1) {
		return true
	}

	if matchTargetPart1(nums, target, current*nums[index], index+1) {
		return true
	}

	return false
}

func matchTargetPart2(nums []int, target, current, index int) bool {
	if index == len(nums) {
		return current == target
	}

	if matchTargetPart2(nums, target, current+nums[index], index+1) {
		return true
	}

	if matchTargetPart2(nums, target, current*nums[index], index+1) {
		return true
	}

	concat := concatInt(current, nums[index])
	if matchTargetPart2(nums, target, concat, index+1) {
		return true
	}

	return false
}

func concatInt(a, b int) int {
	digit := 0
	result := a

	for tmp := b; tmp > 0; tmp /= 10 {
		digit++
	}
	for i := 0; i < digit; i++ {
		result *= 10
	}

	result += b
	return result
}
