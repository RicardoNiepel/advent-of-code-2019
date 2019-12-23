package day11

import (
	"fmt"
	"strconv"
)

type Color int

const (
	Black Color = 0
	White Color = 1
)

type TurnDirection int

const (
	Left90  TurnDirection = 0
	Right90 TurnDirection = 1
)

type Direction int

const (
	Top   Direction = 0
	Right Direction = 1
	Down  Direction = 2
	Left  Direction = 3
)

type Robot struct {
	x         int64
	y         int64
	direction Direction
	panels    map[string]*Panel
}

type Panel struct {
	x     int64
	y     int64
	color Color
}

func newRobot() *Robot {
	robot := Robot{x: 0, y: 0, direction: Top}
	robot.panels = make(map[string]*Panel)
	return &robot
}

func (robot *Robot) run(input <-chan string, output chan<- string, foreever bool) {
	for foreever || len(input) > 0 {
		output <- strconv.Itoa(int(robot.getCurrentColor()))
		paintCode, _ := strconv.Atoi(<-input)
		robot.processPaint(Color(paintCode))
		turnDirectionCode, _ := strconv.Atoi(<-input)
		robot.processDirection(TurnDirection(turnDirectionCode))
	}
}

func (robot *Robot) getCurrentColor() Color {
	panel, prs := robot.panels[fmt.Sprintf("%v,%v", robot.x, robot.y)]
	if !prs {
		panel = &Panel{x: robot.x, y: robot.y, color: Black}
		robot.panels[fmt.Sprintf("%v,%v", robot.x, robot.y)] = panel
	}
	return panel.color
}

func (robot *Robot) processPaint(color Color) {
	panel := robot.panels[fmt.Sprintf("%v,%v", robot.x, robot.y)]
	panel.color = color
}

func (robot *Robot) processDirection(turnDir TurnDirection) {
	switch turnDir {
	case Left90:
		robot.direction = robot.direction - 1
		if robot.direction < 0 {
			robot.direction = 3
		}
	case Right90:
		robot.direction = (robot.direction + 1) % 4
	default:
		panic("TurnDirection not supported " + fmt.Sprint(turnDir))
	}

	switch robot.direction {
	case Top:
		robot.y++
	case Right:
		robot.x++
	case Down:
		robot.y--
	case Left:
		robot.x--
	default:
		panic("Direction not supported " + fmt.Sprint(robot.direction))
	}
}
