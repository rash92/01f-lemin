package lemin

import (
	"fmt"
)

type Ant struct {
	ID         int
	RouteIndex int
}

func IdentifyAnts(numberOfAnts int, antsPerRoute []int, routes [][]Room) (antslice []Ant) {
	antNumber := 0
	for antNumber < numberOfAnts {
		for routeIndex := 0; routeIndex < len(routes); routeIndex++ {
			if antsPerRoute[routeIndex] > 0 {
				antNumber++
				antslice = append(antslice, Ant{ID: antNumber, RouteIndex: routeIndex})
				antsPerRoute[routeIndex]--
			}
		}
	}
	// fmt.Println(antslice)
	return antslice
}

func AssignAnts(routes [][]Room, numberofants int) (antsPerRoute []int) {
	if len(routes) == 1 {
		return []int{numberofants}
	}
	if len(routes) == 0 {
		return antsPerRoute
	}

	for routeIndex := 0; routeIndex < len(routes); routeIndex++ {
		antsPerRoute = append(antsPerRoute, 0)
	}

	antNumber := 0

	for currentRouteIndex := 1; currentRouteIndex < len(routes); {
		for previousRouteIndex := 0; previousRouteIndex < currentRouteIndex; previousRouteIndex++ {

			antsPlusPathPrevious := antsPerRoute[previousRouteIndex] + len(routes[previousRouteIndex])
			antsPlusPathCurrent := antsPerRoute[currentRouteIndex] + len(routes[currentRouteIndex])

			if antsPlusPathPrevious <= antsPlusPathCurrent {

				antsPerRoute[previousRouteIndex]++
				previousRouteIndex = 0
				antNumber++
				// fmt.Println("after ant number", antNumber, "index ", currentRouteIndex, "has ants plus length: ", antsPlusPathCurrent, "index ", previousRouteIndex, "has ants plus length: ", antsPlusPathPrevious)
				// fmt.Println("after ant number", antNumber, "assigned ants are: ", antsPerRoute)
				// break
			}
			if antsPlusPathPrevious > antsPlusPathCurrent && previousRouteIndex == currentRouteIndex-1 {
				antsPerRoute[currentRouteIndex]++
				currentRouteIndex++
				antNumber++
				// fmt.Println("after ant number", antNumber, "index ", currentRouteIndex, "has ants plus length: ", antsPlusPathCurrent, "index ", previousRouteIndex, "has ants plus length: ", antsPlusPathPrevious)

				// fmt.Println("after ant number", antNumber, "assigned ants are: ", antsPerRoute)
				if currentRouteIndex == len(routes) {
					currentRouteIndex = 1
				}
				if antNumber == numberofants {
					return antsPerRoute
				}

				break
			}
			// fmt.Println("after ant number", antNumber, "assigned ants are: ", antsPerRoute)
			if antNumber == numberofants {
				return antsPerRoute
			}

		}
	}

	fmt.Println("assigned ants are: ", antsPerRoute)

	return antsPerRoute
}

func AreAntsFinished(antsPerRoute []int) bool {
	for _, ants := range antsPerRoute {
		if ants != 0 {
			return false
		}
	}
	return true
}

func PrintAnts(allAnts []Ant, routes [][]Room, numberOfAnts int) {
	antsRemaining := numberOfAnts
	for _, ant := range allAnts {
		for timeStep := 0; antsRemaining > 0; timeStep++ {
			for room := 1; room < len(routes[ant.RouteIndex]); room++ {

				fmt.Println("L", ant.ID, "- ", routes[ant.RouteIndex][room].Name)
				numberOfAnts--
			}
		}
		fmt.Println()
	}
}
