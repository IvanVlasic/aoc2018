package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	steps := make(map[string]map[string]struct{})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var s1, s2 string
		fmt.Sscanf(scanner.Text(), "Step %s must be finished before step %s can begin.", &s1, &s2)
		if _, exists := steps[s2]; exists {
			steps[s2][s1] = struct{}{}
		} else {
			steps[s2] = map[string]struct{}{s1: struct{}{}}
		}
		if _, exists := steps[s1]; !exists {
			steps[s1] = map[string]struct{}{}
		}
	}

	// copy map for part2
	stepst := make(map[string]map[string]struct{})
	for s, dep := range steps {
		stepst[s] = map[string]struct{}{}
		for d := range dep {
			stepst[s][d] = struct{}{}
		}
	}

	// part1
	order := make([]string, 0)
	for len(steps) > 0 {
		availableSteps := make([]string, 0)
		for step, dep := range steps {
			if len(dep) == 0 {
				availableSteps = append(availableSteps, step)
			}
		}

		sort.Strings(availableSteps)
		order = append(order, availableSteps[0])
		delete(steps, availableSteps[0])

		for _, dep := range steps {
			delete(dep, availableSteps[0])
		}
	}

	fmt.Printf("Order: %s.\n", strings.Join(order, ""))

	// part2
	workers := make(map[string]int)
	duration := 0
	for len(stepst) > 0 {
		availableSteps := make([]string, 0)
		for step, dep := range stepst {
			if len(dep) == 0 {
				availableSteps = append(availableSteps, step)
			}
		}

		sort.Strings(availableSteps)
		for _, as := range availableSteps {
			if len(workers) == 5 {
				break
			}
			if _, exists := workers[as]; !exists {
				workers[as] = 60 + int(rune(as[0])-'A') + 1
			}
		}

		minTime := 100
		for _, d := range workers {
			if minTime > d {
				minTime = d
			}
		}
		duration += minTime

		stepsDone := make([]string, 0)
		for s := range workers {
			workers[s] -= minTime
			if workers[s] == 0 {
				stepsDone = append(stepsDone, s)
				delete(stepst, s)
				delete(workers, s)
			}
		}

		for _, dep := range stepst {
			for _, sd := range stepsDone {
				delete(dep, sd)
			}
		}
	}

	fmt.Printf("Duration: %d.\n", duration)
}
