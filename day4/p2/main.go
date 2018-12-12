package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

type event struct {
	id   int
	kind eventKind
	time time.Time
}

func (e event) String() string {
	date := e.time.Format("01/02 15:04")
	switch e.kind {
	case eventStart:
		return fmt.Sprintf("[%s] Guard #%d starts", date, e.id)
	case eventAsleep:
		return fmt.Sprintf("[%s] Guard #%d fall asleep", date, e.id)
	case eventAwake:
		return fmt.Sprintf("[%s] Guard #%d awake up", date, e.id)
	default:
		return fmt.Sprintf("unknow event type: %v", e.kind)
	}
}

type eventKind byte

const (
	eventStart eventKind = iota
	eventAsleep
	eventAwake
	eventShift
)

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}
	lines := strings.Split(string(data), "\n")

	sort.Strings(lines)

	var events []event

	for _, line := range lines {

		// skip empty line
		if len(line) == 0 {
			continue
		}
		t := line
		dateEnd := strings.Index(t, "]")
		if dateEnd == -1 {
			log.Fatalf("could not index date in line: %s", line)
		}
		dateText := t[1:dateEnd]

		date, err := time.Parse("2006-01-02 15:04", dateText)
		if err != nil {
			log.Fatalf("could not parse time %v: %v", dateText, err)
		}

		e := event{time: date}
		pieces := strings.Fields(t[dateEnd+1:])
		switch pieces[0] {
		case "Guard":
			id, err := strconv.Atoi(pieces[1][1:])
			if err != nil {
				log.Fatalf("could not convert %v to int; %v", pieces[1][1:], err)
			}
			e.id = id
			e.kind = eventStart
		case "falls":
			e.id = events[len(events)-1].id
			e.kind = eventAsleep
		case "wakes":
			e.id = events[len(events)-1].id
			e.kind = eventAwake
		}

		events = append(events, e)
	}

	id, minutes := findGuard(events)
	fmt.Println(id * minutes)

}

func findGuard(events []event) (id int, minutes int) {

	c := make(map[int][]int)

	maxCount := 0
	maxMinute := 0
	maxID := 0
	for index, e := range events {
		if e.kind == eventAwake {
			for i := events[index-1].time; e.time.After(i); i = i.Add(time.Minute) {
				if _, exist := c[e.id]; !exist {
					c[e.id] = make([]int, 60)
				}

				m := i.Minute()
				c[e.id][m]++

				if c[e.id][m] > maxCount {
					maxCount = c[e.id][m]
					maxMinute = m
					maxID = e.id
				}
			}
		}
	}

	return maxID, maxMinute

}
