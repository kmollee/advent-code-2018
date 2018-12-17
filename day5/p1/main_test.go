package main

import (
	"fmt"
	"log"
	"testing"
)

func TestOpposite(t *testing.T) {
	tt := []struct {
		a, b byte
		res  bool
	}{
		{'a', 'A', true},
		{'A', 'a', true},
		{'a', 'a', false},
		{'a', 'b', false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%c-%c", tc.a, tc.b), func(t *testing.T) {
			if opposite(tc.a, tc.b) != tc.res {
				t.Fatalf("%c compare %c should be %v, but got %v", tc.a, tc.b, tc.res, tc.a == tc.b)
			}
		})
	}
}

func TestStep(t *testing.T) {
	tt := []struct {
		in, out string
		ok      bool
	}{

		{"Aa", "", true},
		{"Ab", "Ab", false},
		{"aAbB", "bB", true},
		{"aAa", "a", true},
		{"dabAcCaCBAcCcaDA", "dabAaCBAcCcaDA", true},
	}
	for _, tc := range tt {
		t.Run(tc.in, func(t *testing.T) {
			out, ok := step(tc.in)
			if out != tc.out {
				log.Fatalf("step(%s) expect '%s', but got '%s'", tc.in, tc.out, out)
			}
			if ok != tc.ok {
				log.Fatalf("step(%s) expect '%v', but got '%v'", tc.in, tc.ok, ok)
			}
		})
	}
}
