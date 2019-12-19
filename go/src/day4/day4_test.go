package day4

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partone_small(t *testing.T) {
	tests := []struct {
		input          int64
		rangeStart     int64
		rangeEnd       int64
		expectedOutput bool
	}{
		{
			111111,
			111111,
			675810,
			true,
		},
		{
			223450,
			134792,
			675810,
			false,
		},
		{
			123789,
			134792,
			675810,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			result := meetsCriteria(tt.input, tt.rangeStart, tt.rangeEnd, false)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func Test_partone_possibilities(t *testing.T) {
	var startRange int64 = 134792
	var endRange int64 = 675810
	meetsCriteriaCount := 0

	for i := startRange; i <= endRange; i++ {
		if meetsCriteria(i, startRange, endRange, false) {
			meetsCriteriaCount++
		}
	}

	assert.Equal(t, 1955, meetsCriteriaCount)
}

func Test_parttwo_small(t *testing.T) {
	tests := []struct {
		input          int64
		rangeStart     int64
		rangeEnd       int64
		expectedOutput bool
	}{
		{
			112233,
			100000,
			675810,
			true,
		},
		{
			123444,
			100000,
			675810,
			false,
		},
		{
			111122,
			100000,
			675810,
			true,
		},
		{
			112222,
			100000,
			675810,
			true,
		},
		{
			111223,
			100000,
			675810,
			true,
		},
		{
			111112,
			100000,
			675810,
			false,
		},
		{
			111111,
			100000,
			675810,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			result := meetsCriteria(tt.input, tt.rangeStart, tt.rangeEnd, true)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func Test_parttwo_possibilities(t *testing.T) {
	var startRange int64 = 134792
	var endRange int64 = 675810
	meetsCriteriaCount := 0

	for i := startRange; i <= endRange; i++ {
		if meetsCriteria(i, startRange, endRange, true) {
			meetsCriteriaCount++
		}
	}

	assert.Equal(t, 1319, meetsCriteriaCount)
}
