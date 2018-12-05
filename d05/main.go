package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.TrimSpace(string(dat))

	// part1
	frp := react(input)
	fmt.Printf("Fully reacting len: %d.\n", len(frp))

	// part2
	minLen := len(input)
	for r := 'a'; r < 'z'; r++ {
		rs := strings.Replace(input, string(r), "", -1)
		rs = strings.Replace(rs, strings.ToUpper(string(r)), "", -1)
		res := react(rs)
		if len(res) < minLen {
			minLen = len(res)
		}
	}
	fmt.Printf("Shortest polymer: %d.\n", minLen)
}

func react(s string) string {
	i := 0
	for i < len(s)-1 {
		diff := abs(int(rune(s[i]) - rune(s[i+1])))
		if diff == 32 {
			if i == len(s)-2 {
				s = s[:len(s)-2]
				break
			} else {
				s = s[:i] + s[i+2:]
				if i > 0 {
					i--
				}
			}
		} else {
			i++
		}
	}
	return s
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
