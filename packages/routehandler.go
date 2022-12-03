package lemin

import (
	"fmt"
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
	fmt.Println("find route input info, starting room: ", startingRoom.Name, "starting room pointers :", startingRoom.LinksAsStrings, startingRoom.LinksAsPointers)
	existingRoute = append(existingRoute, startingRoom.Name)
	if startingRoom.Name == endingRoom.Name {
		fmt.Println("discovered route at endpoint is: ", existingRoute)
		return existingRoute
	}

	for i := 0; i < len(startingRoom.LinksAsPointers); i++ {
		currentLinkedRoom := startingRoom.LinksAsPointers[i]

		if !IsRoomContainedInRoute((*currentLinkedRoom).Name, existingRoute) && len(FindRoute(*currentLinkedRoom, endingRoom, allRooms, existingRoute, existingRoutes)) != 0 {
			potentialRoute := FindRoute(*currentLinkedRoom, endingRoom, allRooms, existingRoute, existingRoutes)
			if !IsRouteConatinedInRoutes(potentialRoute, *existingRoutes) && len(potentialRoute) != 0 {
				fmt.Println("discovered route in long conditional is: ", existingRoute)
				*existingRoutes = append(*existingRoutes, potentialRoute)
				existingRoute = potentialRoute
				fmt.Println("all discovered routes after long conditional are: ", existingRoutes)
				return existingRoute
			}
		}
	}
	if existingRoute[len(existingRoute)-1] == endingRoom.Name {
		fmt.Println("got to checking if route ends at ending room")
		return existingRoute
	}

	fmt.Println("got to end without meeting any other conditions")
	return []string{}
}
