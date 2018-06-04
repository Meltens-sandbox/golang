// Package main provides ...
package main

import (
	"fmt"
)

type (
	// People is
	People struct {
		id   int
		name string
		sex  string
	}
	// Event is
	Event struct {
		People
		Done struct{}
		Err  error
	}
)

func main() {
	eventChan := make(chan Event, 1)

	go func() {
		// defer close(event)
		var e Event
		e.People.id = 1
		e.People.name = "Meltens"
		e.People.sex = "Men"
		eventChan <- e
	}()

	evt := <-eventChan
	if evt.Err != nil {
		fmt.Println(evt.Err)
	} else {
		fmt.Printf("%#v\n", evt)
	}
}
