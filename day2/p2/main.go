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

	var words []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		words = append(words, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	var mostCommonLen int
	var mostCommonWord string

	for idx, i := range words {
		for _, j := range words[idx+1:] {
			common, ok := compare(i, j)
			if ok {
				if len(common) > mostCommonLen {
					mostCommonLen = len(common)
					mostCommonWord = common
				}
			}
		}
	}

	fmt.Println(mostCommonLen, mostCommonWord)

}

// compare :compare two string common word, if two word length is not equal return false
func compare(a, b string) (string, bool) {
	if len(a) != len(b) {
		return "", false
	}

	var commonLetter []byte
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			commonLetter = append(commonLetter, a[i])
		}
	}

	return string(commonLetter[:]), true
}
