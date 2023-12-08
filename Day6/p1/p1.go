package main

import "fmt"

func calcDist(hold, left int) int {
	return hold * left
}

func p1(times, records []int) {

	var races [][]int

	for i := 0; i < len(times); i++ {
		races = append(races, []int{times[i], records[i]})
	}

	var ways []int

	for _, race := range races {
		wins := 0
		for i := 0; i <= race[0]; i++ {
			if calcDist(i, race[0]-i) > race[1] {
				wins++
			}
		}
		ways = append(ways, wins)
	}

	result := 1

	for _, val := range ways {
		result *= val
	}

	fmt.Printf("Part 1: %d\n", result)
}

func main() {
	times := []int{45, 97, 72, 95}
	records := []int{305, 1062, 1110, 1695}
	p1(times, records)
}
