package day16

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Coordinate struct {
	x, y int
}

type Problem struct {
	start Coordinate
	end   Coordinate
	grid  map[Coordinate]rune
}

type Dir int

type Position struct {
	coord Coordinate
	dir   Dir
}

type State struct {
	pos    Position
	points int
}

const (
	North Dir = iota
	East
	South
	West
)

func part1(problem Problem) {

	bestScoreResult := 0
	start := State{Position{problem.start, East}, 0}
	openStatesSet := make(map[State]bool)
	openStatesSet[start] = true
	openPQ := make(PriorityQueue, 1)
	openPQ[0] = &PQ{
		state:    start,
		priority: h(start, problem),
		index:    0,
	}
	heap.Init(&openPQ)

	gScore := initializeScores(problem, math.MaxInt)
	fScore := initializeScores(problem, math.MaxInt)

	gScore[start.pos] = 0
	fScore[start.pos] = h(start, problem)

	for len(openStatesSet) > 0 {
		cur := heap.Pop(&openPQ).(*PQ).state
		delete(openStatesSet, cur)

		if cur.pos.coord == problem.end {
			bestScoreResult = cur.points
			break
		}

		for _, n := range nextStates(cur, problem) {
			tentativeGScore := gScore[cur.pos] + (n.points - cur.points)
			if tentativeGScore < gScore[n.pos] {
				gScore[n.pos] = tentativeGScore
				fScore[n.pos] = tentativeGScore + h(n, problem)
				if _, ok := openStatesSet[n]; !ok {
					openStatesSet[n] = true
					pqState := &PQ{state: n, priority: fScore[n.pos]}
					heap.Push(&openPQ, pqState)
				}
			}
		}
	}

	fmt.Println(bestScoreResult)
}

func part2(problem Problem) {

	allTiles := []Coordinate{}
	start := State{Position{problem.start, East}, 0}
	openStatesSet := make(map[State]bool)
	openStatesSet[start] = true

	openPQ := make(PriorityQueue, 1)
	openPQ[0] = &PQ{
		state:    start,
		priority: h(start, problem),
		index:    0,
	}
	heap.Init(&openPQ)

	cameFrom := make(map[Position][]Position)
	gScore := make(map[Position]int)

	for coord := range problem.grid {
		for _, dir := range []Dir{North, East, South, West} {
			gScore[Position{coord, dir}] = math.MaxInt
		}
	}

	gScore[start.pos] = 0
	fScore := make(map[Position]int)

	for coord := range problem.grid {
		for _, dir := range []Dir{North, East, South, West} {
			fScore[Position{coord, dir}] = math.MaxInt
		}
	}

	fScore[start.pos] = h(start, problem)
	bestpoints := math.MaxInt

	for len(openStatesSet) > 0 {
		cur := heap.Pop(&openPQ).(*PQ).state
		delete(openStatesSet, cur)
		if cur.points > bestpoints {
			allTiles = findAllLocs(cameFrom, problem.end)
			break
		}

		if cur.pos.coord == problem.end {
			if cur.points < bestpoints {
				bestpoints = cur.points
			}
		}

		for _, n := range nexts(cur, problem) {
			tentative_gScore := gScore[cur.pos] + (n.points - cur.points)

			if tentative_gScore == gScore[n.pos] {
				cameFrom[n.pos] = append(cameFrom[n.pos], cur.pos)
			} else if tentative_gScore < gScore[n.pos] {
				cameFrom[n.pos] = []Position{cur.pos}
				gScore[n.pos] = tentative_gScore
				fScore[n.pos] = tentative_gScore + h(n, problem)
				if _, ok := openStatesSet[n]; !ok {
					openStatesSet[n] = true
					pqstate := &PQ{state: n, priority: fScore[n.pos]}
					heap.Push(&openPQ, pqstate)
				}
			}
		}
	}

	fmt.Println(len(allTiles))
}

func Run() {
	file, err := os.ReadFile("16/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(file)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make(map[Coordinate]rune)
	var start, end Coordinate

	for y, line := range lines {
		for x, c := range line {
			coord := Coordinate{x, y}
			if c == 'S' {
				start = coord
				grid[coord] = '.'
			} else if c == 'E' {
				end = coord
				grid[coord] = '.'
			} else {
				grid[coord] = c
			}
		}
	}
	problem := Problem{start, end, grid}

	part1(problem)
	part2(problem)
}

func h(next State, problem Problem) int {
	return abs(next.pos.coord.x-problem.end.x) + abs(next.pos.coord.y-problem.end.y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func delta(dir Dir) (int, int) {
	switch dir {
	case North:
		return 0, -1
	case East:
		return 1, 0
	case South:
		return 0, 1
	case West:
		return -1, 0
	default:
		return 0, 0
	}
}

func nextStates(cur State, problem Problem) []State {
	x, y, dir := cur.pos.coord.x, cur.pos.coord.y, cur.pos.dir
	points := cur.points
	grid := problem.grid

	var ns []State
	left, right := (dir-1+4)%4, (dir+1)%4
	dx, dy := delta(dir)

	ns = append(ns,
		State{Position{cur.pos.coord, left}, points + 1000},
		State{Position{cur.pos.coord, right}, points + 1000},
	)

	forward := Coordinate{x + dx, y + dy}
	if grid[forward] != '#' {
		ns = append(ns, State{Position{forward, dir}, points + 1})
	}

	return ns
}

func findAllTiles(cameFrom map[Position][]Position, current Coordinate) []Coordinate {
	tilesSet := make(map[Coordinate]bool)
	tilesSet[current] = true

	currentSet := []Position{}
	for _, dir := range []Dir{North, East, South, West} {
		pos := Position{current, dir}
		if _, ok := cameFrom[pos]; ok {
			currentSet = append(currentSet, pos)
		}
	}

	for len(currentSet) > 0 {
		nextSet := []Position{}
		for _, cur := range currentSet {
			tilesSet[cur.coord] = true
			if nextPositions, ok := cameFrom[cur]; !ok {
				continue
			} else {
				nextSet = append(nextSet, nextPositions...)
			}
		}
		currentSet = nextSet
	}

	var tilesList []Coordinate
	for loc := range tilesSet {
		tilesList = append(tilesList, loc)
	}

	return tilesList
}

func initializeScores(problem Problem, defaultValue int) map[Position]int {
	scores := make(map[Position]int)
	for coord := range problem.grid {
		for _, dir := range []Dir{North, East, South, West} {
			scores[Position{coord, dir}] = defaultValue
		}
	}
	return scores
}

func nexts(cur State, problem Problem) []State {
	x, y, dir := cur.pos.coord.x, cur.pos.coord.y, cur.pos.dir
	points := cur.points
	grid := problem.grid

	ns := []State{}
	left, right := dir-1, dir+1
	if left < 0 {
		left += 4
	}
	if right >= 4 {
		right -= 4
	}
	dx, dy := delta(dir)
	ns = append(ns,
		State{Position{cur.pos.coord, left}, points + 1000},
		State{Position{cur.pos.coord, right}, points + 1000},
	)
	forward := Coordinate{x + dx, y + dy}
	if grid[forward] != '#' {
		ns = append(ns, State{Position{forward, dir}, points + 1})
	}
	return ns
}

func findAllLocs(cameFrom map[Position][]Position, current Coordinate) []Coordinate {
	locsSet := make(map[Coordinate]bool)
	locsSet[current] = true

	currentset := []Position{}
	for _, dir := range []Dir{North, East, South, West} {
		pos := Position{current, dir}
		_, ok := cameFrom[pos]
		if ok {
			currentset = append(currentset, pos)
		}
	}
	var done bool = false
	for {
		for _, cur := range currentset {
			locsSet[cur.coord] = true
			if _, ok := cameFrom[cur]; !ok {
				done = true
			}
		}
		if done {
			break
		}
		nextset := []Position{}
		for _, cur := range currentset {
			nextset = append(nextset, cameFrom[cur]...)
		}
		currentset = nextset
	}

	locsList := []Coordinate{}
	for loc := range locsSet {
		locsList = append(locsList, loc)
	}
	return locsList
}
