package functions

// CountTunnels counts the number of tunnels leading to rooms closer to the destination
func (farm *Farm) CountTunnels(chamber *AntRoom) int {
	// Initialize the tunnel count
	tunnelCount := 0

	// Start with the first connected room
	connection := chamber.connectedRooms.firstNode

	// Iterate through all connected rooms
	for connection != nil {
			// If the connected room is closer to the destination, increment the count
			if farm.roomPaths[connection.data] < farm.roomPaths[chamber] {
					tunnelCount++
			}
			// Move to the next connection
			connection = connection.nextConnection
	}

	// Return the total count of tunnels leading to closer rooms
	return tunnelCount
}

// Connection represents a link between rooms
type Connection struct {
	data           *AntRoom    // The room this connection leads to
	nextConnection *Connection // The next connection in the list
}

// RoomCollection is a linked list of connected rooms
type RoomCollection struct {
	firstNode *Connection // The first room in the collection
}

// AddRoom adds a new room to the collection
func (list *RoomCollection) AddRoom(roomToAdd *AntRoom) {
	// Create a new connection node for the room
	nodeToAdd := &Connection{data: roomToAdd, nextConnection: nil}

	// If the list is empty, make this the first node
	if list.firstNode == nil {
			list.firstNode = nodeToAdd
			return
	}

	// Otherwise, find the last node in the list
	currentNode := list.firstNode
	for currentNode.nextConnection != nil {
			currentNode = currentNode.nextConnection
	}

	// Add the new node to the end of the list
	currentNode.nextConnection = nodeToAdd
}

