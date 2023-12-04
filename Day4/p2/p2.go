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
//	fmt.Println(arr1, arr2)
	r := 0
	for i2 := range arr2 {
		for i1 := range arr1 {
			if arr2[i2] == arr1[i1] {
				r++
			}
		} 
	}
	return r
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
	var arr []string
	var arr2 []int
	for scanner.Scan() {
		line := scanner.Text()
		p := (strings.IndexAny(line, ":") + 1)
		arr = append(arr, line[p:])
		arr2 = append(arr2, 1)
	}	
	for i1 := 0; i1 < len(arr); i1++ {
		for i2 := arr2[i1]; i2 > 0; i2--{
			num1 := get_winnings(arr[i1])
			for num := num1; num > 0;num-- {
				if i1+num < len(arr2){
					arr2[i1+num]++;
				}
			}
		}
	}
	for i3 := 0; i3 < len(arr2); i3++ {
		result += arr2[i3]
	}
	fmt.Println(result)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
