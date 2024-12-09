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
	default:
		log.Println("Day not implemented yet.")
	}
}
