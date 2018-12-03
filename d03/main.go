package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const SIZE = 1000

type Claim struct {
	id, l, r, w, h int
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cells := make([]int, SIZE*SIZE)
	claims := make([]Claim, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := Claim{}
		fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &c.id, &c.l, &c.r, &c.w, &c.h)
		claims = append(claims, c)

	}

	for _, c := range claims {
		for i := 0; i < c.h; i++ {
			for j := 0; j < c.w; j++ {
				idx := SIZE*(c.r+i) + c.l + j
				cells[idx]++
			}
		}
	}

	// part1
	oc := 0
	for _, i := range cells {
		if i > 1 {
			oc++
		}
	}
	fmt.Printf("Overlapping cells: %d.\n", oc)

	// part2
claimcheck:
	for _, c := range claims {
		for i := 0; i < c.h; i++ {
			for j := 0; j < c.w; j++ {
				idx := SIZE*(c.r+i) + c.l + j
				if cells[idx] > 1 {
					continue claimcheck
				}
			}
		}
		fmt.Printf("Intact claim: %d.\n", c.id)
	}
}
