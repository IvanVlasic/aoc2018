package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	es := []Entry{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		t, err := time.Parse("2006-01-02 15:04", s[1:17])
		if err != nil {
			log.Fatal(err)
		}

		es = append(es, Entry{t, s[19:]})
	}

	sort.Sort(Entries(es))

	guards := map[int][]int{}
	var startsSleep time.Time
	gid := -1
	for _, e := range es {
		switch e.s {
		case "falls asleep":
			startsSleep = e.t
		case "wakes up":
			for i := startsSleep.Minute(); i < e.t.Minute(); i++ {
				guards[gid][i]++
			}
		default:
			fmt.Sscanf(e.s, "Guard #%d begins shift", &gid)
			if _, exists := guards[gid]; !exists {
				guards[gid] = make([]int, 60)
			}
		}
	}

	// part1
	asleep, maxt, idx := 0, 0, 0
	for id, minutes := range guards {
		sum, currMax, midx := 0, 0, 0
		for i, m := range minutes {
			sum += m
			if m > currMax {
				currMax, midx = m, i
			}
		}
		if sum > asleep {
			asleep, maxt, idx = sum, midx, id
		}
	}
	fmt.Printf("Strategy 1: %d\n", maxt*idx)

	// part2
	currMax := 0
	for id, minutes := range guards {
		for i, m := range minutes {
			if m > currMax {
				currMax, maxt, idx = m, i, id
			}
		}
	}
	fmt.Printf("Strategy 2: %d\n", maxt*idx)
}

type Entry struct {
	t time.Time
	s string
}

type Entries []Entry

func (es Entries) Len() int           { return len(es) }
func (es Entries) Swap(i, j int)      { es[i], es[j] = es[j], es[i] }
func (es Entries) Less(i, j int) bool { return es[i].t.Before(es[j].t) }
