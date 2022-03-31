package main

import "fmt"

// Liskov Substitution Principle

type ITransport interface {
	drive()
}

type ManualCar struct {
	distance int
}

type Toyota struct {
	ManualCar
	Convertibles string
}

func (t Toyota) drive() {
	//TODO implement me
	panic("implement me")
}

type Ford struct {
	ManualCar
	Convertibles string
}

func (f Ford) drive() {
	//TODO implement me
	panic("implement me")
}

type Tesla struct {
	ManualCar
	Rich string
}

type Run struct {
}

func (r Run) drive(t ITransport) {
	t.drive()
}

func main() {
	fmt.Println("Liskov Substitution Principle")
	toyota := Toyota{ManualCar{distance: 100}, "Convertibles"}
	ford := Ford{ManualCar{distance: 200}, "Convertibles"}
	tesla := Tesla{ManualCar{distance: 300}, "Rich"}
	run := Run{}
	run.drive(toyota)
	run.drive(ford)
	// error: cannot use tesla (type Tesla) as type ITransport in argument to run.drive:
	// Tesla does not implement ITransport (drive method has pointer receiver)
	// Tesla cannot be used as ManualCar
	run.drive(tesla)
}
