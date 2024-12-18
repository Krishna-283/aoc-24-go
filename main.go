package main

import (
	day1 "2024/01"
	day2 "2024/02"
	day3 "2024/03"
	day4 "2024/04"
	day5 "2024/05"
	day6 "2024/06"
	day7 "2024/07"
	day8 "2024/08"
	day9 "2024/09"
	day10 "2024/10"
	day11 "2024/11"
	day12 "2024/12"
	day13 "2024/13"
	day14 "2024/14"
	day15 "2024/15"
	day16 "2024/16"
	day17 "2024/17"
	day18 "2024/18"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: go run main.go <day>")
	}

	dayArg := os.Args[1]

	day, err := strconv.Atoi(dayArg)
	if err != nil || day < 1 || day > 25 {
		log.Fatalln("Please enter valid number between 1 - 25")
	}

	switch day {
	case 1:
		day1.Run()
	case 2:
		day2.Run()
	case 3:
		day3.Run()
	case 4:
		day4.Run()
	case 5:
		day5.Run()
	case 6:
		day6.Run()
	case 7:
		day7.Run()
	case 8:
		day8.Run()
	case 9:
		day9.Run()
	case 10:
		day10.Run()
	case 11:
		day11.Run()
	case 12:
		day12.Run()
	case 13:
		day13.Run()
	case 14:
		day14.Run()
	case 15:
		day15.Run()
	case 16:
		day16.Run()
	case 17:
		day17.Run()
	case 18:
		day18.Run()
	default:
		log.Println("Day not implemented yet.")
	}
}
