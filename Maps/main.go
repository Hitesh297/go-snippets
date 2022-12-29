package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Test data structures")

	var mylist = make([]string, 10)

	for index := range mylist {
		mylist[index] = "list item " + strconv.Itoa(index)
	}

	for index, val := range mylist {
		println(index, val)
	}

	var mydictionary = make(map[int]string)

	for i := 0; i < 10; i++ {
		mydictionary[i] = "Value with key " + strconv.Itoa(i)

	}

	for i := 0; i < len(mydictionary); i++ {
		fmt.Printf("Value of item %v is %v \n", i, mydictionary[i])
	}
}
