package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
	"strings"
	"strconv"
)

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

func lowest_location(sl []string, seeds []int)int {
	var location []int
	for i := 0; i < len(seeds);i++ {
		seed := seeds[i]
		for j := 0; j < len(sl); j++ {
			for ; j < len(sl) && unicode.IsNumber(rune(sl[j][0]));j++ {
				sl_nb := get_nb(sl[j])
				if sl_nb[1] <= seed && seed < (sl_nb[1] + sl_nb[2]) {
					seed += sl_nb[0] - sl_nb[1]
					for ; j < len(sl) && unicode.IsNumber(rune(sl[j][0]));j++ {}
				}
			}
		}
		location = append(location, seed)
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
		fmt.Println(sl[i])
		i++	
	}
	var sl_nb []int
	sl_nb = get_nb(sl[0])
	fmt.Println(sl_nb)
	result = lowest_location(sl, sl_nb)
	fmt.Println(result)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
