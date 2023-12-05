package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day5 struct{}

func (d Day5) Part1(input string) {
	lines := strings.Split(input, "\r\n")

	var seeds []int
	var seedsToSoil [][]int
	var soilToFertilizer [][]int
	var fertilizerToWater [][]int
	var waterToLight [][]int
	var lightToTemperature [][]int
	var temperatureToHumidity [][]int
	var humidityToLocation [][]int

	var current = ""
	for i, line := range lines {
		if line == "" {
			continue
		}

		if i == 0 {
			seedsLine := strings.Split(line, ":")
			seedsStrings := strings.Split(strings.TrimSpace(seedsLine[1]), " ")
			for _, seedsString := range seedsStrings {
				seed, _ := strconv.Atoi(seedsString)
				seeds = append(seeds, seed)
			}
			continue
		}

		if line == "seed-to-soil map:" {
			current = "seed-to-soil"
			continue
		} else if line == "soil-to-fertilizer map:" {
			current = "soil-to-fertilizer"
			continue
		} else if line == "fertilizer-to-water map:" {
			current = "fertilizer-to-water"
			continue
		} else if line == "water-to-light map:" {
			current = "water-to-light"
			continue
		} else if line == "light-to-temperature map:" {
			current = "light-to-temperature"
			continue
		} else if line == "temperature-to-humidity map:" {
			current = "temperature-to-humidity"
			continue
		} else if line == "humidity-to-location map:" {
			current = "humidity-to-location"
			continue
		}

		values := strings.Split(line, " ")
		if current == "seed-to-soil" {
			var soil []int
			for _, value := range values {
				v, _ := strconv.Atoi(value)
				soil = append(soil, v)
			}
			seedsToSoil = append(seedsToSoil, soil)
		} else if current == "soil-to-fertilizer" {
			var fertilizer []int
			for _, value := range values {
				v, _ := strconv.Atoi(value)
				fertilizer = append(fertilizer, v)
			}
			soilToFertilizer = append(soilToFertilizer, fertilizer)
		} else if current == "fertilizer-to-water" {
			var water []int
			for _, value := range values {
				v, _ := strconv.Atoi(value)
				water = append(water, v)
			}
			fertilizerToWater = append(fertilizerToWater, water)
		} else if current == "water-to-light" {
			var light []int
			for _, value := range values {
				v, _ := strconv.Atoi(value)
				light = append(light, v)
			}
			waterToLight = append(waterToLight, light)
		} else if current == "light-to-temperature" {
			var temperature []int
			for _, value := range values {
				v, _ := strconv.Atoi(value)
				temperature = append(temperature, v)
			}
			lightToTemperature = append(lightToTemperature, temperature)
		} else if current == "temperature-to-humidity" {
			var humidity []int
			for _, value := range values {
				v, _ := strconv.Atoi(value)
				humidity = append(humidity, v)
			}
			temperatureToHumidity = append(temperatureToHumidity, humidity)
		} else if current == "humidity-to-location" {
			var location []int
			for _, value := range values {
				v, _ := strconv.Atoi(value)
				location = append(location, v)
			}
			humidityToLocation = append(humidityToLocation, location)
		}
	}

	var locations []int
	for _, seed := range seeds {
		soil := seed
		for _, temp := range seedsToSoil {
			if seed >= temp[1] && seed < temp[1]+temp[2] {
				soil = temp[0] + (seed - temp[1])
			}
		}

		fertilizer := soil
		for _, temp := range soilToFertilizer {
			if soil >= temp[1] && soil < temp[1]+temp[2] {
				fertilizer = temp[0] + (soil - temp[1])
			}
		}

		water := fertilizer
		for _, temp := range fertilizerToWater {
			if fertilizer >= temp[1] && fertilizer < temp[1]+temp[2] {
				water = temp[0] + (fertilizer - temp[1])
			}
		}

		light := water
		for _, temp := range waterToLight {
			if water >= temp[1] && water < temp[1]+temp[2] {
				light = temp[0] + (water - temp[1])
			}
		}

		temperature := light
		for _, temp := range lightToTemperature {
			if light >= temp[1] && light < temp[1]+temp[2] {
				temperature = temp[0] + (light - temp[1])
			}
		}

		humidity := temperature
		for _, temp := range temperatureToHumidity {
			if temperature >= temp[1] && temperature < temp[1]+temp[2] {
				humidity = temp[0] + (temperature - temp[1])
			}
		}

		location := humidity
		for _, temp := range humidityToLocation {
			if humidity >= temp[1] && humidity < temp[1]+temp[2] {
				location = temp[0] + (humidity - temp[1])
			}
		}

		locations = append(locations, location)
	}

	var smallest = locations[0]
	for _, location := range locations {
		if location < smallest {
			smallest = location
		}
	}

	fmt.Println(smallest)
}

func (d Day5) Part2(input string) {
	lines := strings.Split(input, "\r\n")

	var seeds [][]int
	var seedsToSoil [][]int
	var soilToFertilizer [][]int
	var fertilizerToWater [][]int
	var waterToLight [][]int
	var lightToTemperature [][]int
	var temperatureToHumidity [][]int
	var humidityToLocation [][]int

	var current = ""
	for i, line := range lines {
		if line == "" {
			continue
		}

		if i == 0 {
			seedsLine := strings.Split(line, ":")
			seedsStrings := strings.Split(strings.TrimSpace(seedsLine[1]), " ")
			var previous int
			for i, seedsString := range seedsStrings {
				if i%2 == 0 {
					previous, _ = strconv.Atoi(seedsString)
					continue
				} else {
					seed, _ := strconv.Atoi(seedsString)
					seeds = append(seeds, []int{previous, seed})
				}
			}
			continue
		}

		if line == "seed-to-soil map:" {
			current = "seed-to-soil"
			continue
		} else if line == "soil-to-fertilizer map:" {
			current = "soil-to-fertilizer"
			continue
		} else if line == "fertilizer-to-water map:" {
			current = "fertilizer-to-water"
			continue
		} else if line == "water-to-light map:" {
			current = "water-to-light"
			continue
		} else if line == "light-to-temperature map:" {
			current = "light-to-temperature"
			continue
		} else if line == "temperature-to-humidity map:" {
			current = "temperature-to-humidity"
			continue
		} else if line == "humidity-to-location map:" {
			current = "humidity-to-location"
			continue
		}

		values := strings.Split(line, " ")
		if current == "seed-to-soil" {
			var soil []int
			for _, value := range values {
				v, _ := strconv.Atoi(value)
				soil = append(soil, v)
			}
			seedsToSoil = append(seedsToSoil, soil)
		} else if current == "soil-to-fertilizer" {
			var fertilizer []int
			for _, value := range values {
				v, _ := strconv.Atoi(value)
				fertilizer = append(fertilizer, v)
			}
			soilToFertilizer = append(soilToFertilizer, fertilizer)
		} else if current == "fertilizer-to-water" {
			var water []int
			for _, value := range values {
				v, _ := strconv.Atoi(value)
				water = append(water, v)
			}
			fertilizerToWater = append(fertilizerToWater, water)
		} else if current == "water-to-light" {
			var light []int
			for _, value := range values {
				v, _ := strconv.Atoi(value)
				light = append(light, v)
			}
			waterToLight = append(waterToLight, light)
		} else if current == "light-to-temperature" {
			var temperature []int
			for _, value := range values {
				v, _ := strconv.Atoi(value)
				temperature = append(temperature, v)
			}
			lightToTemperature = append(lightToTemperature, temperature)
		} else if current == "temperature-to-humidity" {
			var humidity []int
			for _, value := range values {
				v, _ := strconv.Atoi(value)
				humidity = append(humidity, v)
			}
			temperatureToHumidity = append(temperatureToHumidity, humidity)
		} else if current == "humidity-to-location" {
			var location []int
			for _, value := range values {
				v, _ := strconv.Atoi(value)
				location = append(location, v)
			}
			humidityToLocation = append(humidityToLocation, location)
		}
	}

	sort.Slice(seedsToSoil, func(i, j int) bool {
		return seedsToSoil[i][1] < seedsToSoil[j][1]
	})
	sort.Slice(soilToFertilizer, func(i, j int) bool {
		return soilToFertilizer[i][1] < soilToFertilizer[j][1]
	})
	sort.Slice(fertilizerToWater, func(i, j int) bool {
		return fertilizerToWater[i][1] < fertilizerToWater[j][1]
	})
	sort.Slice(waterToLight, func(i, j int) bool {
		return waterToLight[i][1] < waterToLight[j][1]
	})
	sort.Slice(lightToTemperature, func(i, j int) bool {
		return lightToTemperature[i][1] < lightToTemperature[j][1]
	})
	sort.Slice(temperatureToHumidity, func(i, j int) bool {
		return temperatureToHumidity[i][1] < temperatureToHumidity[j][1]
	})
	sort.Slice(humidityToLocation, func(i, j int) bool {
		return humidityToLocation[i][1] < humidityToLocation[j][1]
	})

	var seedToSoilRanges [][][]int
	for _, seed := range seeds {
		answer := d.Do(seed, seedsToSoil)
		seedToSoilRanges = append(seedToSoilRanges, answer)
	}

	var soilToFertilizerRanges [][][]int
	for _, seedToSoilRange := range seedToSoilRanges {
		for _, seed := range seedToSoilRange {
			answer := d.Do(seed, soilToFertilizer)
			soilToFertilizerRanges = append(soilToFertilizerRanges, answer)
		}
	}

	var fertilizerToWaterRanges [][][]int
	for _, soilToFertilizerRange := range soilToFertilizerRanges {
		for _, seed := range soilToFertilizerRange {
			answer := d.Do(seed, fertilizerToWater)
			fertilizerToWaterRanges = append(fertilizerToWaterRanges, answer)
		}
	}

	var waterToLightRanges [][][]int
	for _, fertilizerToWaterRange := range fertilizerToWaterRanges {
		for _, seed := range fertilizerToWaterRange {
			answer := d.Do(seed, waterToLight)
			waterToLightRanges = append(waterToLightRanges, answer)
		}
	}

	var lightToTemperatureRanges [][][]int
	for _, waterToLightRange := range waterToLightRanges {
		for _, seed := range waterToLightRange {
			answer := d.Do(seed, lightToTemperature)
			lightToTemperatureRanges = append(lightToTemperatureRanges, answer)
		}
	}

	var temperatureToHumidityRanges [][][]int
	for _, lightToTemperatureRange := range lightToTemperatureRanges {
		for _, seed := range lightToTemperatureRange {
			answer := d.Do(seed, temperatureToHumidity)
			temperatureToHumidityRanges = append(temperatureToHumidityRanges, answer)
		}
	}

	var humidityToLocationRanges [][][]int
	for _, temperatureToHumidityRange := range temperatureToHumidityRanges {
		for _, seed := range temperatureToHumidityRange {
			answer := d.Do(seed, humidityToLocation)
			humidityToLocationRanges = append(humidityToLocationRanges, answer)
		}
	}

	var locations []int
	for _, humidityToLocationRange := range humidityToLocationRanges {
		for _, seed := range humidityToLocationRange {
			locations = append(locations, seed[0])
		}
	}

	var smallest = locations[0]
	for _, location := range locations {
		if location < smallest {
			smallest = location
		}
	}

	fmt.Println(smallest)
}

func (d Day5) Do(seed []int, seedsToSoil [][]int) [][]int {
	seedStart := seed[0]
	seedEnd := seed[0] + seed[1]

	var index int
	var personalSeedsToSoil [][]int
	for seedStart != seedEnd {
		if index >= len(seedsToSoil) {
			personalSeedsToSoil = append(personalSeedsToSoil, []int{
				seedStart,
				seedEnd - seedStart,
			})
			seedStart = seedEnd
			continue
		}

		soil := seedsToSoil[index]

		destination := soil[0]
		source := soil[1]
		sourceRange := soil[2]

		if seedStart > source+sourceRange || seedEnd < source {
			index++
			continue
		}

		if source < seedStart {
			destination += seedStart - source
			sourceRange += source - seedStart
			source = seedStart
		}

		if source+sourceRange > seedEnd {
			sourceRange = sourceRange - ((source + sourceRange) - seedEnd)
		}

		if source != seedStart {
			personalSeedsToSoil = append(personalSeedsToSoil, []int{
				seedStart,
				source - seedStart,
			})
		}

		personalSeedsToSoil = append(personalSeedsToSoil, []int{
			destination,
			sourceRange,
		})

		seedStart = source + sourceRange
		index++
	}

	return personalSeedsToSoil
}
