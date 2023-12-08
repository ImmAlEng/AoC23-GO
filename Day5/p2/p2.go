package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
	"strings"
	"strconv"
	"sort"
)

func biggest_nb(l_map [][][]int, seeds []int)int {
	biggest := 0;
	for i := 0; i < len(seeds) - 1; i+=2 {
		if biggest < seeds[i] + seeds[i + 1] {
			biggest = seeds[i] + seeds[i + 1]
		}
	}
	for k := 0; k < len(l_map); k++ {
		for i := 0; i < len(l_map[k]);i++ {
			if biggest < l_map[k][i][0] + l_map[k][i][2] {
				biggest =  l_map[k][i][0] + l_map[k][i][2] 
			}
		}
	}
	return biggest
}

func insert_slice(triplet []int, l_map [][][]int, k int) {
	i := 0
	for ; i < len(l_map[k]);i++ {
		if triplet[1] < l_map[k][i][1] {
			l_map[k] = append(l_map[k][:i], append([][]int{triplet}, l_map[k][i:]...)...)
			return
		}
	}
	l_map[k] = append(l_map[k], triplet)
}

func fill_maps(l_map [][][]int) {

	for k := 0; k < len(l_map); k++{
		if l_map[k][0][1] != 0 {
			insert_slice([]int{0, 0, l_map[k][0][1]}, l_map, k)
		}
		for j := 0; j < len(l_map[k]) - 1; j++{
			if l_map[k][j][1] + l_map[k][j][2] < l_map[k][j+1][1] {
				insert_slice([]int{l_map[k][j][1] + l_map[k][j][2], l_map[k][j][1] + l_map[k][j][2], l_map[k][j+1][1] - l_map[k][j][1] + l_map[k][j][2]}, l_map, k)
				fill_maps(l_map)
			}
		}
	}
	for k := 0; k < len(l_map); k++{
		for j := 0; j < len(l_map[k]); j++{
			if j + 1 == len(l_map[k]) {
				insert_slice([]int{l_map[k][j][1] + l_map[k][j][2], l_map[k][j][1] + l_map[k][j][2], (5000000000 - l_map[k][j][1] - l_map[k][j][2])},l_map, k)  
				j++
			}
		}
	}
}

func intersection(triplet []int, seeds []int) []int {
	var inter []int
	beginTriplet, endTriplet := triplet[1], triplet[1]+triplet[2]
	beginSeeds, endSeeds := seeds[0], seeds[0]+seeds[1]

	// max && min
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	if beginSeeds < endTriplet && endSeeds > beginTriplet {
		interBegin := max(beginTriplet, beginSeeds)
		interEnd := min(endTriplet, endSeeds)
		inter = []int{interBegin, interEnd - interBegin}
	}
	return inter
}

func remainder(inter []int, seeds []int) []int {
	var rem []int
	if inter == nil || seeds == nil {
		return nil
	}
	if len(inter) == 2 {
		beginRem := inter[1] + inter[0] 
		endRem := seeds[0] + seeds[1]  

		if endRem > beginRem {
			rem = []int{beginRem, endRem - beginRem}
		}
	}
	return rem
}


func min_elem(sl []int)int {
	if (len(sl) == 0){
		return -1
	}
	low := sl[0]
	for i := 1; i < len(sl); i++ {
		if low > sl[i] {
			low = sl[i]
		}	
	}
	fmt.Println(sl)
	return low
}

func get_nb(line string)[]int {
	seperator := strings.IndexAny(line, ":")
	var sl1 []string = strings.Fields(line[seperator+1:])
	var sl_nb []int
	for i := 0; i < len(sl1); i++ {
		number, err := strconv.Atoi(sl1[i])                 
		if err != nil {
			fmt.Println("Error", err)
        }
		sl_nb = append(sl_nb, number)
	}
	return sl_nb
}



func get_orderd_lmaps(sl []string)[][][]int{
	var locations [][]int
	var l_maps [][][]int
	for j := 0; j < len(sl); j++ {
		locations = nil
		for ; j < len(sl) && unicode.IsNumber(rune(sl[j][0])); j++ {
			sl_nb := get_nb(sl[j])
			locations = append(locations, sl_nb)
		
		}
		if locations == nil {
			continue
		}
		sort.Slice(locations, func(i, j int) bool {
		return locations[i][1] < locations[j][1]
		})
		l_maps = append(l_maps, locations)
	}
	return l_maps
}

func get_lowest_location(lowest int, k int, l_map[][][]int, seeds []int)int{
	var inter []int
	i := 0
	var nb int
	for ; i < len(l_map[k]); i++ {
		inter = intersection(l_map[k][i], seeds)
		if inter == nil {
			continue
		}
		seeds = remainder(inter, seeds)
		inter[0] += l_map[k][i][0] - l_map[k][i][1]
		if k == len(l_map) - 1 && lowest > inter[0] {
			lowest = inter[0]
		}
		if k < (len(l_map) - 1) {
			nb = get_lowest_location(lowest, k + 1, l_map, inter)
			if lowest > nb {
				lowest = nb
			}
		}
		if seeds == nil {
			break
		}
	}
	if seeds != nil {
		seeds[0] += l_map[k][i][0] - l_map[k][i][1]
		if k < (len(l_map) - 1) {
			nb = get_lowest_location(lowest, k + 1, l_map, seeds)
		} else {
			nb = seeds[0]
		}
		if lowest > nb {
			lowest = nb
		}
	}
	return lowest
}


func lowest_location(l_map [][][]int, seeds []int)int {
	var location []int
	for i := 0; i < (len(seeds) - 1);i+=2 {
		seed := []int{seeds[i], seeds[i+1]}
		location = append(location, get_lowest_location(5000000000, 0, l_map, seed))
	}
	return min_elem(location)
}


func main(){
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening File")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := 0
	var sl []string
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		sl = append(sl, line)
		i++	
	}
	seeds := get_nb(sl[0])
	fmt.Println(seeds)
	l_maps := get_orderd_lmaps(sl)
	fill_maps(l_maps)
	for k := 0; k < len(l_maps); k++ {
		fmt.Println("Map:", k)
		for i := 0; i < len(l_maps[k]); i++ {
			fmt.Println(l_maps[k][i])
		}
	}
	fmt.Println(biggest_nb(l_maps, seeds))
	result = lowest_location(l_maps, seeds)
	fmt.Println(result)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
