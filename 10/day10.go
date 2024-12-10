package day10

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var dirs = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func part1(lines []string) {
	totalTrailScore := 0
	for i, line := range lines {
		for j, char := range line {
			if char == '0' {
				totalTrailScore += score(lines, i, j)
			}
		}
	}
	fmt.Println(totalTrailScore)
}

func part2(lines []string) {
	totalTrailRating := 0
	for i, line := range lines {
		for j, char := range line {
			if char == '0' {
				totalTrailRating += rating(lines, i, j)
			}
		}
	}
	fmt.Println(totalTrailRating)
}

func Run() {
	file, err := os.ReadFile("10/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	rawInput := strings.Split(string(file), "\n")
	lines := []string{}
	for _, line := range rawInput {
		if len(strings.TrimSpace(line)) > 0 {
			lines = append(lines, line)
		}
	}

	part1(lines)
	part2(lines)
}

func score(lines []string, i, j int) int {
	count := 0
	rows, cols := len(lines), len(lines[0])

	visited := make([][]bool, rows)
	for k := range visited {
		visited[k] = make([]bool, cols)
	}

	dfs(lines, i, j, &visited)
	for _, row := range visited {
		for _, val := range row {
			if val {
				count++
			}
		}
	}
	return count
}

func dfs(lines []string, i int, j int, visited *[][]bool) {
	if lines[i][j] == '9' {
		(*visited)[i][j] = true
		return
	}

	nxt := byte(int(lines[i][j]-'0') + 1)
	for _, d := range dirs {
		if inRange(lines, i+d[0], j+d[1]) &&
			byte(lines[i+d[0]][j+d[1]]-'0') == nxt {
			dfs(lines, i+d[0], j+d[1], visited)
		}
	}
}

func rating(lines []string, i, j int) int {
	if lines[i][j] == '9' {
		return 1
	}

	nxt := byte(int(lines[i][j]-'0') + 1)
	val := 0

	for _, d := range dirs {
		if inRange(lines, i+d[0], j+d[1]) &&
			byte(lines[i+d[0]][j+d[1]]-'0') == nxt {
			val += rating(lines, i+d[0], j+d[1])
		}
	}
	return val
}

func inRange(lines []string, i, j int) bool {
	return (i >= 0 && i < len(lines) && j >= 0 && j < len(lines[0]))
}
