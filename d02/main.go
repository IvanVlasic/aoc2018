package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	boxes := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		boxes = append(boxes, scanner.Text())
	}

	// part1
	ww2l, ww3l := 0, 0
	for _, box := range boxes {
		runes := make(map[rune]int, 0)
		for _, r := range box {
			runes[r]++
		}

		has2Letters, has3Letters := false, false

		for _, count := range runes {
			if count == 2 && !has2Letters {
				ww2l++
				has2Letters = true
			}
			if count == 3 && !has3Letters {
				ww3l++
				has3Letters = true
			}
		}
	}
	fmt.Printf("Checksum is %d.\n", ww2l*ww3l)

	// part2
	idLen := len(boxes[0])

	for i := 0; i < idLen; i++ {
		words := make(map[string]struct{})
		for _, word := range boxes {
			w := word[:i] + word[i+1:]
			if _, ok := words[w]; ok {
				fmt.Printf("Common letters between two boxes are %s.\n", w)
			}
			words[w] = struct{}{}
		}
	}

}
