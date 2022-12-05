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

	AntHandler(allroutes, numberofants)
}

func AntHandler(routes [][]lemin.Room, numberofants int) {
	// get number of rooms in a route
	lengthOfRoute := len(routes)
	fmt.Println("possible routes", lengthOfRoute)
	fmt.Println("number of ants", numberofants)

	antsOnPath := make([]int, lengthOfRoute)
	fmt.Println("ants on paths PREVIOUS", antsOnPath)

	roomsInPaths := []int{}

	for i := lengthOfRoute; i > 0; i-- {
		numberOfRoomsInRoute := len(routes[i-1])
		roomsInPaths = append(roomsInPaths, numberOfRoomsInRoute)
	}

	fmt.Println("rooms in paths", roomsInPaths)

	for numberofants > 0 {
		fmt.Println("Number of ants", numberofants)
		for i := 0; i < len(roomsInPaths)-1; i++ {
			if (roomsInPaths[i] + antsOnPath[i]) < (roomsInPaths[i+1] + antsOnPath[i+1]) {
				antsOnPath[i]++
			} else {
				antsOnPath[i+1]++
			}
		}
		numberofants--
	}

	fmt.Println("ants on paths AFTER", antsOnPath)

}
