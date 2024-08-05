package functions

func (colony *Farm) RepositionAnts() {
	// Reset all ants to their initial position
	for i := 0; i < colony.antCount; i++ {
			// Set current room to initial room for each ant
			colony.workers[i].currentRoom = colony.initialRoom

			// Clear visited room history for each ant
			for roomKey := range colony.workers[i].visitedRoom {
					colony.workers[i].visitedRoom[roomKey] = false
			}
			// Mark initial room as not visited
			colony.workers[i].visitedRoom[colony.initialRoom] = false
			// Set ant status to "in motion"
			colony.workers[i].inMotion = true
	}
	// Reset all tunnel access after repositioning ants
	colony.ResetTunnelAccess()
}

func (colony *Farm) ResetTunnelAccess() {
	roomCollection := colony.roomMap

	// Iterate through all rooms in the colony
	for roomKey := range roomCollection {
			currentRoom := roomCollection[roomKey]

			// Set all tunnel accessibility to false for each room
			for tunnelID := range currentRoom.accessibility {
					currentRoom.accessibility[tunnelID] = false
			}
	}
}

// RoomStack is a simple stack data structure for rooms
type RoomStack struct {
	roomList []*AntRoom
}

// AddRoomToStack adds a room to the top of the stack
func (stack *RoomStack) AddRoomToStack(room *AntRoom) {
	stack.roomList = append(stack.roomList, room)
}

// RemoveRoom removes and returns the top room from the stack
func (stack *RoomStack) RemoveRoom() *AntRoom {
	if len(stack.roomList) == 0 {
			return nil // Return nil if stack is empty
	}
	removedRoom := stack.roomList[0]
	stack.roomList = stack.roomList[1:] // Update stack by removing first element
	return removedRoom
}

