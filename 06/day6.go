package day6

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Position struct {
	X, Y int
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

var directions = []Position{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

func part1(input string) {
	grid, guardPos, guardDir := parse(input)

	visited := map[Position]bool{}
	rows := len(grid)
	cols := len(grid[0])

	for {
		visited[guardPos] = true

		nextPos := Position{
			X: guardPos.X + directions[guardDir].X,
			Y: guardPos.Y + directions[guardDir].Y,
		}

		if nextPos.X < 0 || nextPos.X >= rows || nextPos.Y < 0 || nextPos.Y >= cols {
			break
		}

		if grid[nextPos.X][nextPos.Y] == '#' {
			guardDir = (guardDir + 1) % 4
		} else {
			guardPos = nextPos
		}
	}

	fmt.Println(len(visited))
}

func part2(input string) {
	grid, guardPos, guardDir := parse(input)
	rows := len(grid)
	cols := len(grid[0])
	validPos := 0

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] == '.' && !(x == guardPos.X && y == guardPos.Y) {
				if simulateObstruction(copyGrid(grid), guardPos, guardDir, Position{x, y}) {
					validPos++
				}
			}
		}
	}

	fmt.Println(validPos)
}

func Run() {
	file, err := os.ReadFile("06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(string(file))
	part2(string(file))
}

func parse(input string) ([][]rune, Position, Direction) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	var guardPos Position
	var guardDir Direction

	for i, line := range lines {
		grid[i] = []rune(line)
		for j, char := range line {
			switch char {
			case '^':
				guardPos = Position{i, j}
				guardDir = Up
			case '>':
				guardPos = Position{i, j}
				guardDir = Right
			case 'v':
				guardPos = Position{i, j}
				guardDir = Down
			case '<':
				guardPos = Position{i, j}
				guardDir = Left
			}
		}
	}

	return grid, guardPos, guardDir
}
func simulateObstruction(grid [][]rune, guardPos Position, guardDir Direction, obstruction Position) bool {
	visitedStates := map[struct {
		Position
		Direction
	}]bool{}
	rows := len(grid)
	cols := len(grid[0])

	if obstruction.X >= 0 && obstruction.Y >= 0 && grid[obstruction.X][obstruction.Y] == '.' {
		grid[obstruction.X][obstruction.Y] = '#'
	}

	for {
		state := struct {
			Position
			Direction
		}{guardPos, guardDir}

		if visitedStates[state] {
			return true
		}
		visitedStates[state] = true

		nextPos := Position{
			X: guardPos.X + directions[guardDir].X,
			Y: guardPos.Y + directions[guardDir].Y,
		}
		if nextPos.X < 0 || nextPos.X >= rows || nextPos.Y < 0 || nextPos.Y >= cols {
			break
		}

		if grid[nextPos.X][nextPos.Y] == '#' {
			guardDir = (guardDir + 1) % 4
		} else {
			guardPos = nextPos
		}
	}

	if obstruction.X >= 0 && obstruction.Y >= 0 {
		grid[obstruction.X][obstruction.Y] = '.'
	}

	return false
}

func copyGrid(grid [][]rune) [][]rune {
	copied := make([][]rune, len(grid))

	for i := range grid {
		copied[i] = append([]rune{}, grid[i]...)
	}

	return copied
}
