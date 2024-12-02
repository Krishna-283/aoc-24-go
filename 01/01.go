package day1

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	input, err := os.ReadFile("01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	strInput := string(input)
	lines := strings.Split(strInput, "\n")

	left := make([]int, 0, len(lines))
	right := make([]int, 0, len(lines))

	for _, line := range lines {
		parts := strings.Fields(line)

		if len(parts) == 2 {
			leftNum, _ := strconv.Atoi(parts[0])
			rightNum, _ := strconv.Atoi(parts[1])

			left = append(left, leftNum)
			right = append(right, rightNum)
		}
	}

	slices.Sort(left)
	slices.Sort(right)

	totalDistance := 0
	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]
		totalDistance += int(math.Abs(float64(distance)))
	}

	fmt.Println("Total Distance: ", totalDistance)

	//  PART 2
	frequency := make(map[int]int)
	for _, num := range right {
		frequency[num]++
	}

	similarityScore := 0
	for _, val := range left {
		similarityScore += (val * frequency[val])
	}

	fmt.Println("Similarity Score: ", similarityScore)
}
