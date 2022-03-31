package main

import "fmt"

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func calculate(shapes ...interface{}) float64 {
	var area float64 = 0
	for _, v := range shapes {
		switch v.(type) {
		case Circle:
			area += 3.14 * v.(Circle).Radius * v.(Circle).Radius
		case Rectangle:
			area += v.(Rectangle).Width * v.(Rectangle).Height
		}
	}
	return area
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 10, Height: 5}
	fmt.Println(calculate(c, r))
}
