package day1

import "math"

func run(moduleMass float64) float64 {
	result := math.Floor(moduleMass/3) - 2
	return result
}

func runRecursive(moduleMass float64) float64 {
	result := run(moduleMass)

	tmp := result
	for {
		tmp = run(tmp)
		if tmp <= 0 {
			break
		} else {
			result += tmp
		}
	}

	return result
}
