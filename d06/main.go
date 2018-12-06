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

	ps := make([]point, 0)

	scanner := bufio.NewScanner(file)
	maxx, maxy := 0, 0
	for scanner.Scan() {
		p := point{}
		fmt.Sscanf(scanner.Text(), "%d, %d", &p.x, &p.y)
		ps = append(ps, p)

		if p.x > maxx {
			maxx = p.x
		}
		if p.y > maxy {
			maxy = p.y
		}
	}

	grid := make([][]info, maxx+1)
	for i := range grid {
		grid[i] = make([]info, maxy+1)
	}

	distances := make(map[int]int)
	for x := 0; x <= maxx; x++ {
		for y := 0; y <= maxy; y++ {
			bd := maxx + maxy
			bp := -1

			for i, p := range ps {
				d := abs(p.x-x) + abs(p.y-y)
				if d == bd {
					bp = -1
				}
				if d < bd {
					bd = d
					bp = i
				}
				grid[x][y].d += d
			}
			if bp > -1 {
				grid[x][y].i = bp
				distances[bp]++
			}
		}
	}

	//remove inf
	for x := 0; x <= maxx; x++ {
		delete(distances, grid[x][0].i)
		delete(distances, grid[x][maxy].i)
	}
	for y := 0; y < maxy; y++ {
		delete(distances, grid[0][y].i)
		delete(distances, grid[maxx][y].i)
	}

	// part1
	largest := 0
	for _, dist := range distances {
		if dist > largest {
			largest = dist
		}
	}
	fmt.Printf("Largest area: %d.\n", largest)

	// part2
	inarea := 0
	for x := 0; x <= maxx; x++ {
		for y := 0; y <= maxy; y++ {
			if grid[x][y].d < 10000 {
				inarea++
			}
		}
	}
	fmt.Printf("Size of region: %d.\n", inarea)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type point struct {
	x, y int
}

type info struct {
	d, i int
}
