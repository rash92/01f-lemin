package main

import (
	"fmt"
	"lemin"
)

func main() {
	numberofants, startingroom, endingroom, allrooms, roomLinks, outputError := lemin.ParseArgs()

	fmt.Println("detected error: ", outputError)
	fmt.Println("number of ants: ", numberofants)
	fmt.Println("starting room: ", startingroom, "ending room: ", endingroom)
	for i := 0; i < len(allrooms); i++ {
		fmt.Println("room ", i, "name: ", allrooms[i].Name)

		for j := 0; j < len(allrooms[i].LinksAsPointers); j++ {
			fmt.Println("room ", i, "pointer ", j, "is: ", (*allrooms[i].LinksAsPointers[j]).Name)
		}
	}
	fmt.Println("all rooms: ", allrooms)
	fmt.Println("room links: ", roomLinks)
	fmt.Println("first room: ", allrooms[0])
	fmt.Println("first room links fields: ", allrooms[0].LinksAsStrings, allrooms[0].LinksAsPointers)
}
