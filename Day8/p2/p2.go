package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)


func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = result * nums[i] / GCD(result, nums[i])
	}

	return result
}

func p1(instructions string, net map[string][2]string, node string, index int) int {
    if node[2] == 'Z' {
        return 0
    } else {
        i := func() int {
            if instructions[index%len(instructions)] == 'R' {
                return 1
            }
            return 0
        }()
        return 1 + p1(instructions, net, net[node][i], index+1)
    }
}

func p2(instructions string,  net map[string][2]string)int {
        var result []int
        for key, _ := range net {
                if key[2] == 'A' {
                        result = append(result, p1(instructions, net, key, 0))
                }
        }
        return LCM(result)
}


func main() {
    fileContent, err := ioutil.ReadFile("../input.txt")
    if err != nil {
	panic(err)
    }
    lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")
    pattern := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
    instructions := lines[0]
    net := map[string][2]string{}

    for _, line := range lines {
	matches := pattern.FindStringSubmatch(line)

    	if len(matches) == 4 {
	   key := matches[1]
	   value1 := matches[2]
	   value2 := matches[3]
	   net[key] = [2]string{value1, value2}
	}
    }
    fmt.Println("Part 2: ", p2(instructions, net))
}
