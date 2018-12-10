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

	var nums []int
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		nums = append(nums, n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	seen := map[int]bool{0: true}
loop:
	for {
		for _, n := range nums {
			sum += n
			if seen[sum] {
				fmt.Println(sum)
				break loop
			}
			seen[sum] = true
		}
	}
	fmt.Printf("total %d\n", sum)

}
