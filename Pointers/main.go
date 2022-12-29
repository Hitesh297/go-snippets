package main

import "fmt"

func main() {
	fmt.Println("test pointers")

	testString := 10
	p := &testString

	fmt.Println(testString, p)
	fmt.Printf("%v\n", *p)
}
