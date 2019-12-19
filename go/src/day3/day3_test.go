package day3

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partone_small(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{
			`R8,U5,L5,D3
U7,R6,D4,L4`,
			"6",
		},
		{
			`R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`,
			"159",
		},
		{
			`R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`,
			"135",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			result, _ := run(tt.input)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func Test_partone_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	result, _ := run(input)

	expected := "5357"
	assert.Equal(t, expected, result)
}

func Test_parttwo_small(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{
			`R8,U5,L5,D3
U7,R6,D4,L4`,
			"30",
		},
		{
			`R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`,
			"610",
		},
		{
			`R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`,
			"410",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			_, result := run(tt.input)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func Test_parttwo_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	_, result := run(input)

	expected := "101956"
	assert.Equal(t, expected, result)
}
