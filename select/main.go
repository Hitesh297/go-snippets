package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("using select with channels")

	ch1 := make(chan string)
	ch2 := make(chan string)

	history := make([]string, 1)

	go func() {
		input := ""
		for input != "x" {
			fmt.Scanln(&input)
			fmt.Println("UserInput: ", input)
			history = append(history, input)
			ch1 <- input

			go func() {
				time.Sleep(time.Second * 1)
				ch2 <- input
			}()
		}

		close(ch1)
		close(ch2)
		fmt.Println("History", history)
	}()

ForLoop:
	for {
		select {
		case chl1msg, isopen := <-ch1:
			if !isopen {
				break ForLoop
			}
			fmt.Println("Channel 1:", chl1msg)

		case chl2msg, isopen := <-ch2:
			if !isopen {
				break ForLoop
			}
			fmt.Println("Channel 2:", chl2msg)

		}
	}

	fmt.Println("Program exited!")

}
