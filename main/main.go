package main

import (
	"fmt"
	lemin "lemin/packages"
)

func main() {
	numberofants, startingRoom, endingRoom, allRooms, roomLinks, outputError := lemin.ParseArgs()

	fmt.Println("detected error: ", outputError)
	fmt.Println("number of ants: ", numberofants)
	fmt.Println("starting room: ", startingRoom, "ending room: ", endingRoom)
	for i := 0; i < len(allRooms); i++ {
		fmt.Println("room ", i, "name: ", allRooms[i].Name)

		for j := 0; j < len(allRooms[i].LinksAsPointers); j++ {
			fmt.Println("room ", i, "pointer ", j, "is: ", (*allRooms[i].LinksAsPointers[j]).Name)
		}
	}
	fmt.Println("all rooms: ", allRooms)
	fmt.Println("room links: ", roomLinks)
	fmt.Println("first room: ", allRooms[0])
	fmt.Println("first room links fields: ", allRooms[0].LinksAsStrings, allRooms[0].LinksAsPointers)
	fmt.Println("find route input info, starting room: ", startingRoom.Name, "starting room pointers :", startingRoom.LinksAsStrings, startingRoom.LinksAsPointers)
	fmt.Println("find route input info, ending room: ", endingRoom.Name, "ending room pointers :", endingRoom.LinksAsStrings, endingRoom.LinksAsPointers)

	routes := [][]string{}
	allroutes := lemin.FindAllRoutes(startingRoom, endingRoom, allRooms, &routes)
	fmt.Println("all routes found are: ", allroutes)
}
