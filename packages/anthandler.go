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
	newAntsPerRoute := make([][]Ant, len(antsPerRoute))
	copy(newAntsPerRoute, antsPerRoute)
	// fmt.Println("first ant per route before are: ", newAntsPerRoute[0][0].CurrentRoom.Name)
	for routeIndex, route := range antsPerRoute {
		for antIndex := len(route) - 1; antIndex >= 0; antIndex-- {

			currentAnt := route[antIndex]
			// fmt.Println("current ant location is: ", currentAnt.CurrentRoom.Name, "for ant", currentAnt.ID)
			if antIndex > 0 {
				previousAnt := route[antIndex-1]

				// if currently waiting and previous ant not waiting i.e. is in a room, go to first room
				if currentAnt.CurrentRoom.Name == "" && previousAnt.CurrentRoom.Name != "" {
					currentAnt.CurrentRoomIndex = 0
					currentAnt.CurrentRoom = currentAnt.Route[currentAnt.CurrentRoomIndex]
				}
			}
			// if not currently waiting i.e. already in a room, go to the next room in the route
			if currentAnt.CurrentRoom.Name != "" {
				if currentAnt.CurrentRoomIndex == len(currentAnt.Route)-1 {
					currentAnt.CurrentRoom.Name = ""
					break
				}
				currentAnt.CurrentRoomIndex = currentAnt.CurrentRoomIndex + 1

				currentAnt.CurrentRoom = currentAnt.Route[currentAnt.CurrentRoomIndex]
				// fmt.Println("current room index is: ", currentAnt.CurrentRoomIndex, "current room name is: ", currentAnt.CurrentRoom.Name)
				// fmt.Println("current ant after is: ", currentAnt)
			}
			// fmt.Println("current ant is: ", currentAnt.ID)
			// fmt.Println("current ant location after is: ", currentAnt.CurrentRoom.Name, "for ant", currentAnt.ID)
			newAntsPerRoute[routeIndex][antIndex] = currentAnt
		}
	}
	// fmt.Println("first ant per route after are: ", newAntsPerRoute[0][0].CurrentRoom.Name)
	return newAntsPerRoute
}

func PrintAnts(antsPerRoute [][]Ant, allRoutes [][]Room, numberOfAnts int) {
	// initial assign first ant in each route to first room
	fmt.Println("final printing is: ")
	for _, route := range antsPerRoute {
		route[0].CurrentRoom = route[0].Route[0]
		fmt.Print(" L", route[0].ID, "-", route[0].CurrentRoom.Name)
	}
	fmt.Println()

	timeSteps := FindMaxTimeSteps(antsPerRoute, allRoutes)

	for timeStep := 0; timeStep < timeSteps; timeStep++ {

		antsPerRoute = MoveAnts(antsPerRoute)

		for id := 1; id <= numberOfAnts; id++ {
			for _, route := range antsPerRoute {
				for _, ant := range route {
					if ant.ID == id && ant.CurrentRoom.Name != "" {
						fmt.Print(" L", ant.ID, "-", ant.CurrentRoom.Name)
					}
				}
			}
		}
		fmt.Println()
	}
}
