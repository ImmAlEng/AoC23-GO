package main

import (
	"io/ioutil"
	"fmt"
	"strconv"
	"strings"
)

func readInput(filename string) ([][]int, error) {
	fileContent, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		return nil, err	
	}
	var numbers []int
	var result [][]int
	lines := strings.Split(string(fileContent), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			fields := strings.Fields(line)
			for _, field := range fields {
				num, err := strconv.Atoi(field)
				if err != nil {
					return nil, err
				}
				numbers = append(numbers, num)
			}
			result = append(result, numbers)
			numbers = nil
		}
	}
	return result, nil
}

func recursiveSum(l []int) int {
	if sumNonZero(l) == 0 {
		return 0
	}
	var m []int
	for i := 0; i < len(l)-1; i++ {
		m = append(m, l[i+1]-l[i])
	}
	return l[len(l)-1] + recursiveSum(m)
}

func sumNonZero(l []int) int {
	count := 0
	for _, v := range l {
		if v != 0 {
			count++
		}
	}
	return count
}

func main() {
	input, err := readInput("../input.txt")
	if err != nil {
		panic(err)
	}

	result := 0
	for _, l := range input {
		result += recursiveSum(l)
	}
	fmt.Println(result)
}

