package day10

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partone_pre_small(t *testing.T) {
	input := `#...Z..*..
...A......
...B..*...
.EDCG....*
..F.*.*...
H....*....
*.***.*.**
.......*..
*...*...*.
...*..*..*`

	expected := `#...Z..z..
...A......
...B..a...
.EDCG....a
..F.c.b...
H....c....
h.efd.c.gb
.......c..
h...f...c.
...e..d..c`

	_, spaceMap := getDetectedAsteroids(input, "#")
	output := printMap(spaceMap)

	assert.Equal(t, expected, output)
}

func Test_partone_small(t *testing.T) {
	tests := []struct {
		input                    string
		expectedStationsOverview string
		expectedStationX         int
		expectedStationY         int
		expectedDetected         int
	}{
		{
			`.#..#
.....
#####
....#
...##`,
			`.7..7
.....
67775
....7
...87`,
			3,
			4,
			8,
		},
		{
			`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`,
			"",
			5,
			8,
			33,
		},
		{
			`#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`,
			"",
			1,
			2,
			35,
		},
		{
			`.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`,
			"",
			6,
			3,
			41,
		},
		{
			`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`,
			"",
			11,
			13,
			210,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.expectedDetected), func(t *testing.T) {

			bestStation, detectedAsteroids, spaceMap := getBestStationAsteroid(tt.input)

			if tt.expectedStationsOverview != "" {
				assert.Equal(t, tt.expectedStationsOverview, printMap(spaceMap))
				fmt.Println(printMap(spaceMap))
				fmt.Println()
			}

			assert.Equal(t, tt.expectedStationX, bestStation.x)
			assert.Equal(t, tt.expectedStationY, bestStation.y)
			assert.Equal(t, tt.expectedDetected, len(detectedAsteroids))
		})
	}
}

func Test_partone_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	bestStation, detectedAsteroids, _ := getBestStationAsteroid(input)

	assert.Equal(t, 14, bestStation.x)
	assert.Equal(t, 17, bestStation.y)
	assert.Equal(t, 260, len(detectedAsteroids))
}

func Test_parttwo_small(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....X...###..
..#.#.....#....##`,
			`.#....###24...#..
##...##.13#67..9#
##...#...5.8####.
..#.....X...###..
..#.#.....#....##`,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			spaceMap := getMap(tt.input)
			x, y := findPosition(&spaceMap, "X")
			station := asteroid{x: x, y: y}

			_ = blastAsteroids(&spaceMap, station, 9)
			assert.Equal(t, tt.expected, printMap(spaceMap))
		})
	}
}

func Test_parttwo_main(t *testing.T) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)
	spaceMap := getMap(input)
	station := asteroid{x: 14, y: 17}

	last := blastAsteroids(&spaceMap, station, 200)
	assert.Equal(t, 6, last.x)
	assert.Equal(t, 8, last.y)
}

func Test_getRotatingLaserAngleFromCoordAngle(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{90, 180},
		{120, 210},
		{180, 270},
		{270, 0},
		{0, 90},
		{30, 120},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			output := getRotatingLaserAngleFromCoordAngle(tt.input)
			assert.Equal(t, tt.expected, output)
		})
	}
}
