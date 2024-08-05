package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadAndParseFile(inputPath string, colony *Farm) (bool, string) {
    // Initialize variables
    firstline := false
    inputFile, err := os.Open(inputPath)
    if err != nil {
        fmt.Println("ERROR: unable to open file", err)
        os.Exit(1)
    }
    defer inputFile.Close()
    
    inputInfo := ""
    lineIndex := 0
    fileScanner := bufio.NewScanner(inputFile)
    isBeginningRoom := false
    isFinalRoom := false
    AntsCount := 0
    startRoomSet := false
    endRoomSet := false

    // Start scanning the file line by line
    for fileScanner.Scan() {
        line := fileScanner.Text()
        inputInfo += line + "\n"
        lineIndex++

        // Handle comments and special commands
        if strings.HasPrefix(line, "#") {
            if line == "##start" {
                if startRoomSet {
                    fmt.Printf("ERROR: invalid data format, multiple start rooms defined (%s)\n", line)
                    os.Exit(1)
                }
                startRoomSet = true
                isBeginningRoom = true
                continue
            } else if line == "##end" {
                if endRoomSet {
                    fmt.Printf("ERROR: invalid data format, multiple end rooms defined (%s)\n", line)
                    os.Exit(1)
                }
                endRoomSet = true
                isFinalRoom = true
                continue
            } else {
                continue
            }
        }

        // Parse the first line (number of ants)
        if !firstline {
            roomFields := strings.Fields(line)
            if len(roomFields) != 1 {
                fmt.Printf("⛔Invalid data format detected⛔ (%s)\n", line)
                os.Exit(1)
            }
            AntsCount, err = strconv.Atoi(roomFields[0])
            if err != nil || AntsCount <= 0 {
                fmt.Printf("ERROR: invalid data format, invalid number of ants (%s)\n", line)
                os.Exit(1)
            }
            firstline = true
            continue
        }

        // Handle start room definition
        if isBeginningRoom {
            isBeginningRoom = false
            parameters := strings.Fields(line)
            if len(parameters) != 3 {
                fmt.Printf("ERROR: invalid data format, invalid room definition (%s)\n", line)
                os.Exit(1)
            }
            room := parameters[0]
            xCoordinate, err1 := strconv.Atoi(parameters[1])
            yCoordinate, err2 := strconv.Atoi(parameters[2])
            if err1 != nil || err2 != nil {
                fmt.Printf("ERROR: invalid data format, invalid room definition (%s)\n", line)
                os.Exit(1)
            }
            colony.RegisterRoom(room, "start", xCoordinate, yCoordinate)
            continue
        }

        // Handle end room definition
        if isFinalRoom {
            isFinalRoom = false
            parameters := strings.Fields(line)
            if len(parameters) != 3 {
                fmt.Printf("ERROR: invalid data format, invalid room definition (%s)\n", line)
                os.Exit(1)
            }
            room := parameters[0]
            xCoordinate, err1 := strconv.Atoi(parameters[1])
            yCoordinate, err2 := strconv.Atoi(parameters[2])
            if err1 != nil || err2 != nil {
                fmt.Printf("ERROR: invalid data format, invalid room definition (%s)\n", line)
                os.Exit(1)
            }
            colony.RegisterRoom(room, "end", xCoordinate, yCoordinate)
            continue
        }

        // Handle room connections
        if strings.Contains(line, "-") && strings.Count(line, "-") == 1 {
            parameters := strings.Split(line, "-")
            if !colony.linkRooms(parameters[0], parameters[1], true) {
                os.Exit(1)
            }
            continue
        }

        // Handle normal room definitions
        roomDetails := strings.Fields(line)
        if len(roomDetails) != 3 && (!isBeginningRoom || !isFinalRoom) {
            continue
        }
        room := roomDetails[0]
        xCoordinate, _ := strconv.Atoi(roomDetails[1])
        yCoordinate, _ := strconv.Atoi(roomDetails[2])
        colony.RegisterRoom(room, "normal", xCoordinate, yCoordinate)
    }

    // Final checks
    if AntsCount == 0 {
        fmt.Printf("⛔Data format issue - No ants specified!⛔ (%s)\n")
        os.Exit(1)
    }

    if !startRoomSet || !endRoomSet {
        fmt.Println("ERROR: invalid data format, no start/end room found")
        os.Exit(1)
    }

    // Setup ants and return
    colony.SetupAnts(AntsCount)
    inputInfo += "\n"
    return true, inputInfo
}
