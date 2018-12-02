package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nums := make([]int, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, i)
	}

	// part1
	s := 0
	for _, n := range nums {
		s += n
	}
	fmt.Printf("Resulting frequency is %d.\n", s)

	// part2
	curr := 0
	vis := make(map[int]struct{})

	for i := 0; ; i = (i + 1) % len(nums) {
		curr += nums[i]
		if _, exists := vis[curr]; exists {
			fmt.Printf("Resulting frequency is %d.\n", curr)
			break
		}
		vis[curr] = struct{}{}
	}

}
