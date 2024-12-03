package main

import (
	day1 "2024/01"
	day2 "2024/02"
	day3 "2024/03"
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
	default:
		log.Println("Day not implemented yet.")
	}
}
