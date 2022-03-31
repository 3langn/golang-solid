package main

import "fmt"

type IAnimal interface {
	Eat()
	Move()
	Swim()
	Fly()
}

type Cow struct {
	name string
}

func (c Cow) Eat() {
}

func (c Cow) Move() {
}

func (c Cow) Swim() {
}

func (c Cow) Fly() {
}

type Fish struct {
	name string
}

func (f Fish) Eat() {
}

func (f Fish) Move() {
}

func (f Fish) Swim() {
}

func (f Fish) Fly() {
	fmt.Printf("%s can't fly\n", f.name)
}

type Bird struct {
	name string
}

func (b Bird) Eat() {
}

func (b Bird) Move() {
}

func (b Bird) Swim() {
	fmt.Printf("%s can't swim\n", b.name)
}

func (b Bird) Fly() {
}

func DoSomething(animal IAnimal) {
	animal.Eat()
	animal.Move()
	animal.Swim()
	animal.Fly()
}

func main() {
	DoSomething(Cow{})
	DoSomething(Fish{})
	DoSomething(Bird{})
}
