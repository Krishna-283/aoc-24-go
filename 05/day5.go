package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(rules [][2]int, updates [][]int) {
	var correctUpdates [][]int
	sum := 0

	for _, update := range updates {
		if isValid(update, rules) {
			correctUpdates = append(correctUpdates, update)
		}
	}

	for _, update := range correctUpdates {
		middle := update[len(update)/2]
		sum += middle
	}

	fmt.Println(sum)
}

func part2(rules [][2]int, updates [][]int) {
	var correctUpdates [][]int
	var incorrectUpdates [][]int
	sum := 0

	for _, update := range updates {
		if isValid(update, rules) {
			correctUpdates = append(correctUpdates, update)
		} else {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	var correctedUpdates [][]int

	for _, update := range incorrectUpdates {
		sortedUpdate := sortUpdate(update, rules)
		correctedUpdates = append(correctedUpdates, sortedUpdate)
	}

	for _, update := range correctedUpdates {
		middle := update[len(update)/2]
		sum += middle
	}

	fmt.Println(sum)
}

func Run() {
	file, err := os.Open("05/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	section := 0
	var rules [][2]int
	var updates [][]int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			section++
			continue
		}

		if section == 0 {
			split := strings.Split(line, "|")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			rules = append(rules, [2]int{x, y})
		} else if section == 1 {
			nums := strings.Split(line, ",")
			var update []int

			for _, num := range nums {
				n, _ := strconv.Atoi(num)
				update = append(update, n)
			}

			updates = append(updates, update)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	part1(rules, updates)
	part2(rules, updates)
}

func isValid(update []int, rules [][2]int) bool {
	for _, rule := range rules {
		x, y := rule[0], rule[1]
		xIndex, yIndex := indexOf(update, x), indexOf(update, y)

		if xIndex != -1 && yIndex != -1 && xIndex > yIndex {
			return false
		}
	}
	return true
}

func sortUpdate(update []int, rules [][2]int) []int {
	sort.Slice(update, func(i, j int) bool {
		a, b := update[i], update[j]
		for _, rule := range rules {
			x, y := rule[0], rule[1]
			if a == x && b == y {
				return true
			}

			if b == x && a == y {
				return false
			}
		}
		return a < b
	})
	return update
}

func indexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}
