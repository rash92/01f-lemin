package lemin

import "fmt"

func IsContainedIn(element string, slice []string) bool {
	for _, potentialelement := range slice {
		if element == potentialelement {
			return true
		}
	}
	return false
}

func FindRoute(startingRoom Room, endingRoom Room, allRooms []Room, existingRoute []string) (routeNames []string) {
	fmt.Println("find route input info, starting room: ", startingRoom.Name, "starting room pointers :", startingRoom.LinksAsStrings, startingRoom.LinksAsPointers)

	if startingRoom.Name == endingRoom.Name {

		existingRoute = append(existingRoute, endingRoom.Name)
		fmt.Println("discovered route at endpoint is: ", existingRoute)
		return existingRoute
	}

	for _, currentLinkedRoom := range startingRoom.LinksAsPointers {
		if (*currentLinkedRoom).Name == endingRoom.Name {

			existingRoute = append(existingRoute, (*currentLinkedRoom).Name)
			fmt.Println("discovered route is: ", routeNames)
			return existingRoute
		}
		if !IsContainedIn((*currentLinkedRoom).Name, existingRoute) {
			existingRoute = append(existingRoute, (*currentLinkedRoom).Name)
			existingRoute = FindRoute(*currentLinkedRoom, endingRoom, allRooms, existingRoute)
			fmt.Println("discovered route is: ", existingRoute)
			return existingRoute
		}
	}

	return existingRoute
}
