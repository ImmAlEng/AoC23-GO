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

func nb_is_valid(arr [140]string, x int, y int, slice string) bool {
	f := func(c rune) bool {
		return c != '.' && !unicode.IsNumber(c)
	}	
	x2 := x + len(slice)
	for y1 := y - 1; y1 <= y + 1; y1++ {
		for x1 := x - 1; x1 <= x2; x1++ {
			if in_bounds(arr, x1, y1) {
				if f(rune(arr[y1][x1])){ 
					return true
				}
			}
		}
	}
	return false
}

func get_sum(arr [140]string)int{
	result := 0
	for y := 0; y < 140; y++ {
		for x := 0; x < len(arr[y]); x++ {
			if unicode.IsNumber(rune(arr[y][x])) {
				slice := get_slice(arr[y], x)
				if nb_is_valid(arr, x, y, slice){
					number, err := strconv.Atoi(slice)
					if err != nil {
						fmt.Println("Error", err)
						return -1
					}
					result += number
				}
				x += len(slice) - 1
			}
		}
	}
	return result
}

func main(){
	file, err := os.Open("../test.txt")
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
	result = get_sum(arr)
	fmt.Println(result)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
