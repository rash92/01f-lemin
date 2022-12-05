package main

import (
	"fmt"
	lemin "lemin/packages"
	"log"
)

func main() {
	numberofants, startingRoom, endingRoom, allRooms, roomLinks, outputError := lemin.ParseArgs()
	if outputError != nil {
		log.Fatal("couldn't parse file")
	}

	fmt.Println("detected error: ", outputError)
	fmt.Println("number of ants: ", numberofants)
	fmt.Println("starting room: ", startingRoom, "ending room: ", endingRoom)
	for i := 0; i < len(allRooms); i++ {
		fmt.Println("room ", i, "name: ", allRooms[i].Name)

		for j := 0; j < len(allRooms[i].LinksAsPointers); j++ {
			fmt.Println("room ", i, "has pointer to:", (*allRooms[i].LinksAsPointers[j]).Name)
		}
	}
	fmt.Println("all rooms: ", allRooms)
	fmt.Println("room links: ", roomLinks)

	routes := [][]string{}
	allroutes := lemin.FindAllRoutes(startingRoom, endingRoom, allRooms, routes)
	fmt.Println("all routes found are: ", allroutes)

	fmt.Println("sorted routes are: ", allroutes)
}
