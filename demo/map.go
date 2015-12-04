package main

import (
	"fmt"
)

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo

	personDB = make(map[string]PersonInfo)

	p1 := PersonInfo{"123", "to", "you"}
	personDB[p1.ID] = p1
	p2 := PersonInfo{"456", "go", "me"}
	personDB[p2.ID] = p2

	fmt.Println(personDB)
	if person, ok := personDB["456"]; ok {
		fmt.Println("person: ", person.Name)
	} else {
		fmt.Println("no")
	}

	// delete()
}
