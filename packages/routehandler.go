package lemin

import (
	"fmt"
	"reflect"
)

func IsRoomContainedInRoute(element Room, slice []Room) bool {
	for _, potentialelement := range slice {
		if element.Name == potentialelement.Name {
			return true
		}
	}
	return false
}

func IsRouteConatinedInRoutes(route []Room, routes [][]Room) bool {
	for _, potentialroute := range routes {
		if reflect.DeepEqual(route, potentialroute) {
			return true
		}
	}
	return false
}

// takes in a list of routes and returns the same list sorted by length from shortest to longest
func RouteSorter(routes [][]Room) (routesSortedByLength [][]Room) {
	noOfRoutes := len(routes)
	for length := 0; len(routesSortedByLength) < noOfRoutes; length++ {
		for routeIndex := 0; routeIndex < len(routes); routeIndex++ {
			currentRoute := routes[routeIndex]
			lengthOfRoute := len(currentRoute)
			if lengthOfRoute == length {
				routesSortedByLength = append(routesSortedByLength, currentRoute)
			}
		}
	}
	return routesSortedByLength
}

func RemoveIncomplete(existingRoutes [][]Room, endingRoom Room) (outputRoutes [][]Room) {
	for _, route := range existingRoutes {
		if route[len(route)-1].Name == endingRoom.Name {
			outputRoutes = append(outputRoutes, route)
		}
	}
	return outputRoutes
}

func FindLengthOfRouteForAnts(setOfRoutes [][]Room, numberOfAnts int) int {
	routeLength := 0

	antsForRoute1 := AssignNumberOfAnts(setOfRoutes, numberOfAnts)

	for routeIndex, route := range setOfRoutes {
		potentialLength := len(route) + antsForRoute1[routeIndex]
		if potentialLength > routeIndex {
			routeLength = potentialLength
		}
	}
	return routeLength
}

func IsSetOfRoutesShorter(routes1 [][]Room, routes2 [][]Room, numberOfAnts int) bool {
	return FindLengthOfRouteForAnts(routes1, numberOfAnts) < FindLengthOfRouteForAnts(routes2, numberOfAnts)
}

// takes a list of routes and returns a checks if any go through the same middle rooms, true if they don't
func IsSetOfRoutesIndependent(existingRoutes [][]Room) bool {
	for route1Index := 0; route1Index < len(existingRoutes); route1Index++ {
		route1 := existingRoutes[route1Index]
		for route2Index := 0; route2Index < len(existingRoutes); route2Index++ {
			route2 := existingRoutes[route2Index]
			for roomIndex := 1; roomIndex < len(route1)-1; roomIndex++ {
				room := route1[roomIndex]
				if IsRoomContainedInRoute(room, route2) && route2Index != route1Index {
					return false
				}
			}
		}
	}

	return true
}

func FindInitialValidSetOfPaths(existingRoutes [][]Room) [][]Room {
	outputRoutes := RouteSorter(existingRoutes)
	// outputRoutes = append(outputRoutes[2:])
	for route1Index := 0; route1Index < len(outputRoutes); route1Index++ {
		route1 := outputRoutes[route1Index]
		for route2Index := 0; route2Index < len(outputRoutes); route2Index++ {
			route2 := outputRoutes[route2Index]
			for roomIndex := 1; roomIndex < len(route1)-1; roomIndex++ {
				room := route1[roomIndex]
				if IsRoomContainedInRoute(room, route2) && route2Index != route1Index && len(route1) <= len(route2) {
					outputRoutes = append(outputRoutes[:route2Index], outputRoutes[route2Index+1:]...)
					route2Index--
					if route2Index < route1Index {
						route1Index--
					}
					break
				}
			}
		}
	}
	return outputRoutes
}

func RemoveDuplicates(allRoutesIncludingDuplicates [][]Room, numberOfAnts int) [][]Room {
	// fmt.Println("current length of attempted set of routes is: ", len(allRoutesIncludingDuplicates))
	// start with longest possible times all ants going down single route

	initialGuess := FindInitialValidSetOfPaths(allRoutesIncludingDuplicates)
	initialLength := FindLengthOfRouteForAnts(initialGuess, numberOfAnts)
	var outputGuess [][]Room
	outputGuess = initialGuess
	shortestLength := initialLength
	allRoutesIncludingDuplicates = RouteSorter(allRoutesIncludingDuplicates)

	fmt.Println("initial guess, length, outputguess are: ", len(initialGuess), initialLength, len(outputGuess))

	if IsSetOfRoutesIndependent(allRoutesIncludingDuplicates) && FindLengthOfRouteForAnts(allRoutesIncludingDuplicates, numberOfAnts) < initialLength {
		fmt.Println("is current attempted set valid")
		return allRoutesIncludingDuplicates
	}

	for routeIndex := 0; routeIndex < len(allRoutesIncludingDuplicates); routeIndex++ {
		potentialSetOfRoutes := append(allRoutesIncludingDuplicates[routeIndex:])
		fmt.Println("potential set of routes has length: ", len(potentialSetOfRoutes), "current routeIndex is: ", routeIndex, "current shortest length is: ", shortestLength)
		// fmt.Println("potential set of routes has first element: ", potentialSetOfRoutes[0])
		if IsSetOfRoutesIndependent(potentialSetOfRoutes) {
			lengthOfCurrentSetOfRoutes := FindLengthOfRouteForAnts(potentialSetOfRoutes, numberOfAnts)
			if lengthOfCurrentSetOfRoutes < shortestLength {
				shortestLength = lengthOfCurrentSetOfRoutes
				fmt.Println("here before copy")
				outputGuess = potentialSetOfRoutes

				// fmt.Println("current length of running best set of routes is before the else: ", len(RunningBestSetOfRoutes))
			}
		} else {
			potentialSetOfRoutes := FindInitialValidSetOfPaths(potentialSetOfRoutes)
			fmt.Println("potential set of routes has length: ", len(potentialSetOfRoutes), "current routeIndex is: ", routeIndex, "current shortest length is: ", shortestLength)
			lengthOfCurrentSetOfRoutes := FindLengthOfRouteForAnts(potentialSetOfRoutes, numberOfAnts)
			fmt.Println("inside the else - length of current potential routes is: ", lengthOfCurrentSetOfRoutes)
			if lengthOfCurrentSetOfRoutes < shortestLength {
				shortestLength = lengthOfCurrentSetOfRoutes

				outputGuess = potentialSetOfRoutes
			}
		}
	}
	// fmt.Println("current length of running best set of routes is: ", len(RunningBestSetOfRoutes))
	if len(outputGuess) == 0 {
		fmt.Println("here")
		return [][]Room{}
	}
	return outputGuess
}

// finds a route that hasn't already been found and put in existingRoutes if possible, or returns an empty route
func FindRoute(startingRoom Room, endingRoom Room, allRooms []Room, existingRoute []Room, existingRoutes *[][]Room) (routeNames []Room) {
	existingRoute = append(existingRoute, startingRoom)

	if startingRoom.Name == endingRoom.Name {
		// fmt.Println("new route ending is: ", existingRoute[len(existingRoute)-1].Name)
		if existingRoute[len(existingRoute)-1].Name != endingRoom.Name {
			return []Room{}
		}
		return existingRoute
	}

	for i := 0; i < len(startingRoom.LinksAsPointers); i++ {
		currentLinkedRoom := startingRoom.LinksAsPointers[i]

		if !IsRoomContainedInRoute((*currentLinkedRoom), existingRoute) && len(FindRoute(*currentLinkedRoom, endingRoom, allRooms, existingRoute, existingRoutes)) != 0 {
			potentialRoute := FindRoute(*currentLinkedRoom, endingRoom, allRooms, existingRoute, existingRoutes)
			if !IsRouteConatinedInRoutes(potentialRoute, *existingRoutes) && len(potentialRoute) != 0 {

				*existingRoutes = append(*existingRoutes, potentialRoute)
				existingRoute = potentialRoute
				if existingRoute[len(existingRoute)-1].Name != endingRoom.Name {
					return []Room{}
				}

				return existingRoute
			}
		}
	}

	if existingRoute[len(existingRoute)-1].Name != endingRoom.Name {
		return []Room{}
	}

	return []Room{}
}

// finds all possible valid routes, and sorts them from shortest to longest
func FindAllRoutes(startingRoom Room, endingRoom Room, allRooms []Room, existingRoutes [][]Room, numberOfAnts int) (allRoutesNames [][]Room) {
	FindRoute(startingRoom, endingRoom, allRooms, []Room{}, &existingRoutes)

	allRoutesNames = RemoveDuplicates((RemoveIncomplete(existingRoutes, endingRoom)), numberOfAnts)

	return allRoutesNames
}
