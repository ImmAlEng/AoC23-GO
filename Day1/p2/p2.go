package main

// p1: 54667 | p2: 54203

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
	"strings"
)

func writtenNumber(line string)int{
	arr := [9]string {"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i := 0; i < 9; i++ {
		if strings.HasPrefix(line, arr[i]){
			return (i + 1)
		}
	}
	return -1
}

func getNumber(line string )int{
	nb_f := 0
	nb_l := 0
	for i := 0; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			nb_f = int(line[i] - '0')
			break
		}
		nb_n := writtenNumber(string(line[i:]))
		if nb_n != -1 {
			nb_f = nb_n
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			nb_l = int(line[i] - '0')
			break
		}
		nb_n := writtenNumber(string(line[i:]))
		if nb_n != -1 {
			nb_l = nb_n
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
