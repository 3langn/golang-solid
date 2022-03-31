package main

import "fmt"

type Shape interface {
	Area() float64
}

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

type rectangle struct {
	width, height float64
}

func (r rectangle) Area() float64 {
	return r.width * r.height
}

func Calculate(shapes ...Shape) float64 {
	var area float64
	for _, shape := range shapes {
		area += shape.Area()
	}
	return area
}

func main() {
	c := circle{radius: 5}
	r := rectangle{width: 10, height: 5}
	fmt.Println(Calculate(c, r))
}
