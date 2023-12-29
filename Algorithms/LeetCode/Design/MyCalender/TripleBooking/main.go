package main

/*
 * A triple booking happens when three events have some non-empty intersection
 * (i.e., some moment is common to all the three events.).
 */

import (
	"fmt"
	"sort"
)

type Event struct {
	time, count int
}

type MyCalendarTwo struct {
	events []Event
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (cal *MyCalendarTwo) Book(start, end int) bool {
	// keep a copy of all events before adding current event to revert if needed
	cachedEvents := append(make([]Event, 0, len(cal.events)), cal.events...)
	cal.events = append(cal.events, Event{start, 1}, Event{end, -1})

	// Sort the events based on time
	sort.Slice(cal.events, func(i, j int) bool {
		if cal.events[i].time == cal.events[j].time {
			// if times are equal, prioritize end events (count =  -1) over start events
			return cal.events[i].count < cal.events[j].count
		}
		return cal.events[i].time < cal.events[j].time
	})

	count := 0

	for _, event := range cal.events {
		count += event.count

		if count > 2 {
			// remove the current event added at start
			cal.events = cachedEvents
			return false
		}
	}
	return true
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Book(10, 20)) // true
	fmt.Println(obj.Book(50, 60)) // true
	fmt.Println(obj.Book(10, 40)) // true
	fmt.Println(obj.Book(5, 15))  // false
	fmt.Println(obj.Book(5, 10))  // true
	fmt.Println(obj.Book(25, 55)) // true
}
