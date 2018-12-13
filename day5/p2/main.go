package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	data = bytes.TrimSuffix(data, []byte("\n"))

	polymers := string(data)

	var d byte

	shortestpoly := len(data)
	var removeChar rune

	for c := 'A'; c <= 'Z'; c++ {
		newpolymers := strings.Replace(polymers, string(c), "", -1)
		newpolymers = strings.Replace(newpolymers, string(c+32), "", -1)
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
		if len(newpolymers) < shortestpoly {
			shortestpoly = len(newpolymers)
			removeChar = c
		}
	}

	fmt.Println(shortestpoly, string(removeChar))
}
