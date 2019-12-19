package day4

import (
	"fmt"
)

func meetsCriteria(input, rangeStart, rangeEnd int64, adjacentAreNotPartOfLargerGroup bool) bool {
	if input < 100000 {
		return false
	}

	if input < rangeStart || input > rangeEnd {
		return false
	}

	inputAsString := fmt.Sprint(input)

	if inputAsString[0] != inputAsString[1] &&
		inputAsString[1] != inputAsString[2] &&
		inputAsString[2] != inputAsString[3] &&
		inputAsString[3] != inputAsString[4] &&
		inputAsString[4] != inputAsString[5] {
		return false
	}

	if adjacentAreNotPartOfLargerGroup {
		if !(areAdjacentsAreNotPartOfLargerGroup(&inputAsString, 0) ||
			areAdjacentsAreNotPartOfLargerGroup(&inputAsString, 1) ||
			areAdjacentsAreNotPartOfLargerGroup(&inputAsString, 2) ||
			areAdjacentsAreNotPartOfLargerGroup(&inputAsString, 3) ||
			(inputAsString[3] != inputAsString[4] && inputAsString[4] == inputAsString[5])) {
			return false
		}
	}

	if !(inputAsString[0] <= inputAsString[1] &&
		inputAsString[1] <= inputAsString[2] &&
		inputAsString[2] <= inputAsString[3] &&
		inputAsString[3] <= inputAsString[4] &&
		inputAsString[4] <= inputAsString[5]) {
		return false
	}

	return true
}

func areAdjacentsAreNotPartOfLargerGroup(input *string, index int) bool {
	if index+2 < len(*input) {
		if index > 0 {
			return (*input)[index-1] != (*input)[index] && (*input)[index] == (*input)[index+1] && (*input)[index+1] != (*input)[index+2]
		} else {
			return (*input)[index] == (*input)[index+1] && (*input)[index+1] != (*input)[index+2]
		}
	}
	panic("index too large")
}
