package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var twos, threes int
	for s.Scan() {
		var gotTwo, gotThree bool
		count := map[rune]int{}
		for _, r := range s.Text() {
			count[r]++
		}
		for _, c := range count {
			// maybe more than one have 2 or 3 letter
			if c == 2 {
				gotTwo = true
			}
			if c == 3 {
				gotThree = true
			}
		}

		if gotTwo {
			twos++
		}

		if gotThree {
			threes++
		}

	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("total %d\n", twos*threes)

}
