package day10

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"unicode"
)

type asteroid struct {
	x int
	y int
}

func blastAsteroids(spaceMap *[][]string, station asteroid, rounds int) (last asteroid) {
	removeShootNumbers(spaceMap)

	reachableAsteroids := getDetectedAsteroidsInternal(spaceMap, station.x, station.y)
	sort.Slice(reachableAsteroids, func(i, j int) bool {
		return getRotatingLaserAngleFromCoordAngle(getAngle(station, reachableAsteroids[i])) < getRotatingLaserAngleFromCoordAngle(getAngle(station, reachableAsteroids[j]))
	})

	for i := 0; i < rounds && i < len(reachableAsteroids); i++ {
		last = reachableAsteroids[i]
		(*spaceMap)[last.y][last.x] = fmt.Sprint(i + 1)
	}

	return
}

func removeShootNumbers(spaceMap *[][]string) {
	for _, line := range *spaceMap {
		for x, pos := range line {
			if unicode.IsNumber(rune(pos[0])) {
				line[x] = "."
			}
		}
	}
}

func getRotatingLaserAngleFromCoordAngle(angle float64) float64 {
	var tmp float64
	if angle > 270 {
		tmp = angle - 270
	} else {
		tmp = angle + 90
	}

	if tmp == 360 {
		return 0
	}

	return tmp

	// return 360 - tmp
}

func getBestStationAsteroid(rawMap string) (bestStation asteroid, detectedAsteroids []asteroid, spaceMap [][]string) {
	spaceMap = getMap(rawMap)
	possibleStations := getAllAsteroids(&spaceMap)
	for _, possibleStation := range possibleStations {
		spaceMapCopy := spaceMap
		detectedAsteroidsTmp := getDetectedAsteroidsInternal(&spaceMapCopy, possibleStation.x, possibleStation.y)

		spaceMap[possibleStation.y][possibleStation.x] = fmt.Sprint(len(detectedAsteroidsTmp))

		if len(detectedAsteroidsTmp) > len(detectedAsteroids) {
			detectedAsteroids = detectedAsteroidsTmp
			bestStation = possibleStation
		}
	}
	return
}

func getDetectedAsteroids(rawMap, monitoringStationLabel string) (detectedAsteroids []asteroid, spaceMap [][]string) {
	spaceMap = getMap(rawMap)
	x, y := findPosition(&spaceMap, monitoringStationLabel)
	detectedAsteroids = getDetectedAsteroidsInternal(&spaceMap, x, y)
	return
}

func getDetectedAsteroidsInternal(spaceMap *[][]string, x, y int) []asteroid {
	station := asteroid{x: x, y: y}

	asteroids := getAllAsteroids(spaceMap)
	sort.Slice(asteroids, func(i, j int) bool {
		return getDistance(station, asteroids[i]) < getDistance(station, asteroids[j])
	})

	detectedAsteroids := make([]asteroid, 0)
	for _, asteroid := range asteroids[1:] {
		if hidingAsteroid := getHiding(&station, &detectedAsteroids, &asteroid); hidingAsteroid != nil {
			symbol := (*spaceMap)[hidingAsteroid.y][hidingAsteroid.x]
			if unicode.IsLetter(rune(symbol[0])) {
				(*spaceMap)[asteroid.y][asteroid.x] = strings.ToLower(symbol)
			}
		} else {
			detectedAsteroids = append(detectedAsteroids, asteroid)
		}
	}

	return detectedAsteroids
}

func getHiding(station *asteroid, detectedAsteroids *[]asteroid, asteroid *asteroid) *asteroid {
	for _, detectedAsteroid := range *detectedAsteroids {
		if isHiding(station, &detectedAsteroid, asteroid) {
			return &detectedAsteroid
		}
	}

	return nil
}

func isHiding(station *asteroid, detectedAsteroid *asteroid, asteroid *asteroid) bool {
	// Using polar coordinate system
	angle1 := getAngle(*station, *detectedAsteroid)
	angle2 := getAngle(*station, *asteroid)
	return angle1 == angle2
}

func getDistance(asteroid1, asteroid2 asteroid) float64 {
	x1, x2 := float64(asteroid1.x), float64(asteroid2.x)
	y1, y2 := float64(asteroid1.y), float64(asteroid2.y)
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func getAngle(asteroid1, asteroid2 asteroid) float64 {
	x1, x2 := float64(asteroid1.x), float64(asteroid2.x)
	y1, y2 := float64(asteroid1.y), float64(asteroid2.y)
	radian := math.Atan2(y2-y1, x2-x1)
	degree := 180 / math.Pi * radian
	if degree < 0 {
		degree += 360
	}
	return degree
}

func getAllAsteroids(spaceMap *[][]string) []asteroid {
	asteroids := make([]asteroid, 0)

	for y, line := range *spaceMap {
		for x, position := range line {
			if position != "." {
				asteroids = append(asteroids, asteroid{x: x, y: y})
			}
		}
	}

	return asteroids
}

func findPosition(spaceMap *[][]string, element string) (x, y int) {
	for y, line := range *spaceMap {
		for x, position := range line {
			if element == position {
				return x, y
			}
		}
	}

	return -1, -1
}

func getMap(rawMap string) [][]string {
	lines := strings.Split(rawMap, "\n")
	spaceMap := make([][]string, len(lines))

	for y, line := range lines {
		positions := []rune(line)
		spaceMap[y] = make([]string, len(positions))
		for x, position := range positions {
			spaceMap[y][x] = string(position)
		}
	}
	return spaceMap
}

func printMap(spaceMap [][]string) string {
	var sb strings.Builder
	for y, line := range spaceMap {
		for _, position := range line {
			sb.WriteString(position)
		}
		if y < len(spaceMap)-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
