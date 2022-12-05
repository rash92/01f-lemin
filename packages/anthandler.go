package lemin

import "fmt"

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
				fmt.Println("after ant number", antNumber, "index ", currentRouteIndex, "has ants plus length: ", antsPlusPathCurrent, "index ", previousRouteIndex, "has ants plus length: ", antsPlusPathPrevious)
				fmt.Println("after ant number", antNumber, "assigned ants are: ", antsPerRoute)
				// break
			}
			if antsPlusPathPrevious > antsPlusPathCurrent && previousRouteIndex == currentRouteIndex-1 {
				antsPerRoute[currentRouteIndex]++
				currentRouteIndex++
				antNumber++
				fmt.Println("after ant number", antNumber, "index ", currentRouteIndex, "has ants plus length: ", antsPlusPathCurrent, "index ", previousRouteIndex, "has ants plus length: ", antsPlusPathPrevious)

				fmt.Println("after ant number", antNumber, "assigned ants are: ", antsPerRoute)
				if currentRouteIndex == len(routes) {
					currentRouteIndex = 1
				}
				if antNumber == numberofants {
					return antsPerRoute
				}

				break
			}
			fmt.Println("after ant number", antNumber, "assigned ants are: ", antsPerRoute)
			if antNumber == numberofants {
				return antsPerRoute
			}

		}
	}

	fmt.Println("assigned ants are: ", antsPerRoute)

	return antsPerRoute
}

func AntHandler(routes [][]Room, numberofants int) {
	fmt.Println("number of paths is: ", len(routes), "number of ants is: ", numberofants)

	for routeIndex := 0; routeIndex < len(routes); routeIndex++ {
		currentRoute := routes[routeIndex]
		for roomIndex := 1; roomIndex < len(currentRoute); roomIndex++ {
			for antNumber := 0; antNumber < numberofants; antNumber++ {
				if antNumber%len(routes) == routeIndex {
					fmt.Print("L", antNumber, "-", currentRoute[roomIndex].Name, " ")
				}
			}
			fmt.Println()
		}
		// fmt.Println()
	}
	// fmt.Println()
}
