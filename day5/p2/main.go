package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

type result struct {
	len  int
	word rune
}

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	data = bytes.TrimSuffix(data, []byte("\n"))

	polymers := string(data)

	// var d byte

	shortestpoly := len(data)
	var removeChar rune

	poly := make(chan result)

	var wg sync.WaitGroup

	for c := 'A'; c <= 'Z'; c++ {
		wg.Add(1)
		go func(text string, r rune, p chan<- result) {
			defer wg.Done()
			newpolymers := strings.Replace(text, string(r), "", -1)
			newpolymers = strings.Replace(newpolymers, string(r+32), "", -1)
			var d byte
			for {
				isReactive := false
				for i := range newpolymers[:len(newpolymers)-1] {

					if newpolymers[i] > newpolymers[i+1] {
						d = newpolymers[i] - newpolymers[i+1]
					} else {
						d = newpolymers[i+1] - newpolymers[i]
					}

					if d == 32 {
						newpolymers = newpolymers[:i] + newpolymers[i+2:]
						isReactive = true
						break
					}
				}
				if !isReactive {
					break
				}
			}

			p <- result{len(newpolymers), r}
		}(polymers, c, poly)
	}

	go func() {
		wg.Wait()
		close(poly)
	}()

	for r := range poly {
		fmt.Printf("%v\n", r)
		if r.len < shortestpoly {
			shortestpoly = r.len
			removeChar = r.word
		}
	}

	fmt.Println(shortestpoly, string(removeChar))
}
