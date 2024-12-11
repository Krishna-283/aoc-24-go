package day11

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(input []string) {
	for i := 0; i < 25; i++ {
		next := make([]string, 0, len(input)*2)
		for _, stone := range input {
			next = append(next, step(stone)...)
		}
		input = next
	}
	fmt.Println(len(input))
}

func part2(input []string) {
	counts := make(map[string]int)
	for _, stone := range input {
		counts[stone]++
	}

	for i := 0; i < 75; i++ {
		nextCounts := make(map[string]int)
		for stone, count := range counts {
			for _, successor := range step(stone) {
				nextCounts[successor] += count
			}
		}
		counts = nextCounts
	}

	total := 0
	for _, count := range counts {
		total += count
	}
	fmt.Println(total)
}

func Run() {
	file, err := os.ReadFile("11/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := strings.TrimSpace(string(file))
	input := strings.Fields(s)

	part1(input)
	part2(input)
}

func step(data string) []string {
	if data == "0" {
		return []string{"1"}
	}

	if len(data)%2 == 0 {
		halflen := len(data) / 2
		a, b := truncateZeros(data[0:halflen]), truncateZeros(data[halflen:])
		return []string{a, b}
	}

	n, _ := strconv.Atoi(data)
	return []string{fmt.Sprint(n * 2024)}
}

func truncateZeros(data string) string {
	if len(data) > 1 {
		if data[0] == '0' {
			return truncateZeros(data[1:])
		}
		return data
	}
	return data
}
