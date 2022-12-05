package lemin

import "fmt"

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
	}
}
