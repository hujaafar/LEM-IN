package main

import (
	"fmt"
	ants "lem-in/functions"
	"os"
)

func main(){
	// Get command-line arguments
	parameters := os.Args[1:]

	// Check if the correct number of arguments is provided
	if len(parameters) != 1 {
		fmt.Println("⛔ Run the program using: go run <path_to_farm_file> ⛔")
		return
	}

	// Initialize the ant colony
	var colony ants.Farm
	colony.Setup()

	// Read and parse the input file, skipping section between ##start and ##end
	fileReadSuccess, _ := ants.ReadAndParseFile(parameters[0], &colony)
	if !fileReadSuccess {
		return
	}

	// Verify that all rooms have distinct coordinates
	if !colony.CheckDistinctRoomCoordinates() {
		fmt.Println("⛔Data format issue⛔")
		return
	}

	// Set up paths, find the shortest path, print info, and find the optimal ant path
	colony.SetupPaths()
	colony.FindShortestPath()
	colony.FindOptimalAntPath()
}