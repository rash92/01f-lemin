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
	fmt.Println("number of ants: ", numberofants)
	// fmt.Println("starting room: ", startingRoom, "ending room: ", endingRoom)
	// fmt.Println("all rooms: ", allRooms)
	// fmt.Println("room links: ", roomLinks)

	// routes := [][]lemin.Room{}
	allroutes := lemin.FindAllRoutes(startingRoom, endingRoom, allRooms, [][]lemin.Room{})

	for _, route := range allroutes {
		for j, room := range route {
			fmt.Println("room ", j, " in route: ", room.Name)
		}
	}
	// getAntsOnRoute := AntHandler(allroutes, numberofants)
	// fmt.Println(getAntsOnRoute)
	// fmt.Println(allroutes)

	// i := 1

	// for i < numberofants+1 {
	// 	for _, roomlinks := range allroutes {
	// 		// fmt.Println("route:", route)
	// 		for links := range roomlinks {
	// 			fmt.Print("L", i, "-", links)
	// 		}
	// 		fmt.Println()
	// 	}
	// 	i++
	// }

	antsRoute := RoutesAntsWillTake(allroutes, numberofants)
	// fmt.Println("ants route:", antsRoute)

	// roomsInRoute := []int{}
	// lengthOfRoutes := len(allroutes)

	fmt.Println(allRooms)

}

func RoutesAntsWillTake(routes [][]lemin.Room, numberofants int) []int {

	// make a slice with the length of the number of routes
	lengthOfRoutes := len(routes)
	antsOnRoute := make([]int, lengthOfRoutes)

	roomsInRoute := []int{}

	for i := lengthOfRoutes; i > 0; i-- {
		numberOfRoomsInRoute := len(routes[i-1])
		roomsInRoute = append(roomsInRoute, numberOfRoomsInRoute)
	}

	for numberofants > 0 {
		if len(roomsInRoute) == 1 {
			antsOnRoute[0]++
			numberofants--
		}
		for i := len(roomsInRoute) - 1; i > 0; i-- {
			for j := i - 1; j >= 0; j-- {
				if numberofants == 0 {
					break
				}
				prev := roomsInRoute[i] + antsOnRoute[i]
				next := roomsInRoute[j] + antsOnRoute[j]
				if prev > next {
					antsOnRoute[j]++
				} else {
					antsOnRoute[i]++
				}
				numberofants--
			}
		}
	}
	return antsOnRoute
	fmt.Println("length of empty set of rooms is: ", len([]lemin.Room{}))
	fmt.Println("number of paths is: ", len(allroutes), "number of ants is: ", numberofants)
	fmt.Println("ant paths: ", lemin.AssignNumberOfAnts(allroutes, numberofants))
	// lemin.AntHandler(allroutes, numberofants)
	antSlice := lemin.IdentifyAnts(numberofants, lemin.AssignNumberOfAnts(allroutes, numberofants), allroutes)
	fmt.Println("antSlice is: ", antSlice)
	fmt.Println(lemin.AssignAntsPerRoute(antSlice, allroutes))
	// lemin.PrintAnts(lemin.IdentifyAnts(numberofants, lemin.AssignAnts(allroutes, numberofants), allroutes), allroutes, numberofants)
}
