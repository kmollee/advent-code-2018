package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
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

	minLen := -1
	var minChar byte
	const diff = 'a' - 'A'
	for c, used := range units(polymers) {
		if !used {
			continue
		}
		p := trim(polymers, []string{string(c), string(c - diff)})
		r := react(p)
		if len(r) < minLen || minLen == -1 {
			minLen = len(r)
			minChar = c
		}
	}

	fmt.Println(minLen, string(minChar))
}

func trim(s string, replaceChars []string) string {
	for _, c := range replaceChars {
		s = strings.Replace(s, c, "", -1)
	}
	return s
}

func units(s string) map[byte]bool {

	s = strings.ToLower(s)
	counter := make(map[byte]bool)

	for i := 0; i < len(s); i++ {

		counter[s[i]] = true
	}

	return counter
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
