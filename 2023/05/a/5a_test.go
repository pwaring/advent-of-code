package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSeeds(t *testing.T) {
	seeds := []int{}
	seeds = append(seeds, 79)
	seeds = append(seeds, 14)
	seeds = append(seeds, 55)
	seeds = append(seeds, 13)

	assert.Equal(t, seeds, getSeeds("seeds: 79 14 55 13"))
}

func TestGetMappings(t *testing.T) {
	mappingStr := "seed-to-soil map:\n"
	mappingStr += "50 98 2\n"
	mappingStr += "52 50 48"

	mappings := []Mapping{}
	mappings = append(mappings, Mapping{
		sourceName:            "seed",
		destinationName:       "soil",
		sourceRangeStart:      98,
		sourceRangeEnd:        99,
		destinationRangeStart: 50,
		destinationRangeEnd:   51,
		rangeLength:           2,
		rangeDelta:            -48,
	})
	mappings = append(mappings, Mapping{
		sourceName:            "seed",
		destinationName:       "soil",
		sourceRangeStart:      50,
		sourceRangeEnd:        97,
		destinationRangeStart: 52,
		destinationRangeEnd:   99,
		rangeLength:           48,
		rangeDelta:            2,
	})
	assert.Equal(t, mappings, getMappings(mappingStr))

	mappingStr = "soil-to-fertilizer map::\n"
	mappingStr += "0 15 37\n"
	mappingStr += "37 52 2\n"
	mappingStr += "39 0 15"

	mappings = []Mapping{}
	mappings = append(mappings, Mapping{
		sourceName:            "soil",
		destinationName:       "fertilizer",
		sourceRangeStart:      15,
		sourceRangeEnd:        51,
		destinationRangeStart: 0,
		destinationRangeEnd:   36,
		rangeLength:           37,
		rangeDelta:            -15,
	})
	mappings = append(mappings, Mapping{
		sourceName:            "soil",
		destinationName:       "fertilizer",
		sourceRangeStart:      52,
		sourceRangeEnd:        53,
		destinationRangeStart: 37,
		destinationRangeEnd:   38,
		rangeLength:           2,
		rangeDelta:            -15,
	})
	mappings = append(mappings, Mapping{
		sourceName:            "soil",
		destinationName:       "fertilizer",
		sourceRangeStart:      0,
		sourceRangeEnd:        14,
		destinationRangeStart: 39,
		destinationRangeEnd:   53,
		rangeLength:           15,
		rangeDelta:            39,
	})
	assert.Equal(t, mappings, getMappings(mappingStr))
}

func TestMapSourceToDestinationId(t *testing.T) {
	mappings := []Mapping{}

	seedSoilStr := "seed-to-soil map:\n"
	seedSoilStr += "50 98 2\n"
	seedSoilStr += "52 50 48"
	seedSoilMappings := getMappings(seedSoilStr)
	mappings = append(mappings, seedSoilMappings...)

	soilFertilizerStr := "soil-to-fertilizer map:\n"
	soilFertilizerStr += "0 15 37\n"
	soilFertilizerStr += "37 52 2\n"
	soilFertilizerStr += "39 0 15"
	soilFertilizerMappings := getMappings(soilFertilizerStr)
	mappings = append(mappings, soilFertilizerMappings...)

	fertilizerWaterStr := "fertilizer-to-water map:\n"
	fertilizerWaterStr += "49 53 8\n"
	fertilizerWaterStr += "0 11 42\n"
	fertilizerWaterStr += "42 0 7\n"
	fertilizerWaterStr += "57 7 4"
	fertilizerWaterMappings := getMappings(fertilizerWaterStr)
	mappings = append(mappings, fertilizerWaterMappings...)

	waterLightStr := "water-to-light map:\n"
	waterLightStr += "88 18 7\n"
	waterLightStr += "18 25 70"
	waterLightMappings := getMappings(waterLightStr)
	mappings = append(mappings, waterLightMappings...)

	lightTemperatureStr := "light-to-temperature map:\n"
	lightTemperatureStr += "45 77 23\n"
	lightTemperatureStr += "81 45 19\n"
	lightTemperatureStr += "68 64 13"
	lightTemperatureMappings := getMappings(lightTemperatureStr)
	mappings = append(mappings, lightTemperatureMappings...)

	temperatureHumidityStr := "temperature-to-humidity map:\n"
	temperatureHumidityStr += "0 69 1\n"
	temperatureHumidityStr += "1 0 69"
	temperatureHumidityMappings := getMappings(temperatureHumidityStr)
	mappings = append(mappings, temperatureHumidityMappings...)

	humidityLocationStr := "humidity-to-location map:\n"
	humidityLocationStr += "60 56 37\n"
	humidityLocationStr += "56 93 4"
	humidityLocationMappings := getMappings(humidityLocationStr)
	mappings = append(mappings, humidityLocationMappings...)

	assert.Equal(t, seedSoilMappings, getSourceMappings("seed", mappings))
	assert.Equal(t, soilFertilizerMappings, getSourceMappings("soil", mappings))
	assert.Equal(t, fertilizerWaterMappings, getSourceMappings("fertilizer", mappings))
	assert.Equal(t, waterLightMappings, getSourceMappings("water", mappings))
	assert.Equal(t, lightTemperatureMappings, getSourceMappings("light", mappings))
	assert.Equal(t, temperatureHumidityMappings, getSourceMappings("temperature", mappings))
	assert.Equal(t, humidityLocationMappings, getSourceMappings("humidity", mappings))

	// One step at a time
	assert.Equal(t, 81, mapSourceToDestinationId("seed", 79, "soil", mappings))
	assert.Equal(t, 81, mapSourceToDestinationId("soil", 81, "fertilizer", mappings))
	assert.Equal(t, 81, mapSourceToDestinationId("fertilizer", 81, "water", mappings))
	assert.Equal(t, 74, mapSourceToDestinationId("water", 81, "light", mappings))
	assert.Equal(t, 78, mapSourceToDestinationId("light", 74, "temperature", mappings))
	assert.Equal(t, 78, mapSourceToDestinationId("temperature", 78, "humidity", mappings))
	assert.Equal(t, 82, mapSourceToDestinationId("humidity", 78, "location", mappings))

	// Two steps at a time
	assert.Equal(t, 81, mapSourceToDestinationId("seed", 79, "fertilizer", mappings))

	// Three steps at a time
	assert.Equal(t, 81, mapSourceToDestinationId("seed", 79, "water", mappings))

	// Four steps at a time
	assert.Equal(t, 74, mapSourceToDestinationId("seed", 79, "light", mappings))

	// Five steps at a time
	assert.Equal(t, 78, mapSourceToDestinationId("seed", 79, "temperature", mappings))

	// Six steps at a time
	assert.Equal(t, 78, mapSourceToDestinationId("seed", 79, "humidity", mappings))

	// Seven steps at a time
	assert.Equal(t, 82, mapSourceToDestinationId("seed", 79, "location", mappings))
	assert.Equal(t, 43, mapSourceToDestinationId("seed", 14, "location", mappings))
	assert.Equal(t, 86, mapSourceToDestinationId("seed", 55, "location", mappings))
	assert.Equal(t, 35, mapSourceToDestinationId("seed", 13, "location", mappings))
}
