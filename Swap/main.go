package main

import "fmt"

//factorial code
func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func main() {

	listObj := []byte{'a', 'b', 'c', 'c', 'b', 'a'}

	for i, j := 0, len(listObj)-1; i < j; i, j = i+1, j-1 {
		listObj[i], listObj[j] = listObj[j], listObj[i]
	}

	fmt.Printf("%c\n", listObj)

	fmt.Println(factorial(7))
}
