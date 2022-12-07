package main

import (
	"fmt"
	lemin "lemin/packages"
	"log"
)

func main() {
	numberofants, startingRoom, endingRoom, allRooms, _, outputError := lemin.ParseArgs()
	// _, startingRoom, endingRoom, allRooms, _, outputError := lemin.ParseArgs()

	if outputError != nil {
		fmt.Println(outputError)
		log.Fatal("couldn't parse file")
	}

	// fmt.Println("detected error: ", outputError)
	// fmt.Println("number of ants: ", numberofants)
	// fmt.Println("starting room: ", startingRoom, "ending room: ", endingRoom)
	// fmt.Println("all rooms: ", allRooms)
	// fmt.Println("room links: ", roomLinks)

	// routes := [][]lemin.Room{}
	allroutes := lemin.FindAllRoutes(startingRoom, endingRoom, allRooms, [][]lemin.Room{})

	// for i, route := range allroutes {
	// 	fmt.Println("route number: ", i+1)
	// 	for j, room := range route {
	// 		fmt.Println("room ", j, " in route: ", room.Name)
	// 	}
	// }

	getAnts := AntHandler(allroutes, numberofants)

	fmt.Println(getAnts)
}

func AntHandler(routes [][]lemin.Room, numberofants int) []int {
	// get number of rooms in a route
	lengthOfRoute := len(routes)

	antsOnPath := make([]int, lengthOfRoute)

	roomsInPaths := []int{}

	for i := lengthOfRoute; i > 0; i-- {
		numberOfRoomsInRoute := len(routes[i-1])
		roomsInPaths = append(roomsInPaths, numberOfRoomsInRoute)
	}

	for numberofants > 0 {

		if len(roomsInPaths) == 1 {
			antsOnPath[0] = numberofants
			numberofants = 0
		} else {
			for i := len(roomsInPaths) - 1; i > 0; i-- {
				for j := i - 1; j >= 0; j-- {
					if numberofants == 0 {
						break
					}
					prev := roomsInPaths[i] + antsOnPath[i]
					next := roomsInPaths[j] + antsOnPath[j]
					if prev > next {
						antsOnPath[j]++
					} else {
						antsOnPath[i]++
					}
					numberofants--
				}
			}
		}
	}

	return antsOnPath
}
