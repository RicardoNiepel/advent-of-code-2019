package day8

import (
	"fmt"
	"strconv"
	"strings"
)

func decode(imageData string, wide int, tall int) string {
	layersAsString := getLayers(imageData, wide, tall)
	layers := createArrays(layersAsString, wide, tall)

	final := make([][]int, tall)
	for t := 0; t < tall; t++ {
		final[t] = make([]int, wide)
		for w := 0; w < wide; w++ {
			for l := 0; l < len(layers); l++ {
				if layers[l][t][w] == 2 {
					continue
				}
				final[t][w] = layers[l][t][w]
				break
			}
		}
	}

	finalString := ""
	for t := 0; t < tall; t++ {
		for w := 0; w < wide; w++ {
			finalString = fmt.Sprint(finalString, final[t][w])
		}
		if t < tall-1 {
			finalString = fmt.Sprint(finalString, "\n")
		}
	}

	return finalString
}

func createArrays(layersAsString []string, wide int, tall int) [][][]int {
	layers := make([][][]int, len(layersAsString))

	for l := 0; l < len(layers); l++ {
		layers[l] = make([][]int, tall)
		lines := strings.Split(layersAsString[l], "\n")

		for t := 0; t < tall; t++ {
			layers[l][t] = make([]int, wide)

			for w := 0; w < wide; w++ {
				number, _ := strconv.Atoi(string(lines[t][w]))
				layers[l][t][w] = number
				// layers[l] = fmt.Sprint(layers[l], string())
				// remove(&rawImageData, 0)
			}
			// layers[l] = fmt.Sprintln(layers[l])
		}
		// layers[l] = strings.TrimRight(layers[l], "\n")
	}

	return layers
}

func getChecksum(imageData string, wide int, tall int) int {
	layers := getLayers(imageData, wide, tall)

	lowestZeros := 10000
	var result int

	for _, l := range layers {
		if strings.Count(l, "0") < lowestZeros {
			lowestZeros = strings.Count(l, "0")
			result = strings.Count(l, "1") * strings.Count(l, "2")

		}
	}

	return result
}

func getLayers(imageData string, wide int, tall int) []string {
	rawImageData := []rune(imageData)
	layers := make([]string, len(imageData)/(wide*tall))

	for l := 0; l < len(layers); l++ {
		for t := 0; t < tall; t++ {
			for w := 0; w < wide; w++ {
				layers[l] = fmt.Sprint(layers[l], string(rawImageData[0]))
				remove(&rawImageData, 0)
			}
			layers[l] = fmt.Sprintln(layers[l])
		}
		layers[l] = strings.TrimRight(layers[l], "\n")
	}

	return layers
}

func remove(arr *[]rune, index int) {
	(*arr) = append((*arr)[:index], (*arr)[index+1:]...)
}
