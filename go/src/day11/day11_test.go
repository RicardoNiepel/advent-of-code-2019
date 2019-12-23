package day11

import (
	"day9"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partone_small(t *testing.T) {
	tests := []struct {
		input          string
		robotInputs    string
		expectedOutput string
	}{
		{
			`.....
.....
..^..
.....
.....`,
			"1,0,0,0,1,0,1,0,0,1,1,0,1,0",
			`.....
..<#.
...#.
.##..
.....`,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {

			inputChan := make(chan string, 100)
			for _, s := range strings.Split(tt.robotInputs, ",") {
				inputChan <- s
			}
			outputChan := make(chan string, 100)

			robot := newRobot()
			robot.run(inputChan, outputChan, false)

			//assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func Test_partone_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	intcode := string(dat)
	robotBrain := day9.NewIntcodeProgram(intcode)

	inputChan := make(chan string, 1000)
	outputChan := make(chan string, 1000)
	done := make(chan bool, 1)

	robot := newRobot()
	go robot.run(inputChan, outputChan, true)
	go robotBrain.Execute(outputChan, inputChan, done)

	<-done

	assert.Equal(t, 1964, len(robot.panels))
}

func Test_parttwo_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	intcode := string(dat)
	robotBrain := day9.NewIntcodeProgram(intcode)

	inputChan := make(chan string, 1000)
	outputChan := make(chan string, 1000)
	done := make(chan bool, 1)

	robot := newRobot()
	robot.panels["0,0"] = &Panel{x: 0, y: 0, color: White}
	go robot.run(inputChan, outputChan, true)
	go robotBrain.Execute(outputChan, inputChan, done)

	<-done

	printPanels(&robot.panels)
}

func printPanels(panels *map[string]*Panel) {
	var minX, maxX, minY, maxY int64

	for _, panel := range *panels {
		if panel.x < minX {
			minX = panel.x
		}
		if panel.x > maxX {
			maxX = panel.x
		}
		if panel.y < minY {
			minY = panel.y
		}
		if panel.y > maxY {
			maxY = panel.y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			panel, prs := (*panels)[fmt.Sprintf("%v,%v", x, y)]
			if !prs || panel.color == Black {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
