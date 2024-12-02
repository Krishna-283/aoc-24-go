package day2

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() {
	input, err := os.ReadFile("02/input.txt")
	if err != nil {
		log.Println(err)
	}

	lines := strings.Split(string(input), "\n")
	var data [][]int

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		strs := strings.Fields(line)
		var report []int
		for _, str := range strs {
			level, err := strconv.Atoi(str)
			if err != nil {
				log.Println("Error converting string to int:", err)
			}
			report = append(report, level)
		}
		data = append(data, report)
	}

	safeReport := 0
	for _, report := range data {
		if isSafeReport(report) {
			safeReport++
		} else if safeWithRemoval(report) {
			safeReport++
		}
	}
	fmt.Println(safeReport)
}

func safeWithRemoval(report []int) bool {
	length := len(report)
	for i := 0; i < length; i++ {
		var newReport []int
		if i == 0 {
			newReport = make([]int, length-1)
			copy(newReport, report[1:])
		} else if i == length-1 {
			newReport = make([]int, length-1)
			copy(newReport, report[:length-1])
		} else {
			newReport = append([]int{}, report[:i]...)
			newReport = append(newReport, report[i+1:]...)
		}

		if isSafeReport(newReport) {
			return true
		}
	}
	return false
}

func isSafeReport(report []int) bool {
	isDecreasing := true
	isIncreasing := true

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]

		if diff == 0 {
			return false
		}

		if diff < -3 || diff > 3 {
			return false
		}

		if diff > 0 {
			isDecreasing = false
		} else if diff < 0 {
			isIncreasing = false
		}
	}
	return isIncreasing || isDecreasing
}
