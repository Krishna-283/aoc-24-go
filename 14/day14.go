package day14

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

const maxX int = 101
const maxY int = 103

type Robot struct {
	X, Y, Dx, Dy int
}

func part1(data string) {
	quadCount := [2][2]int{}
	midX, midY := maxX/2, maxY/2

	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		robot := parseRobot(line)
		robot.updatePosition()

		if robot.X == midX || robot.Y == midY {
			continue
		}

		quadX, quadY := 0, 0
		if robot.X > midX {
			quadX = 1
		}
		if robot.Y > midY {
			quadY = 1
		}
		quadCount[quadY][quadX]++
	}

	result := 1
	for _, row := range quadCount {
		for _, count := range row {
			result *= count
		}
	}
	fmt.Println(result)
}

func part2(data string) {
	midX, midY := maxX/2, maxY/2

	robots := []Robot{}
	for _, line := range strings.Split(data, "\n") {
		robot := parseRobot(line)
		robots = append(robots, robot)
	}

	result := 0
	minFactor := math.MaxInt
	robotsCopy := []Robot{}
	maxIter := maxX * maxY

	for i := 0; i < maxIter; i += 1 {
		quadCount := [2][2]int{}
		for r := range robots {
			robots[r].X += robots[r].Dx
			robots[r].Y += robots[r].Dy

			robots[r].X = mod(robots[r].X, maxX)
			robots[r].Y = mod(robots[r].Y, maxY)

			if robots[r].X == midX || robots[r].Y == midY {
				continue
			}

			quadX, quadY := 0, 0
			if robots[r].X > midX {
				quadX = 1
			}
			if robots[r].Y > midY {
				quadY = 1
			}
			quadCount[quadY][quadX] += 1
		}

		factor := 1
		for _, q := range quadCount {
			factor *= q[0]
			factor *= q[1]
		}

		if factor < minFactor {
			minFactor = factor
			result = i + 1
			robotsCopy = append([]Robot{}, robots...)
		}
	}

	fmt.Println(result)

	for y := 0; y < maxY; y += 1 {
		for x := 0; x < maxX; x += 1 {
			if slices.IndexFunc(robotsCopy, func(r Robot) bool { return r.X == x && r.Y == y }) == -1 {
				fmt.Printf(".")
			} else {
				fmt.Printf("+")
			}
		}
		fmt.Printf("\n")
	}
}

func Run() {
	file, err := os.ReadFile("14/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	part1(string(file))
	part2(string(file))
}

func mod(a, b int) int {
	result := a % b
	if result < 0 {
		result += b
	}
	return result
}

func parseRobot(line string) Robot {
	r := Robot{}
	fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.X, &r.Y, &r.Dx, &r.Dy)
	return r
}

func (r *Robot) updatePosition() {
	r.X = mod(r.X+r.Dx*100, maxX)
	r.Y = mod(r.Y+r.Dy*100, maxY)
}
