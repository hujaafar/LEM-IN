package functions

import (
	"fmt"
	"os"
)

// RegisterRoom registers a room in the colony with proper error handling.
func (colony *Farm) RegisterRoom(roomName string, roomType string, xCoord int, yCoord int) {
	if _, exists := colony.roomMap[roomName]; exists {
		fmt.Printf("ERROR: invalid data format, duplicate room name '%s'\n", roomName)
		os.Exit(1) // Exit the program with an error code
	}

	switch roomType {
	case "start":
		colony.roomMap[roomName] = &AntRoom{
			roomName:       roomName,
			isBeginning:    true,
			isDestination:  false,
			isUnoccupied:   true,
			xCoordinate:    xCoord,
			yCoordinate:    yCoord,
			connectedRooms: &RoomCollection{},
			accessibility:  make(map[string]bool),
		}
		colony.initialRoom = colony.roomMap[roomName]

	case "end":
		colony.roomMap[roomName] = &AntRoom{
			roomName:       roomName,
			isBeginning:    false,
			isDestination:  true,
			isUnoccupied:   true,
			xCoordinate:    xCoord,
			yCoordinate:    yCoord,
			connectedRooms: &RoomCollection{},
			accessibility:  make(map[string]bool),
		}
		colony.finalRoom = colony.roomMap[roomName]

	case "normal":
		colony.roomMap[roomName] = &AntRoom{
			roomName:       roomName,
			isBeginning:    false,
			isDestination:  false,
			isUnoccupied:   true,
			xCoordinate:    xCoord,
			yCoordinate:    yCoord,
			connectedRooms: &RoomCollection{},
			accessibility:  make(map[string]bool),
		}

	default:
		fmt.Printf("ERROR: invalid data format, invalid room type '%s'\n", roomType)
		os.Exit(1) // Exit the program with an error code
	}
}



// linkRooms links two rooms and handles invalid room names and duplicate connections.
func (colony *Farm) linkRooms(fromRoomName string, toRoomName string, isBidirectional bool) bool {
    fromRoom := colony.roomMap[fromRoomName]
    toRoom := colony.roomMap[toRoomName]

    if fromRoom == nil || toRoom == nil {
        // This checks if either of the rooms does not exist and prints an appropriate error message
        fmt.Println("ERROR: invalid data format, invalid room definition")
        return false
    }

    // Check for duplicate connection
    if fromRoom.HasTunnelTo(toRoom) {
        fmt.Printf("ERROR: duplicate connection between rooms '%s' and '%s'\n", fromRoomName, toRoomName)
        return false
    }

    // Add the connection if no errors were found
    if isBidirectional {
        fromRoom.connectedRooms.AddRoom(toRoom)
        toRoom.connectedRooms.AddRoom(fromRoom)
    } else {
        fromRoom.connectedRooms.AddRoom(toRoom)
    }
    return true
}