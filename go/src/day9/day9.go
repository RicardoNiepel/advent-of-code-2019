package day9

import (
	"fmt"
	"strconv"
	"strings"
)

func getAmplifierSeriesWithFeedbackLoopMaxThrusterSignal(intcode string) (phaseSetting string, thrusterSignal string) {
	possiblePhaseSettings := []string{"5", "6", "7", "8", "9"}

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

func runAmplifierSeries(intcode string, phaseSetting ...string) (thrusterSignal string) {

	ampA := NewIntcodeProgram(intcode)
	ampB := NewIntcodeProgram(intcode)
	ampC := NewIntcodeProgram(intcode)
	ampD := NewIntcodeProgram(intcode)
	ampE := NewIntcodeProgram(intcode)

	inputAChan := make(chan string, 2)
	outputAChan := make(chan string, 1)
	outputBChan := make(chan string, 1)
	outputCChan := make(chan string, 1)
	outputDChan := make(chan string, 1)
	doneA := make(chan bool, 1)
	doneB := make(chan bool, 1)
	doneC := make(chan bool, 1)
	doneD := make(chan bool, 1)
	doneE := make(chan bool, 1)

	inputAChan <- phaseSetting[0]
	inputAChan <- "0"
	outputAChan <- phaseSetting[1]
	outputBChan <- phaseSetting[2]
	outputCChan <- phaseSetting[3]
	outputDChan <- phaseSetting[4]

	go ampA.Execute(inputAChan, outputAChan, doneA)
	go ampB.Execute(outputAChan, outputBChan, doneB)
	go ampC.Execute(outputBChan, outputCChan, doneC)
	go ampD.Execute(outputCChan, outputDChan, doneD)
	go ampE.Execute(outputDChan, inputAChan, doneE)

	<-doneE
	return <-inputAChan
}

func run(intcode string, inputs ...string) (finalIntcode string, output string) {
	intcodeProgram := NewIntcodeProgram(intcode)

	inputChan := make(chan string)
	for _, i := range inputs {
		go func() { inputChan <- i }()
	}

	outputChan := make(chan string, 100)
	done := make(chan bool, 1)
	go intcodeProgram.Execute(inputChan, outputChan, done)
	<-done

	var outputs []string
	for len(outputChan) > 0 {
		outputs = append(outputs, <-outputChan)
	}
	output = strings.Join(outputs, "\n")

	finalIntcode = intcodeProgram.currentProgram
	return
}

func runWithNounVerb(intcode string, noun int, verb int, inputs ...string) (finalIntcode string, output string) {
	intcodeProgram := newIntcodeProgramWithNounVerb(intcode, noun, verb)

	inputChan := make(chan string)
	for _, i := range inputs {
		go func() { inputChan <- i }()
	}

	outputChan := make(chan string, 100)
	done := make(chan bool, 1)
	go intcodeProgram.Execute(inputChan, outputChan, done)
	<-done

	var outputs []string
	for len(outputChan) > 0 {
		outputs = append(outputs, <-outputChan)
	}
	output = strings.Join(outputs, "\n")

	finalIntcode = intcodeProgram.currentProgram
	return
}

type ParameterMode int

const (
	PositionMode  ParameterMode = 0
	ImmediateMode ParameterMode = 1
	RelativeMode  ParameterMode = 2
)

type IntcodeProgram struct {
	initialProgram string
	currentProgram string

	positions    []string
	relativeBase int
}

func NewIntcodeProgram(program string) *IntcodeProgram {
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

	instance := IntcodeProgram{initialProgram: program, relativeBase: 0}
	instance.positions = strings.Split(program, ",")
	return &instance
}

func toParameterMode(modeBit byte) ParameterMode {
	tmp, _ := strconv.Atoi(string(modeBit))
	return ParameterMode(tmp)
}

func (ip *IntcodeProgram) Execute(input <-chan string, output chan<- string, done chan<- bool) {
	currentPosition := 0

	for {
		opcode := ip.positions[currentPosition]
		originalOpcode := opcode

		modeParam1 := PositionMode
		modeParam2 := PositionMode
		modeParam3 := PositionMode
		if len(opcode) > 2 {
			opcode = originalOpcode[len(opcode)-2:]
			modeBits := reverse(originalOpcode[:len(originalOpcode)-2])

			if len(modeBits) > 2 {
				modeParam3 = toParameterMode(modeBits[2])
			}
			if len(modeBits) > 1 {
				modeParam2 = toParameterMode(modeBits[1])
			}
			if len(modeBits) > 0 {
				modeParam1 = toParameterMode(modeBits[0])
			}
		}

		if opcode == "99" {
			break
		}

		switch {
		case opcode == "1" || opcode == "01":
			firstOperant := ip.getParameterRead(currentPosition+1, modeParam1)
			secondOperant := ip.getParameterRead(currentPosition+2, modeParam2)
			storePosition := ip.getParameterWrite(currentPosition+3, modeParam3)
			*storePosition = fmt.Sprint(firstOperant + secondOperant)
			currentPosition += 4
		case opcode == "2" || opcode == "02":
			firstOperant := ip.getParameterRead(currentPosition+1, modeParam1)
			secondOperant := ip.getParameterRead(currentPosition+2, modeParam2)
			storePosition := ip.getParameterWrite(currentPosition+3, modeParam3)
			*storePosition = fmt.Sprint(firstOperant * secondOperant)
			currentPosition += 4
		case opcode == "3" || opcode == "03": // input
			storePosition := ip.getParameterWrite(currentPosition+1, modeParam1)
			*storePosition = <-input
			currentPosition += 2
		case opcode == "4" || opcode == "04": // ouput
			firstOperant := ip.getParameterRead(currentPosition+1, modeParam1)
			output <- fmt.Sprint(firstOperant)
			currentPosition += 2
		case opcode == "5" || opcode == "05": // jump-if-true
			firstParam := ip.getParameterRead(currentPosition+1, modeParam1)
			secParam := ip.getParameterRead(currentPosition+2, modeParam2)
			if firstParam != 0 {
				currentPosition = secParam
			} else {
				currentPosition += 3
			}
		case opcode == "6" || opcode == "06": // jump-if-false
			firstParam := ip.getParameterRead(currentPosition+1, modeParam1)
			secParam := ip.getParameterRead(currentPosition+2, modeParam2)
			if firstParam == 0 {
				currentPosition = secParam
			} else {
				currentPosition += 3
			}
		case opcode == "7" || opcode == "07": // less than
			firstParam := ip.getParameterRead(currentPosition+1, modeParam1)
			secParam := ip.getParameterRead(currentPosition+2, modeParam2)
			storePosition := ip.getParameterWrite(currentPosition+3, modeParam3)

			if firstParam < secParam {
				*storePosition = "1"
			} else {
				*storePosition = "0"
			}
			currentPosition += 4
		case opcode == "8" || opcode == "08": // equals
			firstParam := ip.getParameterRead(currentPosition+1, modeParam1)
			secParam := ip.getParameterRead(currentPosition+2, modeParam2)
			storePosition := ip.getParameterWrite(currentPosition+3, modeParam3)

			if firstParam == secParam {
				*storePosition = "1"
			} else {
				*storePosition = "0"
			}
			currentPosition += 4
		case opcode == "9" || opcode == "09": // adjusts the relative base
			firstParam := ip.getParameterRead(currentPosition+1, modeParam1)
			ip.relativeBase += firstParam
			currentPosition += 2
		default:
			panic("unknown opcode:" + originalOpcode)
		}
	}

	ip.currentProgram = strings.Join(ip.positions, ",")

	done <- true
}

func (ip *IntcodeProgram) getParameterRead(position int, mode ParameterMode) int {
	var parameter int

	switch mode {
	case PositionMode:
		lookupPosition, _ := strconv.Atoi(ip.positions[position])
		ensureMemory(&ip.positions, lookupPosition)
		parameter, _ = strconv.Atoi(ip.positions[lookupPosition])
	case RelativeMode:
		relativeLookupPosition, _ := strconv.Atoi(ip.positions[position])
		ensureMemory(&ip.positions, ip.relativeBase+relativeLookupPosition)
		parameter, _ = strconv.Atoi(ip.positions[ip.relativeBase+relativeLookupPosition])
	case ImmediateMode:
		parameter, _ = strconv.Atoi(ip.positions[position])
	}

	return parameter
}

func (ip *IntcodeProgram) getParameterWrite(position int, mode ParameterMode) *string {
	switch mode {
	case PositionMode:
		lookupPosition, _ := strconv.Atoi(ip.positions[position])
		ensureMemory(&ip.positions, lookupPosition)
		return &ip.positions[lookupPosition]
	case RelativeMode:
		relativeLookupPosition, _ := strconv.Atoi(ip.positions[position])
		ensureMemory(&ip.positions, ip.relativeBase+relativeLookupPosition)
		return &ip.positions[ip.relativeBase+relativeLookupPosition]
	}

	panic("not supported " + fmt.Sprint(position) + " " + fmt.Sprint(mode))
}

func ensureMemory(positions *[]string, index int) {
	if !(index < len(*positions)) {
		*positions = append(*positions, make([]string, index+1-len(*positions))...)
	}
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

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
