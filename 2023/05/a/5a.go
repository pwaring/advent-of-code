package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Mapping struct {
	sourceName            string
	destinationName       string
	sourceRangeStart      int
	sourceRangeEnd        int
	destinationRangeStart int
	destinationRangeEnd   int
	rangeLength           int
	rangeDelta            int
}

func getSeeds(str string) []int {
	seeds := []int{}

	parts := strings.Split(str, ":")
	seedStrings := strings.Fields(parts[1])

	for s := range seedStrings {
		seedId, _ := strconv.Atoi(seedStrings[s])
		seeds = append(seeds, seedId)
	}

	return seeds
}

func getMappings(str string) []Mapping {
	mappings := []Mapping{}

	// Assumption: Unix line endings
	// Go doesn't support \R
	lines := strings.Split(str, "\n")

	// First line contains the source and destination names
	mappingNamesRegex := regexp.MustCompile(`^(?P<source>[a-z]+)-to-(?P<destination>[a-z]+) map:`)
	mappingMatches := mappingNamesRegex.FindStringSubmatch(lines[0])
	sourceIndex := mappingNamesRegex.SubexpIndex("source")
	destinationIndex := mappingNamesRegex.SubexpIndex("destination")
	sourceName := mappingMatches[sourceIndex]
	destinationName := mappingMatches[destinationIndex]

	for lineIndex := 1; lineIndex < len(lines); lineIndex++ {
		if len(lines[lineIndex]) > 0 {
			lineParts := strings.Fields(lines[lineIndex])
			mapping := Mapping{}
			mapping.sourceName = sourceName
			mapping.destinationName = destinationName
			mapping.destinationRangeStart, _ = strconv.Atoi(lineParts[0])
			mapping.sourceRangeStart, _ = strconv.Atoi(lineParts[1])
			mapping.rangeLength, _ = strconv.Atoi(lineParts[2])
			mapping.destinationRangeEnd = mapping.destinationRangeStart + (mapping.rangeLength - 1)
			mapping.sourceRangeEnd = mapping.sourceRangeStart + (mapping.rangeLength - 1)
			mapping.rangeDelta = mapping.destinationRangeStart - mapping.sourceRangeStart
			mappings = append(mappings, mapping)
		}
	}

	return mappings
}

func getSourceMappings(sourceName string, mappings []Mapping) []Mapping {
	sourceMappings := []Mapping{}

	for m := range mappings {
		if mappings[m].sourceName == sourceName {
			sourceMappings = append(sourceMappings, mappings[m])
		}
	}

	return sourceMappings
}

func mapSourceToDestinationId(sourceName string, sourceId int, destinationName string, mappings []Mapping) int {
	destinationId := sourceId
	destinationFound := false
	currentSourceName := sourceName

	for !destinationFound {
		sourceMappings := getSourceMappings(currentSourceName, mappings)

		if len(sourceMappings) > 0 {
			// Map the ID - if no match then the ID doesn't change
			for sm := range sourceMappings {
				if destinationId >= sourceMappings[sm].sourceRangeStart && destinationId <= sourceMappings[sm].sourceRangeEnd {
					destinationId += sourceMappings[sm].rangeDelta

					// Break out of the for loop once we have found a match
					break
				}
			}

			if sourceMappings[0].destinationName == destinationName {
				destinationFound = true
			} else {
				// Move to the next set of mappings
				currentSourceName = sourceMappings[0].destinationName
			}
		}
	}

	return destinationId
}

func main() {
	mappings := []Mapping{}
	inputBytes, _ := os.ReadFile(os.Stdin.Name())
	inputString := string(inputBytes)

	parts := strings.Split(inputString, "\n\n")

	// First part is the seed list
	seeds := getSeeds(parts[0])

	// Each additional part is a mapping
	for p := 1; p < len(parts); p++ {
		if len(strings.TrimSpace(parts[p])) > 0 {
			partMappings := getMappings(parts[p])
			mappings = append(mappings, partMappings...)
		}
	}

	locations := []int{}

	for s := range seeds {
		locations = append(locations, mapSourceToDestinationId("seed", seeds[s], "location", mappings))
	}

	fmt.Println(slices.Min(locations))
}
