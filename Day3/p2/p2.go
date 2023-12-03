package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
//	"strings"
	"strconv"
)

func in_bounds(arr [140]string, x int, y int) bool {
    if x < 0 || y < 0 {
        return false
    }
    if y >= len(arr) {
        return false
    }
    if x >= len(arr[y]) {
        return false
    }
    return true
}


func get_slice(line string, x int)string {
	i := x
	for i < len(line) && unicode.IsNumber(rune(line[i])) {
		i++
	}
	return line[x:i]
}

func gear_is_valid(arr [140]string, x int, y int) bool {
	number := 0
	x2 := x + 1
	for y1 := y - 1; y1 <= y + 1; y1++ {
		for x1 := x - 1; x1 <= x2; x1++ {
			if in_bounds(arr, x1, y1) {
				if unicode.IsNumber(rune(arr[y1][x1])){ 
					if x1 == x - 1 || !unicode.IsNumber(rune(arr[y1][x1 - 1])) {
						number++	
					}
				}
			}
		}
	}
	if number == 2 {
		return true
	}
	return false
}

func get_nb(line string, x int)int{
	for x >= 0 && unicode.IsNumber(rune(line[x])){
		x--
	}	
	x++
	i := x
	for x < len(line) && unicode.IsNumber(rune(line[x])){
		x++
	}
	num, err := strconv.Atoi(line[i:x])
	if err != nil {
		fmt.Println("Error", err)
		return -1
	}
	return num
}

func gear_ratio(arr [140]string, x int, y int) int {
	x1 := -1
	y1 := -1
	nb := 1
	for y1 = -1; y1 <= 1; y1++ {
		for x1 = -1; x1 <= 1; x1++ {
			if in_bounds(arr, x1+x, y1+y) && unicode.IsNumber(rune(arr[y+y1][x+x1])) {
				nb *= get_nb(arr[y+y1], x+x1)
				for x1 <= 1 && unicode.IsNumber(rune(arr[y+y1][x+x1])) {
					x1++		
				}
			}
		}
	}
	return nb 
}

func get_gears(arr [140]string)int{
	result := 0
	for y := 0; y < 140; y++ {
		for x := 0; x < len(arr[y]); x++ {
			if rune(arr[y][x]) == '*'{
				if gear_is_valid(arr, x, y){
					result += gear_ratio(arr, x, y)
				}
			}
		}
	}
	return result
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
	var arr [140]string
	i := 0
	for scanner.Scan() && i < 140 {
		arr[i] = scanner.Text()
		fmt.Println(arr[i])
		i++	
	}	
	result = get_gears(arr)
	fmt.Println(result)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
