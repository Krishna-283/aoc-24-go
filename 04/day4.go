package day4

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func part1(grid [][]string) {
	dirs := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}
	total := 0

	for i, row := range grid {
		for j, char := range row {
			if char != "X" {
				continue
			}
			for _, dir := range dirs {
				di, dj := dir[0], dir[1]
				hit := true
				for k := 0; k < 4; k++ {
					if get(grid, i+k*di, j+k*dj) != string("XMAS"[k]) {
						hit = false
						break
					}
				}
				if hit {
					total++
				}
			}
		}
	}
	fmt.Println(total)
}

func part2(grid [][]string) {
	dirs := [][2]int{
		{1, 1},   // down-right
		{-1, -1}, // up-left
		{1, -1},  // down-left
		{-1, 1},  // up-right
	}
	total := 0

	for i, row := range grid {
		for j, char := range row {
			if char != "A" {
				continue
			}

			word := []string{}
			for _, dir := range dirs {
				dx, dy := dir[0], dir[1]
				word = append(word, get(grid, i+dx, j+dy))
			}

			if word[0] == word[1] {
				continue
			}

			sort.Strings(word)
			if word[0]+word[1]+word[2]+word[3] == "MMSS" {
				total++
			}
		}
	}
	fmt.Println(total)
}

func Run() {
	file, err := os.ReadFile("04/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(file), "\n")
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}

	part1(grid)
	part2(grid)
}

func get(grid [][]string, i, j int) string {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) {
		return grid[i][j]
	}
	return "-"
}
