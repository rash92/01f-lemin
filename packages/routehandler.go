package lemin

import "fmt"

func FindRoute(startingRoom Room, endingRoom Room, allRooms []Room) (route []Room) {
	fmt.Println("find route input info, starting room: ", startingRoom.Name, "starting room pointers :", startingRoom.LinksAsStrings, startingRoom.LinksAsPointers)

	for _, currentLinkedRoom := range startingRoom.LinksAsPointers {
		if (*currentLinkedRoom).Name == endingRoom.Name {
			route = append(route, *currentLinkedRoom)
			fmt.Println("discovered route is: ", route)
			return route
		} else if len(FindRoute(*currentLinkedRoom, endingRoom, allRooms)) != 0 {
			for _, nextLinkedRoom := range currentLinkedRoom.LinksAsPointers {
				if len(FindRoute(*nextLinkedRoom, endingRoom, allRooms)) != 0 && (*nextLinkedRoom).LinksAsStrings[0] != startingRoom.Name {
					route = append(route, FindRoute(*nextLinkedRoom, endingRoom, allRooms)...)
					return route
				}
			}
		}
	}

	return route
}
