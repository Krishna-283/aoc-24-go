package day18

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y int
}

var dirs = []Coordinate{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func Run() {
	file, err := os.ReadFile("18/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	input := strings.TrimSpace(string(file))
	lines := strings.Split(input, "\n")

	coordinates := make([]Coordinate, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		coordinates[i] = Coordinate{x, y}
	}

	part1 := BFS(coordinates, Coordinate{0, 0}, Coordinate{70, 70}, 1024)
	fmt.Println(part1)

	part2 := findOptimalPoint(coordinates)
	fmt.Println(part2.X, part2.Y)
}

func BFS(coordinates []Coordinate, start, end Coordinate, limit int) int {
	visited := make(map[Coordinate]bool)
	pathLengths := make(map[Coordinate]int)
	pathQueue := make([]Coordinate, 0, end.X*end.Y)

	pathQueue = append(pathQueue, start)
	visited[start] = true

	for len(pathQueue) > 0 {
		current := pathQueue[0]
		pathQueue = pathQueue[1:]

		if current == end {
			break
		}

		for _, dir := range dirs {
			next := Coordinate{
				current.X + dir.X,
				current.Y + dir.Y,
			}

			if visited[next] ||
				slices.Contains(coordinates[:limit], next) ||
				next.X < start.X || next.X > end.X ||
				next.Y < start.Y || next.Y > end.Y {
				continue
			}

			pathQueue = append(pathQueue, next)
			visited[next] = true
			pathLengths[next] = pathLengths[current] + 1
		}
	}

	return pathLengths[end]
}

func findOptimalPoint(coordinates []Coordinate) Coordinate {
	left, right := 1024, len(coordinates)-1

	for left <= right {
		mid := (left + right) / 2
		if BFS(coordinates, Coordinate{0, 0}, Coordinate{70, 70}, mid) == 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return coordinates[left-1]
}
