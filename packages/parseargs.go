package lemin

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name            string
	Xcoord          int
	Ycoord          int
	CurrentAnt      Ant
	LinksAsStrings  []string
	LinksAsPointers []*Room
}

// takes as input a slice of rooms and a slice of string that has the room links and populates the LinksAsStrings field with the name of the room
func RoomLinkerStrings(unlinkedRooms []Room, roomLinks [][]string) []Room {
	for linkIndex := 0; linkIndex < len(roomLinks); linkIndex++ {
		for roomIndex := 0; roomIndex < len(unlinkedRooms); roomIndex++ {
			if unlinkedRooms[roomIndex].Name == roomLinks[linkIndex][0] {
				unlinkedRooms[roomIndex].LinksAsStrings = append(unlinkedRooms[roomIndex].LinksAsStrings, roomLinks[linkIndex][1])
			}
			if unlinkedRooms[roomIndex].Name == roomLinks[linkIndex][1] {
				unlinkedRooms[roomIndex].LinksAsStrings = append(unlinkedRooms[roomIndex].LinksAsStrings, roomLinks[linkIndex][0])
			}
		}
	}
	return unlinkedRooms
}

// takes as input a slice of rooms and a slice of string that has the room links and populates the LinksAsPointers field with a pointer to the room
// roomLinks is a slice with elements of the form [room1, room2] showing room1 is linked to room 2.
func RoomLinkerPointers(unlinkedRooms []Room, roomLinks [][]string) []Room {
	// fmt.Println("all room links are: ", roomLinks)
	for linkIndex := 0; linkIndex < len(roomLinks); linkIndex++ {
		for roomIndex := 0; roomIndex < len(unlinkedRooms); roomIndex++ {
			// if the current room is the first element in roomLinks, add the second elemenet as a link
			if unlinkedRooms[roomIndex].Name == roomLinks[linkIndex][0] {
				for roomToLinkIndex := 0; roomToLinkIndex < len(unlinkedRooms); roomToLinkIndex++ {
					if unlinkedRooms[roomToLinkIndex].Name == roomLinks[linkIndex][1] {
						unlinkedRooms[roomIndex].LinksAsPointers = append(unlinkedRooms[roomIndex].LinksAsPointers, &unlinkedRooms[roomToLinkIndex])
					}
				}
			}
			// if the current room is the second element in roomLinks, add the first elemenet as a link
			if unlinkedRooms[roomIndex].Name == roomLinks[linkIndex][1] {
				for roomToLinkIndex := 0; roomToLinkIndex < len(unlinkedRooms); roomToLinkIndex++ {
					if unlinkedRooms[roomToLinkIndex].Name == roomLinks[linkIndex][0] {
						unlinkedRooms[roomIndex].LinksAsPointers = append(unlinkedRooms[roomIndex].LinksAsPointers, &unlinkedRooms[roomToLinkIndex])
					}
				}
			}
		}
	}
	return unlinkedRooms
}

func ParseArgs() (numberofants int, startingroom Room, endingroom Room, allrooms []Room, roomLinks [][]string, outputError error) {
	filename := os.Args[1]

	data, outputError := os.ReadFile(filename)
	if outputError != nil {
		return 0, Room{}, Room{}, []Room{}, [][]string{}, outputError
	}

	instructions := strings.Split(string(data), "\n")

	// first line should always be number of ants
	numberofants, outputError = strconv.Atoi(instructions[0])
	if outputError != nil {
		return 0, Room{}, Room{}, []Room{}, [][]string{}, outputError
	}

	for i := 1; i < len(instructions); i++ {
		line := instructions[i]

		// prints current line as whole file should be printed in output as per instructions
		// fmt.Println(line)

		// splits line based on spaces to differentiate number of ants, room info, room links etc.
		words := strings.Fields(line)

		// only room info should be a line with 3 objects separated by spaces, first object is name second object is Xcoord third object is Ycoord
		if len(words) == 3 {
			currentRoom := Room{}
			currentRoom.Name = words[0]
			currentRoom.Xcoord, outputError = strconv.Atoi(words[1])
			if outputError != nil {
				return 0, Room{}, Room{}, []Room{}, [][]string{}, outputError
			}
			currentRoom.Ycoord, outputError = strconv.Atoi(words[2])
			if outputError != nil {
				return 0, Room{}, Room{}, []Room{}, [][]string{}, outputError
			}

			allrooms = append(allrooms, currentRoom)

			if instructions[i-1] == "##start" {
				startingroom = currentRoom
			}
			if instructions[i-1] == "##end" {
				endingroom = currentRoom
			}
		}

		// apart from number of ants on first line, only other time you should have only one object with no spaces is either room links or start or end tags
		// maybe have more checks for other possibilities for incorrect formats to throw errors
		if len(words) == 1 && line != "##start" && line != "##end" && line[0] != '#' {
			links := strings.Split(line, "-")
			if len(links) != 2 {
				return 0, Room{}, Room{}, []Room{}, [][]string{}, errors.New("roomlink incorrect format")
			}
			roomLinks = append(roomLinks, links)
		}
	}

	allrooms = RoomLinkerPointers(allrooms, roomLinks)
	allrooms = RoomLinkerStrings(allrooms, roomLinks)

	// fixes current and ending room to fill in links
	for _, currentRoom := range allrooms {
		if currentRoom.Name == startingroom.Name {
			startingroom = currentRoom
		}
		if currentRoom.Name == endingroom.Name {
			endingroom = currentRoom
		}
	}

	if startingroom.Name == "" || endingroom.Name == "" {
		fmt.Println("starting or ending room missing")
		return 0, Room{}, Room{}, []Room{}, [][]string{}, errors.New("starting or ending room missing")
	}

	return numberofants, startingroom, endingroom, allrooms, roomLinks, nil
}
