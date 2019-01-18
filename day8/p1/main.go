package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	testFile = "../input.txt"
)

type node struct {
	nChildren int
	nMeta     int
	children  []*node
	meta      []int
}

func (n *node) sumMeta() int {
	sum := 0
	for _, m := range n.meta {
		sum += m
	}

	for _, c := range n.children {
		sum += c.sumMeta()
	}
	return sum
}

func main() {
	data, err := ioutil.ReadFile(testFile)
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	parts := strings.Fields(string(data))

	nums := make([]int, len(parts))
	for i, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			log.Fatalf("could not convert string to int: %v", err)
		}
		nums[i] = n
	}

	n, _ := newNode(nums, 0)
	fmt.Println(n.sumMeta())
}

func newNode(nums []int, start int) (*node, int) {

	pos := start
	nChildren := nums[pos]
	pos++
	nMeta := nums[pos]
	pos++

	children := make([]*node, nChildren)
	for i := 0; i < nChildren; i++ {
		c, childDistance := newNode(nums, pos)
		children[i] = c
		pos += childDistance
	}

	meta := make([]int, nMeta)
	for i := 0; i < nMeta; i++ {
		meta[i] = nums[pos]
		pos++
	}

	distance := pos - start

	return &node{
		nChildren: nChildren,
		nMeta:     nMeta,
		children:  children,
		meta:      meta,
	}, distance
}
