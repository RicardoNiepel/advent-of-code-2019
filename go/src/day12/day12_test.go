package day12

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partone_small1(t *testing.T) {
	input := `-1,0,2
2,-10,-7
4,-8,8
3,5,-1`
	totalEnergy := run(input, 10)

	assert.Equal(t, 179, totalEnergy)
}

func Test_partone_small2(t *testing.T) {
	input := `-8,-10,0
5,5,10
2,-7,3
9,-8,-3`
	totalEnergy := run(input, 100)

	assert.Equal(t, 1940, totalEnergy)
}

func Test_partone_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	totalEnergy := run(input, 1000)

	assert.Equal(t, 12053, totalEnergy)
}

func Test_parttwo_small1(t *testing.T) {
	input := `-1,0,2
2,-10,-7
4,-8,8
3,5,-1`

	steps := findPreviousPointInTime(input)
	assert.Equal(t, int64(2772), steps)
}

func Test_parttwo_small2(t *testing.T) {
	input := `-8,-10,0
5,5,10
2,-7,3
9,-8,-3`
	steps := findPreviousPointInTime(input)
	assert.Equal(t, int64(4686774924), steps)
}

func Test_parttwo_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	steps := findPreviousPointInTime(input)
	assert.Equal(t, int64(320380285873116), steps)
}
