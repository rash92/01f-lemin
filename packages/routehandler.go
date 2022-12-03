package lemin

import (
	"reflect"
)

func IsRoomContainedInRoute(element string, slice []string) bool {
	for _, potentialelement := range slice {
		if element == potentialelement {
			return true
		}
	}
	return false
}

func IsRouteConatinedInRoutes(route []string, routes [][]string) bool {
	for _, potentialroute := range routes {
		if reflect.DeepEqual(route, potentialroute) {
			return true
		}
	}
	return false
}

func RemoveDuplicates(existingRoutes [][]string) [][]string {
	for route1Index := 0; route1Index < len(existingRoutes); route1Index++ {
		route1 := existingRoutes[route1Index]
		for route2Index := 0; route2Index < len(existingRoutes); route2Index++ {
			route2 := existingRoutes[route2Index]
			for roomIndex := 1; roomIndex < len(route1)-1; roomIndex++ {
				room := route1[roomIndex]
				if IsRoomContainedInRoute(room, route2) && route2Index != route1Index && len(route1) <= len(route2) {
					existingRoutes = append(existingRoutes[:route2Index], existingRoutes[route2Index+1:]...)
					route2Index--
					if route2Index < route1Index {
						route1Index--
					}
				}
			}
		}
	}
	return existingRoutes
}

func FindAllRoutes(startingRoom Room, endingRoom Room, allRooms []Room, existingRoutes *[][]string) (allRoutesNames [][]string) {
	potentialRoute := FindRoute(startingRoom, endingRoom, allRooms, []string{}, existingRoutes)
	for len(potentialRoute) != 0 {
		if !IsRouteConatinedInRoutes(potentialRoute, *existingRoutes) {
			*existingRoutes = append(*existingRoutes, potentialRoute)
		}
		potentialRoute = FindRoute(startingRoom, endingRoom, allRooms, []string{}, existingRoutes)
	}
	return *existingRoutes
}

func FindRoute(startingRoom Room, endingRoom Room, allRooms []Room, existingRoute []string, existingRoutes *[][]string) (routeNames []string) {
	existingRoute = append(existingRoute, startingRoom.Name)
	if startingRoom.Name == endingRoom.Name {
		return existingRoute
	}

	for i := 0; i < len(startingRoom.LinksAsPointers); i++ {
		currentLinkedRoom := startingRoom.LinksAsPointers[i]

		if !IsRoomContainedInRoute((*currentLinkedRoom).Name, existingRoute) && len(FindRoute(*currentLinkedRoom, endingRoom, allRooms, existingRoute, existingRoutes)) != 0 {
			potentialRoute := FindRoute(*currentLinkedRoom, endingRoom, allRooms, existingRoute, existingRoutes)
			if !IsRouteConatinedInRoutes(potentialRoute, *existingRoutes) && len(potentialRoute) != 0 {

				*existingRoutes = append(*existingRoutes, potentialRoute)
				existingRoute = potentialRoute

				return existingRoute
			}
		}
	}
	if existingRoute[len(existingRoute)-1] == endingRoom.Name {
		return existingRoute
	}

	return []string{}
}
