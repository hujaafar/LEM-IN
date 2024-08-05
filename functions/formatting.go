package functions

import (
	"fmt"
	"sort"
)

func (colony *Farm) FormatAntLocations() string {
	// Initialize an empty string to store formatted locations
	formattedLocations := ""

	// Create a new slice to hold ordered ants
	orderedAnts := make([]*ColonyWorker, colony.antCount)

	// Copy the workers to the new slice
	copy(orderedAnts, colony.workers)

	// Sort the ants by their number
	sort.SliceStable(orderedAnts, func(first, second int) bool {
		return orderedAnts[first].antNumber < orderedAnts[second].antNumber
	})

	// Iterate through the sorted ants
	for _, worker := range orderedAnts {
		// Check if the ant has completed a move
		if worker.hasCompletedMove {
			// Format the ant's location and add it to the string
			formattedLocations += fmt.Sprintf("L%d-%s ", worker.antNumber, worker.currentRoom.roomName)

			// Reset the move completion flag
			worker.hasCompletedMove = false
		}
	}

	// Return the formatted locations with a newline
	return formattedLocations + "\n"
}
