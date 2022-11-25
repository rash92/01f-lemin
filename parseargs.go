package lemin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name   string
	xcoord int
	ycoord int
}

func ParseArgs() (numberofants int, startingroom Room, endingroom Room, allrooms []Room, roomLinks []string, outputError error) {
	filename := os.Args[1]

	data, outputError := os.ReadFile(filename)
	if outputError != nil {
		return 0, Room{}, Room{}, []Room{}, []string{}, outputError
	}

	instructions := strings.Split(string(data), "\n")

	// first line should always be number of ants
	numberofants, outputError = strconv.Atoi(instructions[0])
	if outputError != nil {
		return 0, Room{}, Room{}, []Room{}, []string{}, outputError
	}

	for i := 1; i < len(instructions); i++ {
		line := instructions[i]

		// prints current line as whole file should be printed in output as per instructions
		fmt.Println(line)

		// splits line based on spaces to differentiate number of ants, room info, room links etc.
		words := strings.Fields(line)

		// only room info should be a line with 3 objects separated by spaces, first object is name second object is xcoord third object is ycoord
		if len(words) == 3 {
			currentRoom := Room{}
			currentRoom.Name = words[0]
			currentRoom.xcoord, outputError = strconv.Atoi(words[1])
			if outputError != nil {
				return 0, Room{}, Room{}, []Room{}, []string{}, outputError
			}
			currentRoom.ycoord, outputError = strconv.Atoi(words[2])
			if outputError != nil {
				return 0, Room{}, Room{}, []Room{}, []string{}, outputError
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
			roomLinks = append(roomLinks, line)
		}
	}

	return numberofants, startingroom, endingroom, allrooms, roomLinks, nil
}
