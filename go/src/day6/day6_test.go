package day6

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partone_small(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput int64
	}{
		{
			`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`,
			42,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			result := getTotalOrbitCount(tt.input)
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
	output := getTotalOrbitCount(input)

	expected := int64(194721)
	assert.Equal(t, expected, output)
}

func Test_parttwo_small(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput int64
	}{
		{
			`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`,
			4,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			result := getMinimumNumberOfOrbitalTransfers(tt.input)
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
	output := getMinimumNumberOfOrbitalTransfers(input)

	expected := int64(316)
	assert.Equal(t, expected, output)
}
