package day7

import (
	"fmt"
	"strconv"
	"strings"
)

func getAmplifierSeriesMaxThrusterSignal(intcode string) (phaseSetting string, thrusterSignal string) {
	possiblePhaseSettings := []string{"0", "1", "2", "3", "4"}

	var currentPhaseSetting []string

	for _, a := range possiblePhaseSettings {
		currentPhaseSetting = []string{a}

		for _, b := range possiblePhaseSettings {
			if contains(&currentPhaseSetting, b) {
				continue
			}
			currentPhaseSetting = append(currentPhaseSetting, b)

			for _, c := range possiblePhaseSettings {
				if contains(&currentPhaseSetting, c) {
					continue
				}
				currentPhaseSetting = append(currentPhaseSetting, c)

				for _, d := range possiblePhaseSettings {
					if contains(&currentPhaseSetting, d) {
						continue
					}
					currentPhaseSetting = append(currentPhaseSetting, d)

					for _, e := range possiblePhaseSettings {
						if contains(&currentPhaseSetting, e) {
							continue
						}
						currentPhaseSetting = append(currentPhaseSetting, e)

						thrusterSignalTmp := runAmplifierSeries(intcode, currentPhaseSetting...)
						tst, _ := strconv.Atoi(thrusterSignalTmp)
						ts, _ := strconv.Atoi(thrusterSignal)
						if tst > ts {
							phaseSetting = strings.Join(currentPhaseSetting, ",")
							thrusterSignal = thrusterSignalTmp
						}

						remove(&currentPhaseSetting, len(currentPhaseSetting)-1)
					}
					remove(&currentPhaseSetting, len(currentPhaseSetting)-1)
				}
				remove(&currentPhaseSetting, len(currentPhaseSetting)-1)
			}
			remove(&currentPhaseSetting, len(currentPhaseSetting)-1)
		}
	}

	return
}

func contains(arr *[]string, elem string) bool {
	for _, e := range *arr {
		if e == elem {
			return true
		}
	}

	return false
}

func remove(arr *[]string, index int) {
	(*arr) = append((*arr)[:index], (*arr)[index+1:]...)
}

func runAmplifierSeries(intcode string, phaseSetting ...string) (thrusterSignal string) {
	ampA := newIntcodeProgram(intcode)
	ampB := newIntcodeProgram(intcode)
	ampC := newIntcodeProgram(intcode)
	ampD := newIntcodeProgram(intcode)
	ampE := newIntcodeProgram(intcode)

	outputA := ampA.execute(phaseSetting[0], "0")
	outputB := ampB.execute(phaseSetting[1], outputA)
	outputC := ampC.execute(phaseSetting[2], outputB)
	outputD := ampD.execute(phaseSetting[3], outputC)
	outputE := ampE.execute(phaseSetting[4], outputD)
	return outputE
}

func run(intcode string, inputs ...string) (finalIntcode string, output string) {
	intcodeProgram := newIntcodeProgram(intcode)

	output = intcodeProgram.execute(inputs...)
	finalIntcode = intcodeProgram.currentProgram
	return
}

func runWithNounVerb(intcode string, noun int, verb int, inputs ...string) (finalIntcode string, output string) {
	intcodeProgram := newIntcodeProgramWithNounVerb(intcode, noun, verb)

	output = intcodeProgram.execute(inputs...)
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

func (ip *IntcodeProgram) execute(inputs ...string) (output string) {
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
			ip.positions[storePosition] = inputs[0]
			inputs = inputs[1:]
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
