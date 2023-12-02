package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
	"strings"
	"strconv"
)


func min_cube_sum(line string)int {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	cubes := map[string]int { "red": 0, "green": 0, "blue": 0}
	arr := strings.FieldsFunc(line, f)
	for i := 0; i < len(arr) - 1; i+=2 {
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return -1
		}
		if cubes[arr[i+1]] < num {
			cubes[arr[i+1]] = num
		}
	}
	fmt.Println(cubes)
	return (cubes["red"] * cubes["blue"] * cubes["green"]) 
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
		result += min_cube_sum(line[p:])
	}	
	fmt.Println(result)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
