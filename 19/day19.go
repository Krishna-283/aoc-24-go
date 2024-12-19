package day19

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func part1(stripes, towels []string) {
	fmt.Println(len(findValidDesigns(stripes, towels)))
}

func part2(towelPatterns, designs []string) {
	validDesigns := findValidDesigns(towelPatterns, designs)
	counts := make(map[string]int)

	totalCombinations := 0
	for _, design := range validDesigns {
		totalCombinations += countWays(design, towelPatterns, counts)
	}

	fmt.Println(totalCombinations)
}

func Run() {
	file, err := os.ReadFile("19/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	parts := strings.Split(strings.TrimSpace(string(file)), "\n\n")
	towelPatterns := strings.Split(strings.TrimSpace(parts[0]), ", ")
	designs := strings.Split(parts[1], "\n")

	part1(towelPatterns, designs)
	part2(towelPatterns, designs)
}

func findValidDesigns(towelPatterns, designs []string) []string {
	var validDesigns []string

	for _, design := range designs {
		queue := []string{design}

		for len(queue) > 0 {
			currentDesign := queue[0]
			queue = queue[1:]

			if currentDesign == "" {
				validDesigns = append(validDesigns, design)
				break
			}

			valid := false
			for _, towelPattern := range towelPatterns {
				n := strings.TrimPrefix(currentDesign, towelPattern)
				if len(n) < len(currentDesign) {
					valid = true
					queue = append([]string{n}, queue...)
				}
			}
			if !valid {
				continue
			}
		}
	}

	return validDesigns
}

func countWays(design string, towelPatterns []string, counts map[string]int) int {
	if count, exists := counts[design]; exists {
		return count
	}

	if design == "" {
		return 1
	}

	totalWays := 0
	for _, towelPattern := range towelPatterns {
		if strings.HasPrefix(design, towelPattern) {
			totalWays += countWays(strings.TrimPrefix(design, towelPattern), towelPatterns, counts)
		}
	}

	counts[design] = totalWays
	return totalWays
}
