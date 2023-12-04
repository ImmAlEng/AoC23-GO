package main

import (
	"fmt"
	"os"
	"bufio"
	//"unicode"
	"strings"
	//"strconv"
)


func get_winnings(line string)int {
	seperator := strings.IndexAny(line, "|")
	var arr1 []string = strings.Fields(line[:seperator-1])
	var arr2 []string = strings.Fields(line[seperator+1:])
	fmt.Println(arr1, arr2)
	r := 0
	for i2 := range arr2 {
		for i1 := range arr1 {
			if arr2[i2] == arr1[i1] {
				r++
			}
		} 
	}
	if r == 0 {
		return 0
	}
	return (1 << (r - 1))
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
	for scanner.Scan() {
		line := scanner.Text()
		p := (strings.IndexAny(line, ":") + 1)
		result += get_winnings(line[p:])
	}	
	fmt.Println(result)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
