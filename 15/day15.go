package day15

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

var directions = map[string]Point{
	"^": {x: 0, y: -1},
	"v": {x: 0, y: 1},
	">": {x: 1, y: 0},
	"<": {x: -1, y: 0},
}

func part1(input []string) {
	parts := splitInput(input, "")
	grid, moves := parseGrid(parts[0]), strings.Split(strings.Join(parts[1], ""), "")
	robotPosition := findRobotPosition(grid)

	for _, move := range moves {
		_, robotPosition, grid = attemptMove(robotPosition, directions[move], grid)
	}

	totalGps := 0
	for position, value := range grid {
		if value == "O" {
			totalGps += 100*position.y + position.x
		}
	}
	fmt.Println(totalGps)
}

func part2(input []string) {
	parts := splitInput(input, "")
	for i := range parts[0] {
		parts[0][i] = strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(parts[0][i], "#", "##"),
					"O", "[]"), ".", ".."),
			"@", "@.")
	}

	grid, moves := parseGrid(parts[0]), strings.Split(strings.Join(parts[1], ""), "")
	robotPosition := findRobotPosition(grid)

	for _, move := range moves {
		_, robotPosition, grid = attemptMove(robotPosition, directions[move], grid)
	}

	totalGPS := 0
	for position, value := range grid {
		if value == "[" {
			totalGPS += 100*position.y + position.x
		}
	}
	fmt.Println(totalGPS)
}

func Run() {
	file, err := os.ReadFile("15/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	input := strings.Split(strings.TrimSpace(string(file)), "\n")

	part1(input)
	part2(input)
}

func splitInput(input []string, separator string) [][]string {
	var result [][]string
	var temp []string
	for _, line := range input {
		if line == separator {
			result = append(result, temp)
			temp = []string{}
		} else {
			temp = append(temp, line)
		}
	}
	result = append(result, temp)
	return result
}

func parseGrid(input []string) map[Point]string {
	grid := make(map[Point]string)
	for y, line := range input {
		for x, c := range strings.Split(line, "") {
			grid[Point{x, y}] = c
		}
	}
	return grid
}

func findRobotPosition(grid map[Point]string) Point {
	for position, value := range grid {
		if value == "@" {
			return position
		}
	}
	return Point{}
}

func attemptMove(robotPosition Point, direction Point, grid map[Point]string) (bool, Point, map[Point]string) {
	saveState := make(map[Point]string)
	for k, v := range grid {
		saveState[k] = v
	}

	newPosition := Point{x: robotPosition.x + direction.x, y: robotPosition.y + direction.y}

	if grid[newPosition] == "." {
		grid[newPosition] = grid[robotPosition]
		grid[robotPosition] = "."
		return true, newPosition, grid
	}

	if grid[newPosition] == "O" {
		moved, _, updatedGrid := attemptMove(newPosition, direction, grid)
		if moved {
			grid[newPosition] = grid[robotPosition]
			grid[robotPosition] = "."
			return true, newPosition, updatedGrid
		}
	}

	if grid[newPosition] == "[" || grid[newPosition] == "]" {
		if direction == directions["<"] || direction == directions[">"] {
			moved, _, updatedGrid := attemptMove(newPosition, direction, grid)
			if moved {
				grid[newPosition] = grid[robotPosition]
				grid[robotPosition] = "."
				return true, newPosition, updatedGrid
			}
		} else {
			var char, partnerChar string
			var partner Point
			if grid[newPosition] == "]" {
				char, partnerChar = "]", "["
				partner = Point{x: newPosition.x + directions["<"].x, y: newPosition.y + directions["<"].y}
			} else {
				char, partnerChar = "[", "]"
				partner = Point{x: newPosition.x + directions[">"].x, y: newPosition.y + directions[">"].y}
			}

			boxCanMove, _, newGrid := attemptMove(newPosition, direction, grid)
			partnerCanMove, _, _ := attemptMove(partner, direction, newGrid)

			if boxCanMove && partnerCanMove {
				grid[newPosition] = grid[robotPosition]
				grid[robotPosition] = "."

				newPos := Point{x: newPosition.x + direction.x, y: newPosition.y + direction.y}
				grid[newPos] = char

				newPos = Point{x: partner.x + direction.x, y: partner.y + direction.y}
				grid[newPos] = partnerChar

				grid[partner] = "."
				return true, newPosition, grid
			}
		}
	}
	return false, robotPosition, saveState
}
