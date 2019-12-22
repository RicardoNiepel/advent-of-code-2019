package day8

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partone_small(t *testing.T) {
	tests := []struct {
		input            string
		wide             int
		tall             int
		expectedChecksum int
	}{
		{
			"123456789012",
			3,
			2,
			1,
		},
		{
			"112456789012",
			3,
			2,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			checksum := getChecksum(tt.input, tt.wide, tt.tall)
			assert.Equal(t, tt.expectedChecksum, checksum)
		})
	}
}

func Test_partone_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	output := getChecksum(input, 25, 6)

	expected := int(1224)
	assert.Equal(t, expected, output)
}

func Test_parttwo_small(t *testing.T) {
	tests := []struct {
		input          string
		wide           int
		tall           int
		expectedOutput string
	}{
		{
			"0222112222120000",
			2,
			2,
			"01\n10",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			output := decode(tt.input, tt.wide, tt.tall)
			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}

func Test_parttwo_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	output := decode(input, 25, 6)

	expected := "1111011100111101001011100\n1000010010000101001010010\n1110011100001001001010010\n1000010010010001001011100\n1000010010100001001010100\n1111011100111100110010010"
	assert.Equal(t, expected, output)
}
