package lemin

import "fmt"

func AntHandler(routes [][]Room, numberofants int) {
	for antNumber := 1; antNumber <= numberofants; antNumber++ {
		for routeIndex := 0; routeIndex < len(routes); routeIndex++ {

			currentRoute := routes[routeIndex]
			for roomIndex := 1; roomIndex < len(currentRoute); roomIndex++ {
				fmt.Print("L", antNumber, "-", currentRoute[roomIndex].Name, " ")
			}
			fmt.Println()
		}
	}
}
