package main

import (
	"fmt"
	"sort"
)

type Triangle struct {
	base, height float32
}

type Square struct {
	lenght float32
}

type Rectangle struct {
	length, breadth float32
}

func (triangle Triangle) Area() float32 {
	return 0.5 + triangle.base*triangle.height
}

func (square Square) Area() float32 {
	return square.lenght * square.lenght
}

func (rectangle Rectangle) Area() float32 {
	return rectangle.length * rectangle.breadth
}

type Area interface {
	Area() float32
}

type Human struct {
	name string
	age  int
}

type AgeFactor []Human

func (a AgeFactor) Len() int {
	return len(a)
}

func (a AgeFactor) Less(i int, j int) bool {
	return a[i].age < a[j].age
}

func (a AgeFactor) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	fmt.Println("Practice Interface")

	triangleObject := Triangle{base: 12, height: 12}
	squareObject := Square{lenght: 12}
	rectangelObject := Rectangle{length: 12, breadth: 12}

	var areaBase Area
	areaBase = triangleObject
	fmt.Println("Area of Triangle = ", areaBase.Area())
	convertedTriangleObject := areaBase.(Triangle)
	fmt.Printf("%v\n", convertedTriangleObject)
	areaBase = squareObject
	fmt.Println("Area of Square = ", areaBase.Area())
	convertedSquareObject := areaBase.(Square)
	fmt.Printf("%v\n", convertedSquareObject)
	areaBase = rectangelObject
	fmt.Println("Area of Rectangle = ", areaBase.Area())
	convertedRectangleObject := areaBase.(Rectangle)
	fmt.Printf("%v\n", convertedRectangleObject)
	fmt.Println("Test Assertion")

	//check type of interface
	switch v := areaBase.(type) {
	case Square:
		println("Type is Square")
	case Rectangle:
		println("Type is Rectangle")
	case Triangle:
		println("Type is Triangle")
	default:
		fmt.Printf("%v\n", v)
	}

	slice1 := []int{1, 2}
	slice2 := []int{3, 4}
	slice3 := slice1
	slice1 = slice2
	fmt.Println(slice1, slice2, slice3)

	// slice1 := []int{1, 2}
	// slice2 := []int{3, 4}
	// slice3 := slice1
	// copy(slice1, slice2)
	// fmt.Println(slice1, slice2, slice3)

	audience := []Human{
		{"Alice", 35},
		{"Bob", 45},
		{"James", 25},
	}

	sort.Sort(AgeFactor(audience))
	fmt.Println(audience)

}
