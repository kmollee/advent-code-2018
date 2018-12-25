package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

type pos struct {
	x, y int
}

func (p *pos) distance(other *pos) int {
	return abs(p.x-other.x) + abs(p.y-other.y)
}

func (p *pos) closest(ps []*pos) int {
	var closestIndex int
	distance := math.MaxInt32
	isInfinite := false
	for index, pp := range ps {
		// close than exist
		d := p.distance(pp)
		if d < distance {
			distance = d
			closestIndex = index
			isInfinite = false
		} else if d == distance {
			isInfinite = true
		}
	}

	if isInfinite {
		return -1
	}
	return closestIndex
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

	counted := make([]int, len(ps))
	for i := 0; i < gridWidth; i++ {
		for j := 0; j < gridHeight; j++ {
			c := (&pos{i, j}).closest(ps)
			//fmt.Println(c)

			grid[i][j] = c

			if c != -1 && counted[c] != -1 {
				// if point is on edge
				if i == bigP.x || i == 0 || j == 0 || j == bigP.y {
					counted[c] = -1
				} else {
					counted[c]++
				}
			}
		}
	}

	for j := 0; j < gridHeight; j++ {
		for i := 0; i < gridWidth; i++ {

			fmt.Printf("%2d", grid[i][j])
		}
		fmt.Println()
	}

	fmt.Printf("counted %+q\n", counted)
	biggestArea := -1
	biggestPosIndex := -1
	for i, c := range counted {
		if c > biggestArea {
			biggestArea = c
			biggestPosIndex = i
		}
	}
	fmt.Println(biggestArea, biggestPosIndex)

}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
