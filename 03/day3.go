package day3

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(input string) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	sum := 0

	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}

	fmt.Println("Part 1 result", sum)
}

func part2(input string) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	enabled := true
	sum := 0

	for _, match := range matches {
		if len(match) == 0 {
			continue
		}

		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		}

		if enabled && len(match) > 2 && match[0][:3] == "mul" {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			sum += num1 * num2
		}
	}

	fmt.Println("Part 2 result", sum)
}

func Run() {
	data, err := os.ReadFile("03/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	input := strings.TrimSpace(string(data))
	part1(input)
	part2(input)
}

