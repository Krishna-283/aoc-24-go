package day12

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

var dirs = []Point{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func part1(lines []string) {
	grid := map[Point]rune{}
	maxPoint := Point{0, 0}

	for y, line := range lines {
		for x, c := range line {
			grid[Point{x, y}] = c
		}
	}

	maxPoint.X = len(lines[0])
	maxPoint.Y = len(lines)

	result := 0
	complete := map[Point]bool{}
	for i := 0; i < maxPoint.X; i++ {
		for j := 0; j < maxPoint.Y; j++ {
			basePoint := Point{i, j}
			baseRune := grid[basePoint]
			if complete[basePoint] {
				continue
			}

			blob := map[Point]bool{basePoint: true}
			touching := map[Point]map[Point]bool{}
			anyFound := true
			for anyFound {
				anyFound = false
				for p := range blob {
					for _, d := range dirs {
						pNew := Point{p.X + d.X, p.Y + d.Y}
						if grid[pNew] != baseRune {
							if touching[pNew] == nil {
								touching[pNew] = map[Point]bool{}
							}
							touching[pNew][p] = true
							continue
						}
						if blob[pNew] {
							continue
						}
						anyFound = true
						blob[pNew] = true
					}
				}
			}
			area := len(blob)
			perimeter := 0
			for _, v := range touching {
				perimeter += len(v)
			}
			result += perimeter * area

			for p := range blob {
				complete[p] = true
			}
		}
	}
	fmt.Println(result)
}

func part2(lines []string) {
	grid := map[Point]rune{}
	maxPoint := Point{0, 0}

	for y, line := range lines {
		for x, c := range line {
			grid[Point{x, y}] = c
		}
	}

	maxPoint.X = len(lines[0])
	maxPoint.Y = len(lines)

	result := 0
	complete := map[Point]bool{}
	for i := 0; i < maxPoint.X; i++ {
		for j := 0; j < maxPoint.Y; j++ {
			basePoint := Point{i, j}
			baseRune := grid[basePoint]
			if complete[basePoint] {
				continue
			}

			blob := map[Point]bool{basePoint: true}
			touching := map[Point]map[Point]bool{}
			anyFound := true

			for anyFound {
				anyFound = false
				for p := range blob {
					for _, d := range dirs {
						pNew := Point{p.X + d.X, p.Y + d.Y}
						if grid[pNew] != baseRune {
							if touching[pNew] == nil {
								touching[pNew] = map[Point]bool{}
							}
							touching[pNew][p] = true
							continue
						}
						if blob[pNew] {
							continue
						}
						anyFound = true
						blob[pNew] = true
					}
				}
			}

			area := len(blob)
			for p := range blob {
				complete[p] = true
			}

			sides := 0
			sidesToWalk := touching

			for len(sidesToWalk) > 0 {
				var pOut, pIn Point

				for po, pi := range sidesToWalk {
					pOut = po
					if len(pi) == 0 {
						break
					}

					for p := range pi {
						pIn = p
						break
					}
					break
				}

				walk(pIn, pOut, Point{-1, 0}, &sidesToWalk)
				walk(pIn, pOut, Point{1, 0}, &sidesToWalk)
				walk(pIn, pOut, Point{0, 1}, &sidesToWalk)
				walk(pIn, pOut, Point{0, -1}, &sidesToWalk)
				delete(sidesToWalk[pOut], pIn)
				if len(sidesToWalk[pOut]) == 0 {
					delete(sidesToWalk, pOut)
				}

				sides++
			}

			result += sides * area
		}
	}

	fmt.Println(result)
}

func Run() {
	file, err := os.ReadFile("12/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	part1(lines)
	part2(lines)
}

func walk(walkIn Point, walkOut Point, dir Point, sidesToWalk *map[Point]map[Point]bool) {
	for {
		walkOutNew := Point{walkOut.X + dir.X, walkOut.Y + dir.Y}
		walkInNew := Point{walkIn.X + dir.X, walkIn.Y + dir.Y}
		if len((*sidesToWalk)[walkOutNew]) == 0 {
			break
		}
		if !(*sidesToWalk)[walkOutNew][walkInNew] {
			break
		}
		walkOut = walkOutNew
		walkIn = walkInNew
		delete((*sidesToWalk)[walkOutNew], walkInNew)
		if len((*sidesToWalk)[walkOutNew]) == 0 {
			delete((*sidesToWalk), walkOutNew)
		}
	}
}
