package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type box struct {
	name          string
	x, y          int
	width, height int

	subBoxes []box
	m        map[string][]string
	good     map[string]bool
}

func newBox(name string, x, y, width, height int) *box {
	return &box{
		name:     name,
		x:        x,
		y:        y,
		width:    width,
		height:   height,
		subBoxes: nil,
		m:        map[string][]string{},
		good:     map[string]bool{},
	}
}

func (b *box) claims(name string, px, py, width, height int) error {
	if px+width > b.width {
		return fmt.Errorf("x:%v with width %v out of range:%v", px, width, b.width)
	}

	if py+height > b.height {
		return fmt.Errorf("y:%v with height %v out of range:%v", py, height, b.height)
	}

	b.good[name] = true

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			coordinate := fmt.Sprintf("%dx%d", px+x, py+y)
			b.m[coordinate] = append(b.m[coordinate], name)

			if len(b.m[coordinate]) == 1 {
				continue
			}

			// if overlap
			b.good[name] = false
			for _, name := range b.m[coordinate] {
				b.good[name] = false
			}

		}
	}

	return nil
}

func (b *box) overLayAmount() int {

	amount := 0
	for _, c := range b.m {
		if len(c) > 1 {
			amount++
		}
	}
	return amount
}

func (b *box) draw() {
	for x := 0; x < b.width; x++ {
		for y := 0; y < b.height; y++ {
			coordinate := fmt.Sprintf("%dx%d", x, y)
			switch len(b.m[coordinate]) {
			case 0:
				// no one claim, print board self
				fmt.Printf("%s ", b.name)
			case 1:
				fmt.Printf("%s ", b.m[coordinate][0])
			case 2:
				// overlap
				fmt.Printf("x ")
			}

		}
		fmt.Println()
	}
}

func (b *box) goodOption() string {
	for name, c := range b.good {
		if c {
			return name
		}
	}
	return ""
}

func main() {

	width := 1000
	height := 1000

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	board := newBox(".", 0, 0, width, height)
	s := bufio.NewScanner(f)
	for s.Scan() {
		var name string
		var x, y, width, height int
		if _, err := fmt.Sscanf(s.Text(), "#%s @ %d,%d: %dx%d", &name, &x, &y, &width, &height); err != nil {
			log.Fatalf("could not scan: %v", err)
		}
		if err := board.claims(name, x, y, width, height); err != nil {
			log.Fatalf("could not add box to box: %v", err)
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	// board.draw()

	// NOTE: this is p1 anwser
	fmt.Println(board.overLayAmount())
	// NOTE: this is p2 anwser
	fmt.Println(board.goodOption())
}
