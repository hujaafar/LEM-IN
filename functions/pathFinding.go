package functions

import (
	"fmt"
	"math"
	"os"
	"sort"
)

var movementOccurred = false // Global flag to track if any ant movement occurred

// printAndExit prints a message and exits the program with the given exit code
func printAndExit(message string, exitCode int) {
    fmt.Println(message)
    os.Exit(exitCode)
}

// FindShortestPath calculates the shortest path from the final room to all other rooms
func (farm *Farm) FindShortestPath() {
    stack := &RoomStack{}
    exploredRooms := make(map[*AntRoom]bool)
    seenRooms := make(map[*AntRoom]bool)

    // Start from the final room
    stack.AddRoomToStack(farm.finalRoom)
    current := stack.RemoveRoom()
    farm.roomPaths[current] = 1
    seenRooms[current] = true

    // Breadth-first search
    for current != nil {
        neighborNode := current.connectedRooms.firstNode
        for neighborNode != nil {
            if !seenRooms[neighborNode.data] {
                if !neighborNode.data.isBeginning {
                    stack.AddRoomToStack(neighborNode.data)
                }
                seenRooms[neighborNode.data] = true
                if farm.roomPaths[current]+1 < farm.roomPaths[neighborNode.data] {
                    farm.roomPaths[neighborNode.data] = farm.roomPaths[current] + 1
                }
            }
            if neighborNode.data == farm.initialRoom {
                exploredRooms[neighborNode.data] = true
            }
            neighborNode = neighborNode.nextConnection
        }
        exploredRooms[current] = true
        current = stack.RemoveRoom()
    }
    
    // Check if initial room was explored
    if !exploredRooms[farm.initialRoom] {
        printAndExit("⛔ Initial Room Not Explored⛔", 0)
    }
}

// AreAllAntsAtFinalRoom checks if all ants have reached the final room
func (colony *Farm) AreAllAntsAtFinalRoom() bool {
    for _, worker := range colony.workers {
        if worker.currentRoom != colony.finalRoom {
            return false
        }
    }
    return true
}

// AdvanceAntsInFarm moves ants through the farm
func (colony *Farm) AdvanceAntsInFarm(toggle bool) {
    activeWorkers := make([]*ColonyWorker, colony.antCount)
    copy(activeWorkers, colony.workers)

    // Sort workers based on the number of tunnels in their current room
    sort.SliceStable(activeWorkers, func(workerIdx1, workerIdx2 int) bool {
        return colony.CountTunnels(activeWorkers[workerIdx1].currentRoom) < colony.CountTunnels(activeWorkers[workerIdx2].currentRoom)
    })

    for workerIndex := 0; workerIndex < len(activeWorkers); workerIndex++ {
        nextPath := colony.DetermineNextRoom(activeWorkers[workerIndex], toggle)

        if nextPath != colony.finalRoom || activeWorkers[workerIndex].currentRoom.isBeginning {
            colony.roomPaths[nextPath]++
        }
        if canMoveToRoom(activeWorkers[workerIndex], nextPath) {
            // Move the ant to the next room
            activeWorkers[workerIndex].currentRoom.accessibility[nextPath.roomName] = true
            activeWorkers[workerIndex].visitedRoom[activeWorkers[workerIndex].currentRoom] = true
            activeWorkers[workerIndex].currentRoom.isUnoccupied = true
            activeWorkers[workerIndex].currentRoom = nextPath
            nextPath.isUnoccupied = false
            if activeWorkers[workerIndex].currentRoom.isDestination {
                activeWorkers[workerIndex].inMotion = false
            }
            activeWorkers[workerIndex].hasCompletedMove = true
            movementOccurred = true
        }

        // Reset worker index if movement occurred and we're at the last worker
        if workerIndex == len(activeWorkers)-1 && movementOccurred {
            workerIndex = 0
            movementOccurred = false
        }
    }
    colony.ResetTunnelAccess()
    colony.FindShortestPath()
}

// FindOptimalAntPath finds the optimal path for ants to reach the final room
func (colony *Farm) FindOptimalAntPath() {
    onSteps, offSteps, onPath, offPath, currentStep := 0, 0, "", "", 0

    // First attempt without optimization
    for !colony.AreAllAntsAtFinalRoom() {
        offSteps++
        colony.AdvanceAntsInFarm(false)
        offPath += colony.FormatAntLocations()
    }

    // Second attempt with optimization
    colony.RepositionAnts()
    for !colony.AreAllAntsAtFinalRoom() {
        currentStep++
        onSteps++
        if onSteps > offSteps {
            break
        }
        colony.AdvanceAntsInFarm(true)
        onPath += colony.FormatAntLocations()
    }

    // Compare results and print the optimal path
    if offSteps == onSteps {
        fmt.Printf("\n%s\nSteps taken: %d\n", onPath, onSteps)
    } else if offSteps < onSteps {
        fmt.Printf("\n%s\nSteps taken: %d\n", offPath, offSteps)
    } else {
        fmt.Printf("\n%s\nSteps taken: %d\n", onPath, onSteps)
    }
}

// DetermineNextRoom finds the best next room for an ant to move to
func (colony *Farm) DetermineNextRoom(ant *ColonyWorker, isStrict bool) *AntRoom {
    shortestPath := math.MaxInt32
    bestRoom := ant.currentRoom.connectedRooms.firstNode.data

    currentConnection := ant.currentRoom.connectedRooms.firstNode
    for currentConnection != nil {
        if isStrict {
            if colony.roomPaths[currentConnection.data] < shortestPath && !ant.visitedRoom[currentConnection.data] {
                bestRoom = currentConnection.data
                shortestPath = colony.roomPaths[bestRoom]
            }
        } else {
            if colony.roomPaths[currentConnection.data] <= shortestPath && !ant.visitedRoom[currentConnection.data] {
                bestRoom = currentConnection.data
                shortestPath = colony.roomPaths[bestRoom]
            }
        }
        currentConnection = currentConnection.nextConnection
    }
    return bestRoom
}

// canMoveToRoom checks if an ant can move to a specific room
func canMoveToRoom(ant *ColonyWorker, destinationRoom *AntRoom) bool {
    return (destinationRoom.isUnoccupied || destinationRoom.isDestination) &&
           !ant.visitedRoom[destinationRoom] &&
           !destinationRoom.isBeginning &&
           !ant.currentRoom.accessibility[destinationRoom.roomName] &&
           ant.inMotion &&
           !destinationRoom.isEndpoint &&
           !ant.hasCompletedMove
}

