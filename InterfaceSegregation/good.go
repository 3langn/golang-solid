package main

import "fmt"

type iFish interface {
	iAnimal
	Swim()
}

type iAnimal interface {
	Eat()
}

type Shark struct {
}

func (s Shark) Eat() {
	fmt.Println("Shark eat")
}

func (s Shark) Swim() {
	fmt.Println("Shark swims")
}

type iDog interface {
	iAnimal
	Barks()
}

type Husky struct {
}

func (h Husky) Barks() {
	fmt.Println("Husky barks")
}

func (h Husky) Eat() {
	fmt.Println("nom nom nom")
}

func animalEat(a iAnimal) {
	a.Eat()
}

func dogBark(b iDog) {
	b.Barks()
}

func fishSwim(f iFish) {
	f.Swim()
}

func main() {
	var a iDog = Husky{}
	var b iFish = Shark{}

	animalEat(a)
	animalEat(b)

	dogBark(a)
	fishSwim(b)
}
