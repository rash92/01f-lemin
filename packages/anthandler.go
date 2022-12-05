package lemin

import "fmt"

func RouteSorter(routes [][]string) (routesSortedByLength [][]string) {
	noOfRoutes := len(routes)
	for length := 0; len(routesSortedByLength) < noOfRoutes; length++ {
		// for routeIndex := 0; routeIndex < len(routes); routeIndex++ {

		// for length := 0; len(routesSortedByLength) < noOfRoutes; length++ {
		for routeIndex := 0; routeIndex < len(routes); routeIndex++ {
			currentRoute := routes[routeIndex]
			lengthOfRoute := len(currentRoute)
			if lengthOfRoute == length {
				fmt.Println(currentRoute)
				routesSortedByLength = append(routesSortedByLength, currentRoute)
			}
		}
	}
	return routesSortedByLength
}
