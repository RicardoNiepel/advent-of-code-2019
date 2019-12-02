package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func run(intcode string) string {
	return runExt(intcode, -1, -1)
}

func runExt(intcode string, noun int, verb int) string {
	positions := strings.Split(intcode, ",")
	currentPosition := 0
	if noun > -1 {
		positions[1] = fmt.Sprint(noun)
	}
	if verb > -1 {
		positions[2] = fmt.Sprint(verb)
	}

	for {
		opcode := positions[currentPosition]

		if opcode == "99" {
			break
		}
		if opcode != "1" && opcode != "2" {
			panic("unknown opcode")
		}

		firstOperant := getOperant(positions, currentPosition+1)
		secondOperant := getOperant(positions, currentPosition+2)
		storePosition, _ := strconv.Atoi(positions[currentPosition+3])

		switch {
		case opcode == "1":
			positions[storePosition] = fmt.Sprint(firstOperant + secondOperant)
		case opcode == "2":
			positions[storePosition] = fmt.Sprint(firstOperant * secondOperant)
		}

		currentPosition += 4
	}

	return strings.Join(positions, ",")
}

func getOperant(positions []string, position int) int {
	lookupPosition, _ := strconv.Atoi(positions[position])
	operant, _ := strconv.Atoi(positions[lookupPosition])
	return operant
}
