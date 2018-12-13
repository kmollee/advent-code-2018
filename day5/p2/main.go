package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("../test.txt")
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	data = bytes.TrimSuffix(data, []byte("\n"))

	polymers := string(data)

	var d byte

	for {
		isReative := false
		for i := range polymers[:len(polymers)-1] {

			if polymers[i] > polymers[i+1] {
				d = polymers[i] - polymers[i+1]
			} else {
				d = polymers[i+1] - polymers[i]
			}

			if d == 32 {
				polymers = polymers[:i] + polymers[i+2:]
				isReative = true
				break
			}
		}
		if !isReative {
			break
		}
	}

	fmt.Println(len(polymers))
}
