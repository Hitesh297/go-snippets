package main

import "fmt"

type Animal struct {
	animalType string
}
type IAnimal interface {
	SetType(s string)
}

func (a *Animal) SetType(s string) {
	a.animalType = s
}

// func (a *Tiger) SetType(s string) {
// 	a.animalType = s
// }

type Tiger struct {
	Animal
	name string
}

func makeAnimal() Animal {
	return Animal{animalType: "Default"}
}

func makeTiger() Tiger {
	return Tiger{name: "Bag The Tiger", Animal: Animal{animalType: "Tiger"}}
}

func ObjectExplorer(a IAnimal) {

	person1, ok := a.(*Animal)
	if ok {

		person1.SetType("Converted to Animal")
		fmt.Printf("%v \n", person1)
	}

	Tiger1, ok := a.(*Tiger)
	if ok {

		Tiger1.SetType("Converted to Tiger")
		fmt.Printf("%v \n", Tiger1)
	}
}

func main() {
	fmt.Println("Test select with channels")

	basetype := makeAnimal()

	fmt.Printf("BaseType:  %v \n", basetype)

	basetype.SetType("Lion")

	fmt.Printf("BaseType after set lion:  %v\n", basetype)

	tigertype := makeTiger()

	fmt.Printf("Tiger object : %v\n", tigertype)

	ObjectExplorer(&basetype)
	ObjectExplorer(&tigertype)

	// ch1 := make(chan string)
	// ch2 := make(chan int)

	// ch1 <- "Hello"
	// ch2 <- 29

	// for {
	// 	select {
	// 	case <-ch1:
	// 		println("Data received from channel 1")
	// 	case <-ch2:
	// 		println("Data received from channel 2")
	// 	}
	// }
}
