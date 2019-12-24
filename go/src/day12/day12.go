package day12

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Moon struct {
	posx int
	posy int
	posz int
	velx int
	vely int
	velz int
}

func findPreviousPointInTime(scan string) int64 {
	lines := strings.Split(scan, "\n")

	moons := make([]*Moon, 0)
	for _, line := range lines {
		values := strings.Split(line, ",")

		moon := Moon{}
		moons = append(moons, &moon)

		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		z, _ := strconv.Atoi(values[2])
		moon.posx = x
		moon.posy = y
		moon.posz = z
		moon.velx = 0
		moon.vely = 0
		moon.velz = 0
	}

	moon0posx := moons[0].posx
	moon0posy := moons[0].posy
	moon0posz := moons[0].posz
	moon0velx := moons[0].velx
	moon0vely := moons[0].vely
	moon0velz := moons[0].velz
	moon1posx := moons[1].posx
	moon1posy := moons[1].posy
	moon1posz := moons[1].posz
	moon1velx := moons[1].velx
	moon1vely := moons[1].vely
	moon1velz := moons[1].velz
	moon2posx := moons[2].posx
	moon2posy := moons[2].posy
	moon2posz := moons[2].posz
	moon2velx := moons[2].velx
	moon2vely := moons[2].vely
	moon2velz := moons[2].velz
	moon3posx := moons[3].posx
	moon3posy := moons[3].posy
	moon3posz := moons[3].posz
	moon3velx := moons[3].velx
	moon3vely := moons[3].vely
	moon3velz := moons[3].velz

	var resultX, resultY, resultZ int64

	for step := int64(0); step < math.MaxInt64; step++ {
		applyGravityX(moons)
		applyVelocityX(moons)

		if moons[0].posx == moon0posx && moons[0].velx == moon0velx && moons[1].posx == moon1posx && moons[1].velx == moon1velx && moons[2].posx == moon2posx && moons[2].velx == moon2velx && moons[3].posx == moon3posx && moons[3].velx == moon3velx {
			resultX = step + 1
			break
		}
	}

	for step := int64(0); step < math.MaxInt64; step++ {
		applyGravityY(moons)
		applyVelocityY(moons)

		if moons[0].posy == moon0posy && moons[0].vely == moon0vely && moons[1].posy == moon1posy && moons[1].vely == moon1vely && moons[2].posy == moon2posy && moons[2].vely == moon2vely && moons[3].posy == moon3posy && moons[3].vely == moon3vely {
			resultY = step + 1
			break
		}
	}

	for step := int64(0); step < math.MaxInt64; step++ {
		applyGravityZ(moons)
		applyVelocityZ(moons)

		if moons[0].posz == moon0posz && moons[0].velz == moon0velz && moons[1].posz == moon1posz && moons[1].velz == moon1velz && moons[2].posz == moon2posz && moons[2].velz == moon2velz && moons[3].posz == moon3posz && moons[3].velz == moon3velz {
			resultZ = step + 1
			break
		}
	}

	return lcm(resultX, resultY, resultZ)
}

func lcm(a, b, c int64) int64 {
	lcm := a * (b / gcd(a, b))
	return c * (lcm / gcd(lcm, c))
}

func gcd(m, n int64) int64 {
	if n == 0 {
		return m
	}
	return gcd(n, m%n)
}

func run(scan string, steps int) int {
	lines := strings.Split(scan, "\n")

	moons := make([]*Moon, 0)
	for _, line := range lines {
		values := strings.Split(line, ",")

		moon := Moon{}
		moons = append(moons, &moon)

		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		z, _ := strconv.Atoi(values[2])
		moon.posx = x
		moon.posy = y
		moon.posz = z
		moon.velx = 0
		moon.vely = 0
		moon.velz = 0
	}

	fmt.Printf("After 0 steps:\n")
	printMoons(moons)
	fmt.Println()

	for step := 0; step < steps; step++ {
		applyGravityX(moons)
		applyGravityY(moons)
		applyGravityZ(moons)
		applyVelocityX(moons)
		applyVelocityY(moons)
		applyVelocityZ(moons)
		fmt.Printf("After %v steps:\n", step+1)
		printMoons(moons)
		fmt.Println()
	}

	return getTotalEnergy(moons)
}

func getTotalEnergy(moons []*Moon) int {
	sum := 0
	for i := 0; i < len(moons); i++ {
		pot := math.Abs(float64(moons[i].posx)) + math.Abs(float64(moons[i].posy)) + math.Abs(float64(moons[i].posz))
		kin := math.Abs(float64(moons[i].velx)) + math.Abs(float64(moons[i].vely)) + math.Abs(float64(moons[i].velz))
		sumMoon := int(pot) * int(kin)
		sum += sumMoon
	}
	return sum
}

func printMoons(moons []*Moon) {
	for i := 0; i < len(moons); i++ {
		fmt.Printf("pos=<x=%v, y=%v, z=%v>, vel=<x=%v, y=%v, z=%v>\n", moons[i].posx, moons[i].posy, moons[i].posz, moons[i].velx, moons[i].vely, moons[i].velz)
	}
}

func applyVelocityX(moons []*Moon) {
	for i := 0; i < len(moons); i++ {
		moons[i].posx += moons[i].velx
	}
}

func applyVelocityY(moons []*Moon) {
	for i := 0; i < len(moons); i++ {
		moons[i].posy += moons[i].vely
	}
}
func applyVelocityZ(moons []*Moon) {
	for i := 0; i < len(moons); i++ {
		moons[i].posz += moons[i].velz
	}
}

func applyGravityX(moons []*Moon) {
	for i := 0; i < len(moons); i++ {
		for k := 0; k < len(moons); k++ {
			if i == k {
				continue
			}
			if moons[i].posx < moons[k].posx {
				moons[i].velx++
				moons[k].velx--
			}
		}
	}
}

func applyGravityY(moons []*Moon) {
	for i := 0; i < len(moons); i++ {
		for k := 0; k < len(moons); k++ {
			if i == k {
				continue
			}
			if moons[i].posy < moons[k].posy {
				moons[i].vely++
				moons[k].vely--
			}
		}
	}
}

func applyGravityZ(moons []*Moon) {
	for i := 0; i < len(moons); i++ {
		for k := 0; k < len(moons); k++ {
			if i == k {
				continue
			}
			if moons[i].posz < moons[k].posz {
				moons[i].velz++
				moons[k].velz--
			}
		}
	}
}
