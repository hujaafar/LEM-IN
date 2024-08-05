package functions

import "math"

// ColonyWorker represents an individual ant in the colony
type ColonyWorker struct {
	antNumber        int            // Unique identifier for the ant
	currentRoom      *AntRoom       // The room where the ant is currently located
	visitedRoom      map[*AntRoom]bool // Keeps track of rooms visited by the ant
	inMotion         bool           // Indicates if the ant is currently moving
	hasCompletedMove bool           // Indicates if the ant has finished its move
}

// AntRoom represents a room in the ant colony
type AntRoom struct {
	connectedRooms *RoomCollection  // Rooms connected to this room
	isBeginning    bool             // True if this is the starting room
	isDestination  bool             // True if this is the destination room
	isUnoccupied   bool             // True if the room is currently empty
	xCoordinate    int              // X-coordinate of the room
	yCoordinate    int              // Y-coordinate of the room
	roomName       string           // Name of the room
	accessibility  map[string]bool  // Indicates which tunnels are accessible
	isEndpoint     bool             // True if this room is an endpoint
}

// Farm represents the entire ant colony
type Farm struct {
	roomMap     map[string]*AntRoom // Map of all rooms in the colony
	antCount    int                 // Total number of ants
	roomPaths   map[*AntRoom]int    // Stores the shortest path to each room
	workers     []*ColonyWorker     // Slice of all ants in the colony
	initialRoom *AntRoom            // Starting room for the ants
	finalRoom   *AntRoom            // Destination room for the ants
}

// SetupAnts initializes the ants in the colony
func (colony *Farm) SetupAnts(antCount int) {
	colony.antCount = antCount
	colony.workers = make([]*ColonyWorker, colony.antCount)

	// Initialize each ant
	for i := 0; i < antCount; i++ {
			colony.workers[i] = new(ColonyWorker)
			colony.workers[i].currentRoom = colony.initialRoom
			colony.workers[i].visitedRoom = make(map[*AntRoom]bool)
			colony.workers[i].visitedRoom[colony.initialRoom] = false
			colony.workers[i].inMotion = true
			colony.workers[i].antNumber = i + 1
	}
}

// SetupPaths initializes the shortest path to each room as infinity
func (colony *Farm) SetupPaths() {
	for index := range colony.roomMap {
			colony.roomPaths[colony.roomMap[index]] = math.MaxInt32
	}
}

// Setup initializes the basic structures of the Farm
func (colony *Farm) Setup() {
	colony.roomMap = make(map[string]*AntRoom)
	colony.roomPaths = make(map[*AntRoom]int)
}

