package main

import (
	"fmt"
)

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

func (cat Cat) SaySomething() string {
	return "mao"
}

func (dog Dog) SaySomething() string {
	return "gou"
}

type Pet interface {
	SaySomething() string
}

func AllPetsSaySomething(pets ...Pet) {

	for _, pet := range pets {
		fmt.Println(pet.SaySomething())
	}
}

func main() {
	cat := Cat{Name: "hot-cat"}
	dog := Dog{Name: "hot-dog"}
	AllPetsSaySomething(cat, dog)

}
