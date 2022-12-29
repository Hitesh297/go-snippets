package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello from Hitesh")

	var textvalue string
	fmt.Println(textvalue)
	fmt.Printf("Type is: %T \n", textvalue)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating")

	input, _ := reader.ReadString('\n')
	fmt.Println("Your rating is:", input)

}
