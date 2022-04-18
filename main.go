package main

import (
	"fmt"
	"math"
)

//Interface in Golang is a way to look at a set of related objects or types
//Both circle/rect have an area both of shape and both have area
//Interface allows you to look at circle and rect the same with both have area
//Define behavior with similar types

//Defined shape interface
//Point is to say anything type or struct that has area method that reuturns float is of type shape
//When interface is implemented we can use an "upper type of struct"
//Circ and rect implement interface of type shape
//If method listed in interface that is not correlated into either struct, they will no longer be of type shape interface

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

//if funcs receive pointers to rect or circle make sure to pass the reciever to the slice of shapes in order to locate via(&)
func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//With interfaces you can use as any type ie:var return or map
//return in main func because has area method
func getArea(s shape) float64 {
	return s.area()
}

//create slice c1 and r1
//want to look through shapes to call area on because both circle and rect have area method
//loop through shapes and sum area
//Interface will allow this (created above on line: 14)
//Slice on line: 47 is throwing error because we do not know what type to pass to shapes that both circle and rect posses
func main() {
	c1 := circle{4.5}
	r1 := rect{5, 7}
	// shapes := []type{c1, r1}
	//plugged c1 and r1 into interface of shape and use type shape
	//use interface as type all of other behavior of circ/rect we cannot access through interface
	//only use .area of interface
	shapes := []shape{c1, r1}

	//valid because area is defined in interface when interface is type with slice
	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
}
