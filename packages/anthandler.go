package lemin

import "fmt"

func AssignAnts(routes [][]Room, numberofants int) (antsPerRoute []int) {
	for routeIndex := 0; routeIndex < len(routes); routeIndex++ {
		antsPerRoute = append(antsPerRoute, 0)
	}

	for antNumber := 0; antNumber < numberofants; antNumber++ {
		for routeIndex := 1; routeIndex < len(routes); routeIndex++ {
			currentRoute := routes[routeIndex]
			antsPlusPath := antsPerRoute[routeIndex-1] + len(routes[routeIndex-1])
			antsPlusPathCurrent := antsPerRoute[routeIndex] + len(currentRoute)
			if antsPlusPathCurrent <= antsPlusPath {
				antsPerRoute[routeIndex]++
				antsPlusPath = antsPlusPathCurrent
				break
			} else {
				antsPerRoute[routeIndex-1]++
				antsPlusPath++

			}

		}
		// fmt.Println("after ant number", antNumber, "assigned ants are: ", antsPerRoute)
	}
	fmt.Println("assigned ants are: ", antsPerRoute)
	if len(routes) > 1 {
		return antsPerRoute
	}
	return []int{numberofants}
}

func AntHandler(routes [][]Room, numberofants int) {
	fmt.Println("number of paths is: ", len(routes), "number of ants is: ", numberofants)

	for routeIndex := 0; routeIndex < len(routes); routeIndex++ {
		currentRoute := routes[routeIndex]
		for roomIndex := 1; roomIndex < len(currentRoute); roomIndex++ {
			for antNumber := 0; antNumber <= numberofants; antNumber++ {
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
