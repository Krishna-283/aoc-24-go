package day9

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(input string) {
	var blocks []int
	fileId := -1

	for i, char := range input {
		blockLength, _ := strconv.Atoi(string(char))

		if i%2 == 0 {
			fileId++
			for j := 0; j < blockLength; j++ {
				blocks = append(blocks, fileId)
			}
		} else {
			for j := 0; j < blockLength; j++ {
				blocks = append(blocks, -1)
			}
		}
	}

	freeIdx := 0
	for endIdx := len(blocks) - 1; endIdx > freeIdx; endIdx-- {
		if blocks[endIdx] != -1 {
			for {
				if freeIdx > endIdx {
					break
				}
				if blocks[freeIdx] == -1 {
					break
				}
				freeIdx++
			}

			if freeIdx > endIdx {
				break
			}
			blocks[freeIdx] = blocks[endIdx]
			blocks[endIdx] = -1
		}
	}

	calculateChecksum(blocks)
}

func part2(input string) {
	blocks, filesData := make([]int, 0), make([][2]int, 0)
	fileId := -1

	for i, char := range input {
		blockLength, _ := strconv.Atoi(string(char))

		if i%2 == 0 {
			fileId++
			filesData = append(filesData, [2]int{len(blocks), blockLength})
			for j := 0; j < blockLength; j++ {
				blocks = append(blocks, fileId)
			}
		} else {
			for j := 0; j < blockLength; j++ {
				blocks = append(blocks, -1)
			}
		}
	}

	for fileIdx := len(filesData) - 1; fileIdx > 0; fileIdx-- {
		file := filesData[fileIdx]

		found, sp := getFreeSpace(blocks, file[1])
		if found && sp < file[0] {
			for z := 0; z < file[1]; z++ {
				blocks[sp+z] = fileIdx
				blocks[file[0]+z] = -1
			}
		}
	}

	calculateChecksum(blocks)
}

func Run() {
	file, err := os.ReadFile("09/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	input := strings.TrimSpace(string(file))
	part1(input)
	part2(input)
}

func getFreeSpace(disk []int, length int) (bool, int) {
	freeSpaceLength, freeSpaceStart := 0, -1

	for idx, block := range disk {
		if block == -1 {
			if freeSpaceStart == -1 {
				freeSpaceStart = idx
			}
			freeSpaceLength++
			if freeSpaceLength >= length {
				return true, freeSpaceStart
			}
		} else {
			freeSpaceStart = -1
			freeSpaceLength = 0
		}
	}
	return false, -1
}

func calculateChecksum(blocks []int) {
	checksum := 0
	for i, block := range blocks {
		if block != -1 {
			checksum += i * block
		}
	}

	fmt.Println(checksum)
}
