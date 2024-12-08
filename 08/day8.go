package day8

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

func part1(matrix *[][]string, antennas *map[string][]Point) {
	antiNodeCount := 0
	visitedNodes := map[string]int{}
	width, height := len((*matrix)[0]), len(*matrix)

	for _, antPos := range *antennas {
		if len(antPos) == 0 {
			continue
		}

		for i := 0; i < len(antPos); i++ {
			for j := 0; j < len(antPos); j++ {
				p1, p2 := antPos[i], antPos[j]
				dy, dx := p2.Y-p1.Y, p2.X-p1.X

				for _, p := range []Point{p1, p2} {
					nP1 := Point{p.X + dx, p.Y + dy}
					nP2 := Point{p.X - dx, p.Y - dy}

					for _, nP := range []Point{nP1, nP2} {
						if nP != p1 && nP != p2 && nP.X >= 0 && nP.X < width && nP.Y >= 0 && nP.Y < height &&
							visitedNodes[str(nP)] != 1 {
							visitedNodes[str(nP)] = 1
							antiNodeCount++
						}
					}
				}
			}
		}
	}
	fmt.Println(antiNodeCount)
}

func part2(matrix *[][]string, antennas *map[string][]Point) {
	antiNodeCount := 0
	visitedNodes := map[string]int{}
	width, height := len((*matrix)[0]), len(*matrix)

	for _, antPos := range *antennas {
		if len(antPos) == 0 {
			continue
		}

		if visitedNodes[str(antPos[0])] != 1 {
			visitedNodes[str(antPos[0])] = 1
			antiNodeCount++
		}

		for _, p1 := range antPos {
			for j, p2 := range antPos {
				if visitedNodes[str(antPos[j])] != 1 {
					visitedNodes[str(antPos[j])] = 1
					antiNodeCount++
				}

				dy, dx := p2.Y-p1.Y, p2.X-p1.X

				for _, p := range []Point{p1, p2} {
					nP := Point{p.X + dx, p.Y + dy}

					for isValidNode(nP, p1, p2, width, height) && (dx != 0 || dy != 0) {
						if nP != p1 && nP != p2 &&
							visitedNodes[str(nP)] != 1 {
							visitedNodes[str(nP)] = 1
							antiNodeCount++
						}
						nP = Point{nP.X + dx, nP.Y + dy}
					}

					nP = Point{p.X - dx, p.Y - dy}
					for isValidNode(nP, p1, p2, width, height) && (dx != 0 || dy != 0) {
						if nP != p1 && nP != p2 &&
							visitedNodes[str(nP)] != 1 {
							visitedNodes[str(nP)] = 1
							antiNodeCount++
						}
						nP = Point{nP.X - dx, nP.Y - dy}
					}
				}
			}
		}
	}

	fmt.Println(antiNodeCount)
}

func Run() {
	file, err := os.ReadFile("08/input.txt")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	inputStr := strings.TrimSpace(string(file))
	input := strings.Split(inputStr, "\n")

	matrix, antennas := parseInput(input)
	part1(matrix, antennas)
	part2(matrix, antennas)
}

func parseInput(input []string) (*[][]string, *map[string][]Point) {
	var matrix [][]string
	antennas := map[string][]Point{}

	for j, line := range input {
		row := strings.Split(line, "")
		for i, val := range row {
			if val != "." {
				antennas[val] = append(antennas[val], Point{i, j})
			}
		}
		matrix = append(matrix, row)
	}

	return &matrix, &antennas
}

func str(p Point) string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func isValidNode(nP, p1, p2 Point, width, height int) bool {
	return nP != p1 && nP != p2 && nP.X >= 0 && nP.X < width && nP.Y >= 0 && nP.Y < height
}
