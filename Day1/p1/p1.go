package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
)

func getNumber(line string )int{
	nb_f := 0
	nb_l := 0
	for i := 0; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			nb_f = int(line[i] - '0')
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			nb_l = int(line[i] - '0')
			break
		}
	}
	return (nb_f * 10) + nb_l
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
		result += getNumber(line)
	}	
	fmt.Println(result)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
