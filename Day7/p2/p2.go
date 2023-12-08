package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"regexp"
)

func get_joker(c map[rune]int) int {
	nb, exists := c['1']
	if exists {
		delete(c, '1')
		return nb
	} else {
		return 0
	}
}

func replace(hand string)string{
	s2 := "B1DEF"
	s1 := "TJQKA"
	//o_hand := hand
	for i := 0; i < len(s2); i++ {
		hand = strings.Replace(hand, string(s1[i]), string(s2[i]), -1) 
	}
	//fmt.Println(hand, o_hand)
	return hand
}

func getValue(hand string)string {
	value := "TUVWXYZ"
	c := map[rune]int{}
	for i := 0; i < len(hand); i++ {	
		c[rune(hand[i])]++
	}
	var p []int
	j := get_joker(c)
	for _, v := range c {
		p = append(p, v)
	}
	sort.Ints(p)
	if len(p) == 0 {
		p = []int{5}
	} else {
		p[len(p) - 1] += j
	}
	f := func(a, b []int) bool {
		return fmt.Sprint(a) == fmt.Sprint(b)
	}

	t := -1
	switch {
	case f(p, []int{1, 1, 1, 1, 1}):
		t = 0
	case f(p, []int{1, 1, 1, 2}):
		t = 1
	case f(p, []int{1, 2, 2}):
		t = 2
	case f(p, []int{1, 1, 3}):
		t = 3
	case f(p, []int{2, 3}):
		t = 4
	case f(p, []int{1, 4}):
		t = 5
	case f(p, []int{5}):
		t = 6
	}
	if t == -1 {
		panic("Something went horribly wrong")
	}
	return string(value[t]) + hand
}

func main() {
	fileContent, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")
	pattern := regexp.MustCompile(`(\w+) (\w+)`)
	hands := map[string]int{}

	for _, line := range lines {
		matches := pattern.FindStringSubmatch(line)

		if len(matches) == 3 {
			key := getValue(replace(matches[1]))
			value, err := strconv.Atoi(matches[2])
			if err != nil {
				panic("Atoi failed")
			}
			hands[key] = value
		}
	}
	var keys []string
	for k, _ := range hands {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	result := 0
	for i := 0; i < len(keys); i++{
		result += hands[keys[i]] * (i + 1)
	}
	fmt.Println(result)
	//fmt.Println(hands)
}

