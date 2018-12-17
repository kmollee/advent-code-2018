package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}
	data = bytes.TrimSuffix(data, []byte("\n"))
	r := react(string(data))
	fmt.Println(len(r))
}

func react(s string) string {
	ok := true

	for ok {
		s, ok = step(s)
	}
	return s
}

func step(s string) (string, bool) {
	for i := 0; i < len(s)-1; i++ {
		if opposite(s[i], s[i+1]) {
			return s[:i] + s[i+2:], true
		}
	}
	return s, false
}

func opposite(a, b byte) bool {
	const diff = 'a' - 'A'
	return a+diff == b || b+diff == a
}
