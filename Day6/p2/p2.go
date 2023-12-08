package main

import "fmt"

func calcDist(hold, left int) int {
	return hold * left
}

func p2(time, record int) {
    ways := 0

    for i := 0; i <= time; i++ {
        if calcDist(i, time-i) > record {
            ways++
        }
    }

    fmt.Printf("Part 2: %d\n", ways)
}


func main() {
	p2(45977295, 305106211101695)
}
