package main

import (
	"os"      // for reading from files
	"strconv" // our old friend string conversion package
	"strings" // for working with strings
)

func ReadElectoralVotes(filename string) map[string]uint {
	electoralVotes := make(map[string]uint)

	fileContents, err := os.ReadFile(filename)
	Check(err)

	giantString := string(fileContents)
	lines := strings.Split(giantString, "\n")

	for _, currentLine := range lines {
		lineElements := strings.Split(currentLine, ",")
		stateName := lineElements[0]
		votes, err := strconv.Atoi(lineElements[1])
		Check(err)
		electoralVotes[stateName] = uint(votes)
	}

	return electoralVotes
}

func ReadPollingData(filename string) map[string]float64 {
	candidate1Percentages := make(map[string]float64)

	fileContents, err := os.ReadFile(filename)
	Check(err)

	giantString := string(fileContents)
	lines := strings.Split(giantString, "\n")

	for _, currentLine := range lines {
		lineElements := strings.Split(currentLine, ",")
		stateName := lineElements[0]
		percentage1, err := strconv.ParseFloat(lineElements[1], 64)
		Check(err)
		candidate1Percentages[stateName] = percentage1 / 100
	}
	return candidate1Percentages
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
