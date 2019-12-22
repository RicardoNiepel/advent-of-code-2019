package day7

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partone_small(t *testing.T) {
	tests := []struct {
		inputIntcode              string
		expectedPhaseSetting      string
		expectedMaxThrusterSignal string
	}{
		{
			"3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0",
			"4,3,2,1,0",
			"43210",
		},
		{
			"3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0",
			"0,1,2,3,4",
			"54321",
		},
		{
			"3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0",
			"1,0,4,3,2",
			"65210",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.inputIntcode), func(t *testing.T) {
			phaseSetting, thrusterSignal := getAmplifierSeriesMaxThrusterSignal(tt.inputIntcode)
			assert.Equal(t, tt.expectedPhaseSetting, phaseSetting)
			assert.Equal(t, tt.expectedMaxThrusterSignal, thrusterSignal)
		})
	}
}

func Test_partone_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	phaseSetting, thrusterSignal := getAmplifierSeriesMaxThrusterSignal(input)

	assert.Equal(t, "0,2,4,3,1", phaseSetting)
	assert.Equal(t, "567045", thrusterSignal)
}

func Test_parttwo_small(t *testing.T) {
	tests := []struct {
		inputIntcode              string
		expectedPhaseSetting      string
		expectedMaxThrusterSignal string
	}{
		{
			"3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5",
			"9,8,7,6,5",
			"139629729",
		},
		{
			"3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10",
			"9,7,8,5,6",
			"18216",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.inputIntcode), func(t *testing.T) {
			phaseSetting, thrusterSignal := getAmplifierSeriesWithFeedbackLoopMaxThrusterSignal(tt.inputIntcode)
			assert.Equal(t, tt.expectedPhaseSetting, phaseSetting)
			assert.Equal(t, tt.expectedMaxThrusterSignal, thrusterSignal)
		})
	}
}

func Test_parttwo_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	phaseSetting, thrusterSignal := getAmplifierSeriesWithFeedbackLoopMaxThrusterSignal(input)

	assert.Equal(t, "6,5,7,8,9", phaseSetting)
	assert.Equal(t, "39016654", thrusterSignal)
}

func Test_day5_partone_small(t *testing.T) {
	tests := []struct {
		inputIntcode    string
		input           string
		expectedIntcode string
		expectedOutput  string
	}{
		{
			"3,0,4,0,99",
			"9876",
			"9876,0,4,0,99",
			"9876",
		},
		{
			"1002,4,3,4,33",
			"",
			"1002,4,3,4,99",
			"",
		},
		{
			"1101,100,-1,4,0",
			"",
			"1101,100,-1,4,99",
			"",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			result, output := run(tt.inputIntcode, tt.input)
			assert.Equal(t, tt.expectedIntcode, result)
			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}

func Test_day5_partone_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input_day5.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	_, output := run(input, "1")

	expected := "0\n0\n0\n0\n0\n0\n0\n0\n0\n15386262"
	assert.Equal(t, expected, output)
}

func Test_day5_parttwo_small(t *testing.T) {
	tests := []struct {
		inputIntcode    string
		input           string
		expectedIntcode string
		expectedOutput  string
	}{
		{
			"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			"7",
			"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,7,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			"999",
		},
		{
			"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			"8",
			"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,1000,8,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			"1000",
		},
		{
			"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			"9",
			"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,1001,9,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			"1001",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			result, output := run(tt.inputIntcode, tt.input)
			assert.Equal(t, tt.expectedIntcode, result)
			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}

func Test_day5_parttwo_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input_day5.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	_, output := run(input, "5")

	expected := "10376124"
	assert.Equal(t, expected, output)
}

func Test_day2_partone_small(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{"1,0,0,0,99", "2,0,0,0,99"},
		{"2,3,0,3,99", "2,3,0,6,99"},
		{"2,4,4,5,99,0", "2,4,4,5,99,9801"},
		{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			result, _ := run(tt.input, "")
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func Test_day2_partone_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input_day2.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	result, _ := runWithNounVerb(input, 12, 2, "")

	expected := "5534943,12,2,2,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,24,1,19,9,27,1,23,9,30,1,10,27,34,1,13,31,39,1,35,10,43,2,39,9,129,1,43,13,134,1,5,47,135,1,6,51,137,1,13,55,142,1,59,6,144,1,63,10,148,2,67,6,296,1,71,5,297,2,75,10,1188,1,79,6,1190,1,83,5,1191,1,87,6,1193,1,91,13,1198,1,95,6,1200,2,99,10,4800,1,103,6,4802,2,6,107,9604,1,13,111,9609,2,115,10,38436,1,119,5,38437,2,10,123,153748,2,127,9,461244,1,5,131,461245,2,10,135,1844980,2,139,9,5534940,1,143,2,5534942,1,5,147,0,99,2,0,14,0"
	assert.Equal(t, expected, result)
}

func Test_day2_parttwo_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input_day2.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	result, _ := runWithNounVerb(input, 76, 3, "")

	expected := "19690720,76,3,2,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,152,1,19,9,155,1,23,9,158,1,10,27,162,1,13,31,167,1,35,10,171,2,39,9,513,1,43,13,518,1,5,47,519,1,6,51,521,1,13,55,526,1,59,6,528,1,63,10,532,2,67,6,1064,1,71,5,1065,2,75,10,4260,1,79,6,4262,1,83,5,4263,1,87,6,4265,1,91,13,4270,1,95,6,4272,2,99,10,17088,1,103,6,17090,2,6,107,34180,1,13,111,34185,2,115,10,136740,1,119,5,136741,2,10,123,546964,2,127,9,1640892,1,5,131,1640893,2,10,135,6563572,2,139,9,19690716,1,143,2,19690719,1,5,147,0,99,2,0,14,0"
	assert.Equal(t, expected, result)
}
