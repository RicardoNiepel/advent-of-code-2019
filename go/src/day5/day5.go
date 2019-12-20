package day5

import (
	"fmt"
	"strconv"
	"strings"
)

func run(intcode string, input string) (finalIntcode string, output string) {

	intcodeProgram := newIntcodeProgram(intcode)

	output = intcodeProgram.execute(input)
	finalIntcode = intcodeProgram.currentProgram
	return
}

func runWithNounVerb(intcode string, input string, noun int, verb int) (finalIntcode string, output string) {

	intcodeProgram := newIntcodeProgramWithNounVerb(intcode, noun, verb)

	output = intcodeProgram.execute(input)
	finalIntcode = intcodeProgram.currentProgram
	return
}

type ParameterMode int

const (
	PositionMode  ParameterMode = 0
	ImmediateMode ParameterMode = 1
)

type IntcodeProgram struct {
	initialProgram string
	currentProgram string

	positions []string
}

func newIntcodeProgram(program string) *IntcodeProgram {
	return newIntcodeProgramWithNounVerb(program, -1, -1)
}

func newIntcodeProgramWithNounVerb(program string, noun int, verb int) *IntcodeProgram {
	positions := strings.Split(program, ",")
	if noun != -1 {
		if noun < 0 || noun > 99 {
			panic("noun out of range 0-99")
		}
		positions[1] = fmt.Sprint(noun)
	}
	if verb != -1 {
		if verb < 0 || verb > 99 {
			panic("verb out of range 0-99")
		}
		positions[2] = fmt.Sprint(verb)
	}
	program = strings.Join(positions, ",")

	instance := IntcodeProgram{initialProgram: program}
	instance.positions = strings.Split(program, ",")
	return &instance
}

func (ip *IntcodeProgram) execute(input string) (output string) {
	currentPosition := 0
	output = ""

	for {
		opcode := ip.positions[currentPosition]

		modeParam1 := PositionMode
		modeParam2 := PositionMode
		if len(opcode) > 2 {
			if len(opcode) == 4 {
				if opcode[0] == '1' {
					modeParam2 = ImmediateMode
				}
				if opcode[1] == '1' {
					modeParam1 = ImmediateMode
				}
				opcode = opcode[2:]
			} else if len(opcode) == 3 {
				if opcode[0] == '1' {
					modeParam1 = ImmediateMode
				}
				opcode = opcode[1:]
			} else {
				panic("not implemented")
			}
		}

		if opcode == "99" {
			break
		}

		switch {
		case opcode == "1" || opcode == "01":
			firstOperant := ip.getParameter(currentPosition+1, modeParam1)
			secondOperant := ip.getParameter(currentPosition+2, modeParam2)
			storePosition, _ := strconv.Atoi(ip.positions[currentPosition+3])
			ip.positions[storePosition] = fmt.Sprint(firstOperant + secondOperant)
			currentPosition += 4
		case opcode == "2" || opcode == "02":
			firstOperant := ip.getParameter(currentPosition+1, modeParam1)
			secondOperant := ip.getParameter(currentPosition+2, modeParam2)
			storePosition, _ := strconv.Atoi(ip.positions[currentPosition+3])
			ip.positions[storePosition] = fmt.Sprint(firstOperant * secondOperant)
			currentPosition += 4
		case opcode == "3" || opcode == "03":
			storePosition, _ := strconv.Atoi(ip.positions[currentPosition+1])
			ip.positions[storePosition] = input
			currentPosition += 2
		case opcode == "4" || opcode == "04":
			firstOperant := ip.getParameter(currentPosition+1, modeParam1)

			if output != "" {
				output = fmt.Sprintln(output)
			}
			output = fmt.Sprint(output, firstOperant)
			currentPosition += 2
		case opcode == "5" || opcode == "05": // jump-if-true
			firstParam := ip.getParameter(currentPosition+1, modeParam1)
			secParam := ip.getParameter(currentPosition+2, modeParam2)
			if firstParam != 0 {
				currentPosition = secParam
			} else {
				currentPosition += 3
			}
		case opcode == "6" || opcode == "06": // jump-if-false
			firstParam := ip.getParameter(currentPosition+1, modeParam1)
			secParam := ip.getParameter(currentPosition+2, modeParam2)
			if firstParam == 0 {
				currentPosition = secParam
			} else {
				currentPosition += 3
			}
		case opcode == "7" || opcode == "07": // less than
			firstParam := ip.getParameter(currentPosition+1, modeParam1)
			secParam := ip.getParameter(currentPosition+2, modeParam2)
			storePosition, _ := strconv.Atoi(ip.positions[currentPosition+3])
			if firstParam < secParam {
				ip.positions[storePosition] = "1"
			} else {
				ip.positions[storePosition] = "0"
			}
			currentPosition += 4
		case opcode == "8" || opcode == "08": // equals
			firstParam := ip.getParameter(currentPosition+1, modeParam1)
			secParam := ip.getParameter(currentPosition+2, modeParam2)
			storePosition, _ := strconv.Atoi(ip.positions[currentPosition+3])
			if firstParam == secParam {
				ip.positions[storePosition] = "1"
			} else {
				ip.positions[storePosition] = "0"
			}
			currentPosition += 4
		default:
			panic("unknown opcode")
		}
	}

	ip.currentProgram = strings.Join(ip.positions, ",")
	return
}

func (ip *IntcodeProgram) getParameter(position int, mode ParameterMode) int {
	var parameter int

	switch mode {
	case PositionMode:
		lookupPosition, _ := strconv.Atoi(ip.positions[position])
		parameter, _ = strconv.Atoi(ip.positions[lookupPosition])
	case ImmediateMode:
		parameter, _ = strconv.Atoi(ip.positions[position])
	}

	return parameter
}
