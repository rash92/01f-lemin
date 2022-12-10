package main

import (
	"fmt"
	lemin "lemin/packages"
)

func main() {
	numberofants, startingRoom, endingRoom, allRooms, instruction, outputError := lemin.ParseArgs()
	if outputError != nil {
		fmt.Println("ERROR: invalid data format")
		return
	}

	if numberofants <= 0 {
		fmt.Println("ERROR: invalid data format, invalid number of Ants")
		return
	}

	routes := [][]lemin.Room{}
	allroutes := lemin.FindAllRoutes(startingRoom, endingRoom, allRooms, routes)

	if len(allroutes) <= 0 {
		fmt.Println("ERROR: invalid data format, no route found")
		return
	}

	antSlice := lemin.IdentifyAnts(numberofants, lemin.AssignNumberOfAnts(allroutes, numberofants), allroutes)

	antsPerRoute := lemin.AssignAntsPerRoute(antSlice, allroutes)

	fmt.Println(numberofants)
	fmt.Println(instruction)
	lemin.PrintAnts(antsPerRoute, allroutes, numberofants)
}
