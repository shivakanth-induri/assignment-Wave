package assignment

import (
	"math"
)

func MaxNetworkQuality(towers [][]int, radius int) []int {
	maxQuality := 0
	var maxCoord []int
	for x := 0; x <= 50; x++ {
		for y := 0; y <= 50; y++ {
			networkQuality := 0
			for _, tower := range towers {
				dist := math.Sqrt(math.Pow(float64(tower[0]-x), 2) + math.Pow(float64(tower[1]-y), 2))
				if dist <= float64(radius) {
					signalQuality := tower[2] / int(1+dist)
					networkQuality += int(signalQuality)
				}
			}
			if networkQuality > maxQuality {
				maxQuality = networkQuality
				maxCoord = []int{x, y}
			}
		}
	}
	return maxCoord
}
