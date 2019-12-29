package day13

import (
	"fmt"
	"strconv"
)

type TileID int

const (
	Empty  TileID = 0
	Wall   TileID = 1
	Block  TileID = 2
	Paddle TileID = 3
	Ball   TileID = 4
)

type Joystick int

const (
	Neutral Joystick = 0
	Left    Joystick = -1
	Right   Joystick = 1
)

type Game struct {
	segmentDisplay int
	joystick       Joystick
	gameMap        [][]TileID
}

func newGame() *Game {
	game := Game{}
	game.gameMap = make([][]TileID, 100)
	return &game
}

func (game *Game) run(gameEngineInput <-chan string, joystickOutput chan<- string, gameEngineDone chan bool, done chan<- bool) {
	for len(gameEngineDone) == 0 {
		xTmp := <-gameEngineInput
		x, _ := strconv.Atoi(xTmp)
		yTmp := <-gameEngineInput
		y, _ := strconv.Atoi(yTmp)

		tileIDTmp, _ := strconv.Atoi(<-gameEngineInput)
		if x == -1 && y == 0 {
			game.segmentDisplay = tileIDTmp
			blockCount := game.countTiles(Block)
			if blockCount == 0 {
				break
			}
			continue
		}
		tileID := TileID(tileIDTmp)

		xMap := game.gameMap[y]
		if xMap == nil {
			xMap = make([]TileID, 100)
			game.gameMap[y] = xMap
		}

		xMap[x] = tileID

		if tileID == Ball {
			xDiff := game.getXDiff(Paddle, Ball)
			if xDiff < 0 {
				joystickOutput <- strconv.Itoa(int(Right))
			} else if xDiff > 0 {
				joystickOutput <- strconv.Itoa(int(Left))
			} else {
				joystickOutput <- strconv.Itoa(int(Neutral))
			}
		}
	}

	done <- true
}

func (game *Game) getXDiff(a, b TileID) int {
	aX, _ := game.findTile(a)
	bX, _ := game.findTile(b)
	return aX - bX
}

func (game *Game) printMap() {
	for _, y := range game.gameMap {
		for _, x := range y {
			switch x {
			case Empty: // empty
				fmt.Print(" ")
			case Wall: // wall
				fmt.Print("#")
			case Block: // block
				fmt.Print("*")
			case Paddle: // paddle
				fmt.Print("-")
			case Ball: // ball
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}

func (game *Game) findTile(tileID TileID) (x, y int) {
	for y, yVal := range game.gameMap {
		for x, xVal := range yVal {
			if xVal == tileID {
				return x, y
			}
		}
	}
	return -1, -1
}

func (game *Game) countTiles(tileID TileID) int {
	result := 0
	for _, y := range game.gameMap {
		for _, x := range y {
			if x == tileID {
				result++
			}
		}
	}
	return result
}
