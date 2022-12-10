package lemin

import (
	"fmt"
)

type Ant struct {
	ID               int
	RouteIndex       int
	CurrentRoom      Room
	CurrentRoomIndex int
	Route            []Room
}

func AssignAntsPerRoute(antSlice []Ant, allRoutes [][]Room) (antsPerRoute [][]Ant) {
	numberOfRoutes := len(allRoutes)
	for routeIndex := 0; routeIndex < numberOfRoutes; routeIndex++ {
		antsForCurrentRoute := []Ant{}
		for _, ant := range antSlice {
			if ant.RouteIndex == routeIndex {
				antsForCurrentRoute = append(antsForCurrentRoute, ant)
			}
		}
		antsPerRoute = append(antsPerRoute, antsForCurrentRoute)
	}

	return antsPerRoute
}

func IdentifyAnts(numberOfAnts int, antsPerRoute []int, routes [][]Room) (antslice []Ant) {
	antNumber := 0
	for antNumber < numberOfAnts {
		for routeIndex := 0; routeIndex < len(routes); routeIndex++ {
			if antsPerRoute[routeIndex] > 0 {
				antNumber++
				antslice = append(antslice, Ant{ID: antNumber, RouteIndex: routeIndex, Route: routes[routeIndex]})
				antsPerRoute[routeIndex]--
			}
		}
	}
	// fmt.Println(antslice)
	return antslice
}

func AssignNumberOfAnts(routes [][]Room, numberofants int) (antsPerRoute []int) {
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

func FindMaxTimeSteps(antsPerRoute [][]Ant, allRoutes [][]Room) int {
	currentMax := 0
	for routeIndex, ants := range antsPerRoute {
		potentialMax := len(ants) + len(allRoutes[routeIndex])
		if potentialMax > currentMax {
			currentMax = potentialMax
		}
	}
	return currentMax
}

func MoveAnts(antsPerRoute [][]Ant) [][]Ant {
	for _, route := range antsPerRoute {
		for antIndex := len(route) - 1; antIndex > 0; antIndex-- {
			currentAnt := route[antIndex]
			previousAnt := route[antIndex-1]

			// if currently waiting and previous ant not waiting i.e. is in a room, go to first room
			if currentAnt.CurrentRoom.Name == "" && previousAnt.CurrentRoom.Name != "" {
				currentAnt.CurrentRoomIndex = 0
				currentAnt.CurrentRoom = currentAnt.Route[currentAnt.CurrentRoomIndex]
			}
			// if not currently waiting i.e. already in a room, go to the next room in the route
			if currentAnt.CurrentRoom.Name != "" {
				currentAnt.CurrentRoomIndex++
				currentAnt.CurrentRoom = currentAnt.Route[currentAnt.CurrentRoomIndex]
			}
		}
	}
	return antsPerRoute
}

func PrintAnts(antsPerRoute [][]Ant, allRoutes [][]Room) {
}
