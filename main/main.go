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
	// 	fmt.Println("route number: ", i)
	// 	for j, room := range route {
	// 		fmt.Println("room ", j, " in route: ", room.Name)
	// 	}
	// }
	// fmt.Println("ant paths: ")
	AntHandler(allroutes, numberofants)

}

func AntHandler(routes [][]lemin.Room, numberofants int) {

}
