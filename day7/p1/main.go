package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const (
	instructionsFile = "../input.txt"
)

type node struct {
	name       rune
	subscribes []*node
	followed   int
}

func (n *node) String() string {
	return fmt.Sprintf("name: %c follwed %d", n.name, n.followed)
}

func newNode(name rune) *node {
	return &node{
		name:       name,
		subscribes: []*node{},
	}
}

func (n *node) subscribe(other *node) {
	n.followed++
	other.subscribes = append(other.subscribes, n)
}

func (n *node) unsubscribe() {
	for _, sn := range n.subscribes {
		sn.followed--
	}
	n.subscribes = make([]*node, 1)
}

func main() {
	f, err := os.Open(instructionsFile)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}

	nodeMap := make(map[rune]*node)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var stepDone, stepAfter rune
		_, err := fmt.Sscanf(line, "Step %c must be finished before step %c can begin.", &stepDone, &stepAfter)
		if err != nil {
			log.Fatalf("could not scan string: %v", err)
		}
		// Step C must be finished before step A can begin.

		var exist bool
		var stepDoneNode, stepAfterNode *node

		stepDoneNode, exist = nodeMap[stepDone]
		if !exist {
			log.Printf("add node: %c", stepDone)
			stepDoneNode = newNode(stepDone)
			nodeMap[stepDone] = stepDoneNode

		}
		stepAfterNode, exist = nodeMap[stepAfter]
		if !exist {
			log.Printf("add node: %c", stepAfter)
			stepAfterNode = newNode(stepAfter)
			nodeMap[stepAfter] = stepAfterNode
		}
		stepAfterNode.subscribe(stepDoneNode)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scan error: %v", err)
	}

	for {
		if len(nodeMap) == 0 {
			break
		}
		candidates := []*node{}

		for _, node := range nodeMap {
			if node.followed == 0 {
				candidates = append(candidates, node)
			}
		}
		if len(candidates) == 0 {
			log.Fatalf("can not search to canidate node")
		}

		// log.Printf("get candidates: %v with len: %d\n", candidates, len(candidates))
		var removeNode *node
		if len(candidates) == 1 {
			removeNode = candidates[0]

		} else {
			sort.Slice(candidates, func(i, j int) bool {
				return candidates[i].name < candidates[j].name
			})
			removeNode = candidates[0]
		}
		fmt.Printf("%c", removeNode.name)
		removeNode.unsubscribe()
		delete(nodeMap, removeNode.name)

	}

	fmt.Println()
}
