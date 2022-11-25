package main

import (
	"fmt"
	"lemin"
)

func main() {
	numberofants, startingroom, endingroom, allrooms, roomLinks, outputError := lemin.ParseArgs()

	fmt.Println("number of ants: ", numberofants)
	fmt.Println("starting room: ", startingroom, "ending room: ", endingroom)
	fmt.Println("all rooms: ", allrooms)
	fmt.Println("room links: ", roomLinks)
	fmt.Println("first room: ", allrooms[0])
	fmt.Println("first room links fields: ", allrooms[0].LinksAsStrings, allrooms[0].LinksAsPointers)
	fmt.Println("detected error: ", outputError)
}
