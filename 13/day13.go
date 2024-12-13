package day13

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Coordinates struct {
	X, y int
}

type Machine struct {
	BtnA, BtnB, Prize Coordinates
}

func part1(machines []Machine) {
	totalCost := 0

	for _, machine := range machines {
		btnAx, btnAy := machine.BtnA.X, machine.BtnA.y
		btnBx, btnBy := machine.BtnB.X, machine.BtnB.y
		prizeX, prizeY := machine.Prize.X, machine.Prize.y

		det := btnAx*btnBy - btnBx*btnAy
		if det == 0 {
			continue
		}

		b := ((prizeY * btnAx) - (btnAy * prizeX)) / det
		a := (prizeX - btnBx*b) / btnAx
		if a >= 0 && b >= 0 && btnAx*a+btnBx*b == prizeX && btnAy*a+btnBy*b == prizeY {
			totalCost += a*3 + b
		}
	}

	fmt.Println(totalCost)
}

func part2(machines []Machine) {
	totalCost := 0

	for _, machine := range machines {
		btnAx, btnAy := machine.BtnA.X, machine.BtnA.y
		btnBx, btnBy := machine.BtnB.X, machine.BtnB.y
		prizeX, prizeY := machine.Prize.X, machine.Prize.y

		prizeX, prizeY = prizeX+10000000000000, prizeY+10000000000000

		b := ((prizeY * btnAx) - (btnAy * prizeX)) / (-(btnBx * btnAy) + (btnBy * btnAx))
		a := (prizeX - btnBx*b) / btnAx

		if btnAx*a+btnBx*b == prizeX && btnAy*a+btnBy*b == prizeY {
			totalCost += a*3 + b
		}
	}

	fmt.Println(totalCost)
}

func Run() {
	file, err := os.ReadFile("13/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	machines := []Machine{}

	for i := 0; i < len(lines); i += 4 {
		var machine Machine

		fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &machine.BtnA.X, &machine.BtnA.y)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &machine.BtnB.X, &machine.BtnB.y)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &machine.Prize.X, &machine.Prize.y)

		machines = append(machines, machine)
	}

	part1(machines)
	part2(machines)
}
