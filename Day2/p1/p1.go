package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
	"strings"
	"strconv"
)


func game_possible(line string)bool {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	cubes := map[string]int { "red": 0, "green": 0, "blue": 0}
	arr := strings.FieldsFunc(line, f)
	for i := 0; i < len(arr) - 1; i+=2 {
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return false
		}
		cubes[arr[i+1]] = num
		if cubes["red"] > 12 || cubes["green"] > 13 || cubes["blue"] > 14 {
			return false
		}
	}
	fmt.Println(cubes)
	return true
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
	i := 0
	for scanner.Scan() {
		i++
		line := scanner.Text()
		p := (strings.IndexAny(line, ":") + 1)
		if game_possible(line[p:]) {
			result += i
		}
	}	
	fmt.Println(result)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
