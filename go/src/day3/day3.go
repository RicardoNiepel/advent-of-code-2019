package day3

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type wire struct {
	x      int64
	y      int64
	length int64
}

func run(wiresPathString string) (minDistance string, fewestSteps string) {
	wireA := make(map[string]wire)
	wireB := make(map[string]wire)

	wires := strings.Split(wiresPathString, "\n")

	initialize(&wireA, wires[0])
	initialize(&wireB, wires[1])

	intersections := intersectKey(wireA, wireB)

	minDistanceInt := int64(math.MaxInt64)
	fewestStepsInt := int64(math.MaxInt64)

	for _, inter := range intersections {
		distance := int64(math.Abs(float64(wireA[inter].x)) + math.Abs(float64(wireA[inter].y)))
		if distance < minDistanceInt {
			minDistanceInt = distance
		}

		steps := wireA[inter].length + wireB[inter].length
		if steps < fewestStepsInt {
			fewestStepsInt = steps
		}
	}

	return fmt.Sprint(minDistanceInt), fmt.Sprint(fewestStepsInt)
}

func intersectKey(as, bs map[string]wire) []string {
	i := make([]string, 0)
	for a := range as {
		if a != "0,0" {

			if _, ok := bs[a]; ok {
				i = append(i, a)
			}
		}
	}
	return i
}

func initialize(wireType *map[string]wire, wirePathString string) {
	(*wireType)["0,0"] = wire{x: 0, y: 0, length: 0}
	wireEnd := (*wireType)["0,0"]

	wirePath := strings.Split(wirePathString, ",")

	for _, item := range wirePath {
		direction := string(item[0])
		way, _ := strconv.ParseInt(string(item[1:]), 10, 64)

		for i := int64(0); i < way; i++ {
			var newX, newY, newLength int64

			switch direction {
			case "R":
				newX = wireEnd.x + 1
				newY = wireEnd.y
			case "L":
				newX = wireEnd.x - 1
				newY = wireEnd.y
			case "U":
				newX = wireEnd.x
				newY = wireEnd.y + 1
			case "D":
				newX = wireEnd.x
				newY = wireEnd.y - 1
			}

			key := fmt.Sprintf("%v,%v", newX, newY)

			_, prs := (*wireType)[key]
			if prs {
				newLength = (*wireType)[key].length
			} else {
				newLength = wireEnd.length + 1
			}

			(*wireType)[key] = wire{x: newX, y: newY, length: newLength}
			wireEnd = wire{x: newX, y: newY, length: wireEnd.length + 1}
		}
	}
}
