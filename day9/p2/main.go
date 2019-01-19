package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"os"
)

const (
	inputFile = "../input.txt"
)

func main() {

	f, err := os.OpenFile(inputFile, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalf("could not open input file: %v", err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var players, points int
		_, err := fmt.Sscanf(scanner.Text(), "%d players; last marble is worth %d points", &players, &points)
		if err != nil {
			log.Fatalf("could not scan file content: %v", err)
		}

		fmt.Printf("ans: %d\n", marble(players, points*100))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func marble(players int, lastScore int) int {
	circle := &ring.Ring{Value: 0}
	scores := make([]int, players)

	for i := 1; i < lastScore; i++ {
		if i%23 == 0 {
			//  In addition, the marble 7 marbles counter-clockwise from the current marble is removed
			// shift 7 marbles and plus 1 marbles to remove, total is 8
			circle = circle.Move(-8)
			removed := circle.Unlink(1)
			v, ok := removed.Value.(int)
			if !ok {
				log.Fatalf("could not convert ring's value to int")
			}
			scores[i%players] += i + v
			// clockwise of the marble that was removed becomes the new current marble.
			circle = circle.Next()
		} else {
			// clockwise of the marble
			circle = circle.Next()
			// append new marble
			circle = circle.Link(&ring.Ring{Value: i})
			// the marble that was just placed then becomes the current marble
			// shift to last one as current marble
			circle = circle.Prev()
		}
	}
	var highestScore int
	for _, score := range scores {
		if score > highestScore {
			highestScore = score
		}
	}
	return highestScore
}
