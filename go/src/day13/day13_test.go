package day13

import (
	"day9"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partone_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	intcode := string(dat)
	gameEngine := day9.NewIntcodeProgram(intcode)
	_ = gameEngine

	game := newGame()

	inputChan := make(chan string, 10000)
	outputChan := make(chan string, 10000)
	done := make(chan bool, 1)
	done2 := make(chan bool, 1)

	go gameEngine.Execute(inputChan, outputChan, done)
	go game.run(outputChan, inputChan, done, done2)
	<-done2

	blockTilesCount := game.countTiles(Block)
	game.printMap()

	assert.Equal(t, 193, blockTilesCount)
}

func Test_parttwo_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	intcode := string(dat)

	// Play for free
	positions := strings.Split(intcode, ",")
	positions[0] = "2"
	intcode = strings.Join(positions, ",")

	gameEngine := day9.NewIntcodeProgram(intcode)
	_ = gameEngine

	game := newGame()

	inputChan := make(chan string, 10000)
	outputChan := make(chan string, 10000)
	done := make(chan bool, 1)
	done2 := make(chan bool, 1)

	go gameEngine.Execute(inputChan, outputChan, done)
	go game.run(outputChan, inputChan, done, done2)
	<-done2
	<-done

	assert.Equal(t, 10547, game.segmentDisplay)
}
