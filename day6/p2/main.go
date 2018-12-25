package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type pos struct {
	x, y int
}

func (p *pos) distance(other *pos) int {
	return abs(p.x-other.x) + abs(p.y-other.y)
}
func (p *pos) totalDistance(ps []*pos) int {
	total := 0
	for _, pp := range ps {
		total += p.distance(pp)
	}
	return total
}

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	lines := strings.Split(string(data), "\n")

	var bigP pos
	ps := []*pos{}

	for _, line := range lines {
		// empty line, skip
		if len(line) == 0 {
			continue
		}

		var p pos
		_, err := fmt.Sscanf(line, "%d, %d", &p.x, &p.y)
		if err != nil {
			log.Fatalf("could not scan line in file: %v", err)
		}

		if p.x > bigP.x {
			bigP.x = p.x
		}
		if p.y > bigP.y {
			bigP.y = p.y
		}

		ps = append(ps, &p)

	}

	// grid is larger 1 than biggest point
	gridWidth, gridHeight := bigP.x+1, bigP.y+1
	grid := make([][]int, gridWidth)
	for i := range grid {
		grid[i] = make([]int, gridHeight)
	}

	maxDistance := 10000

	safePos := []*pos{}
	for i := 0; i < gridWidth; i++ {
		for j := 0; j < gridHeight; j++ {
			p := &pos{i, j}
			if p.totalDistance(ps) < maxDistance {
				safePos = append(safePos, p)
			}
		}
	}

	fmt.Println(len(safePos))

}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
