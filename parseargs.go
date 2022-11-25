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
	LinksAsStrings  []string
	LinksAsPointers []*Room
}

// takes as input a slice of rooms and a slice of string that has the room links and populates the links field in the rooms
func RoomLinkerStrings(unlinkedRooms []Room, roomLinks [][]string) []Room {
	for roomLinkIndex := 0; roomLinkIndex < len(roomLinks); roomLinkIndex++ {
		for roomIndex := 0; roomIndex < len(unlinkedRooms); roomIndex++ {
			if unlinkedRooms[roomIndex].Name == roomLinks[roomLinkIndex][0] {
				unlinkedRooms[roomIndex].LinksAsStrings = append(unlinkedRooms[roomIndex].LinksAsStrings, roomLinks[roomLinkIndex][1])
			}
		}
	}
	return unlinkedRooms
}

func RoomLinkerPointers(unlinkedRooms []Room, roomLinks [][]string) []Room {
	for roomLinkIndex := 0; roomLinkIndex < len(roomLinks); roomLinkIndex++ {
		for roomIndex := 0; roomIndex < len(unlinkedRooms); roomIndex++ {
			if unlinkedRooms[roomIndex].Name == roomLinks[roomLinkIndex][0] {
				for roomToLinkIndex := 0; roomToLinkIndex < len(unlinkedRooms); roomToLinkIndex++ {
					if unlinkedRooms[roomToLinkIndex].Name == roomLinks[roomToLinkIndex][1] {
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
		fmt.Println(line)

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
		if len(words) == 1 && line != "##start" && line != "##end" {
			links := strings.Split(line, "-")
			if len(links) != 2 {
				return 0, Room{}, Room{}, []Room{}, [][]string{}, errors.New("roomlink incorrect format")
			}
			roomLinks = append(roomLinks, links)
		}
	}

	allrooms = RoomLinkerPointers(allrooms, roomLinks)
	allrooms = RoomLinkerStrings(allrooms, roomLinks)

	return numberofants, startingroom, endingroom, allrooms, roomLinks, nil
}
