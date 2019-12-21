package day6

import (
	"strings"
)

type object struct {
	name     string
	parent   *object
	children []*object
}

func getTotalOrbitCount(orbitMapString string) int64 {
	orbits := initialize(orbitMapString)
	return sumOrbits(orbits["COM"], 0)
}

func getMinimumNumberOfOrbitalTransfers(orbitMapString string) int64 {
	orbits := initialize(orbitMapString)

	from := orbits["YOU"].parent
	to := orbits["SAN"].parent

	var alreadyTested []*object

	found, length := findLength(from, to, &alreadyTested)
	if !found {
		panic("not found")
	}
	return length
}

func findLength(from *object, to *object, alreadyTested *[]*object) (found bool, length int64) {
	if from.parent == to {
		return true, 1
	}

	if from.parent != nil && !contains(alreadyTested, from.parent) {
		*alreadyTested = append(*alreadyTested, from.parent)
		parFound, parLength := findLength(from.parent, to, alreadyTested)
		if parFound {
			return true, parLength + 1
		}
	}

	if from.children != nil {
		for _, child := range from.children {
			if child == to {
				return true, 1
			}
			if contains(alreadyTested, child) {
				continue
			}

			*alreadyTested = append(*alreadyTested, child)
			childFound, childLength := findLength(child, to, alreadyTested)
			if childFound {
				return true, childLength + 1
			}
		}
	}

	return false, -1
}

func contains(arr *[]*object, elem *object) bool {
	for _, n := range *arr {
		if elem == n {
			return true
		}
	}
	return false
}

func initialize(orbitMapString string) map[string]*object {
	orbitDefs := strings.Split(orbitMapString, "\n")

	orbits := make(map[string]*object)

	for _, orbitDefTmp := range orbitDefs {
		orbitDef := strings.Split(orbitDefTmp, ")")
		parent := orbitDef[0]
		child := orbitDef[1]

		orbitParent, prs := orbits[parent]
		if !prs {
			orbitParent = &object{name: parent}
			orbits[parent] = orbitParent
		}

		orbitChild, prs := orbits[child]
		if !prs {
			orbitChild = &object{name: child}
			orbits[child] = orbitChild
		}

		orbitParent.children = append(orbitParent.children, orbitChild)
		orbitChild.parent = orbitParent
	}

	return orbits
}

func sumOrbits(parent *object, hierarchy int64) (totalOrbits int64) {
	totalOrbits += hierarchy
	hierarchy++
	if parent.children != nil {
		for _, child := range parent.children {
			totalOrbits += sumOrbits(child, hierarchy)

		}
	}
	return
}
