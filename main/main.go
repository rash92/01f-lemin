package main

import (
	"fmt"
	lemin "lemin/packages"
	"log"
)

func main() {
	numberofants, startingRoom, endingRoom, allRooms, outputError := lemin.ParseArgs()
	if outputError != nil {
		fmt.Println(outputError)
		log.Fatal("couldn't parse file")
	}

	// fmt.Println("detected error: ", outputError)
	// fmt.Println("number of ants: ", numberofants)
	// fmt.Println("starting room: ", startingRoom.Name, "ending room: ", endingRoom.Name)
	// fmt.Println("all rooms: ", allRooms)
	// fmt.Println("room links: ", roomLinks)

	routes := [][]lemin.Room{}
	allroutes := lemin.FindAllRoutes(startingRoom, endingRoom, allRooms, routes)
	// for i, route := range allroutes {
	// 	fmt.Println("route number: ", i)
	// 	for j, room := range route {
	// 		fmt.Println("room ", j, " in route: ", room.Name)
	// 	}
	// }
	// fmt.Println("length of empty set of rooms is: ", len([]lemin.Room{}))
	// fmt.Println("number of paths is: ", len(allroutes), "number of ants is: ", numberofants)
	// fmt.Println("ant paths: ", lemin.AssignNumberOfAnts(allroutes, numberofants))
	// lemin.AntHandler(allroutes, numberofants)
	antSlice := lemin.IdentifyAnts(numberofants, lemin.AssignNumberOfAnts(allroutes, numberofants), allroutes)
	// fmt.Println("antSlice is: ", antSlice)
	antsPerRoute := lemin.AssignAntsPerRoute(antSlice, allroutes)

	lemin.PrintAnts(antsPerRoute, allroutes, numberofants)
	// lemin.PrintAnts(lemin.IdentifyAnts(numberofants, lemin.AssignAnts(allroutes, numberofants), allroutes), allroutes, numberofants)
}
