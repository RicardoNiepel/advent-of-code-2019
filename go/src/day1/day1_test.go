package day1

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partone_small(t *testing.T) {
	tests := []struct {
		input          float64
		expectedOutput float64
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			result := run(tt.input)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func Test_partone_main(t *testing.T) {
	var totalFuel int64 = 0

	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		moduleMass, _ := strconv.ParseFloat(line, 64)
		fuel := run(moduleMass)
		totalFuel += int64(fuel)
	}

	assert.Equal(t, int64(3297866), totalFuel)
}

func Test_parttwo_small(t *testing.T) {
	tests := []struct {
		input          float64
		expectedOutput float64
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			result := runRecursive(tt.input)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func Test_parttwo_main(t *testing.T) {
	var totalFuel int64 = 0

	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		moduleMass, _ := strconv.ParseFloat(line, 64)
		fuel := runRecursive(moduleMass)
		totalFuel += int64(fuel)
	}

	assert.Equal(t, int64(4943923), totalFuel)
}
