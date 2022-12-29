package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}
var ch = make(chan int)

func main() {
	fmt.Println("Practice channels")
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println(<-ch)
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		ch <- 10
	}(ch, wg)

	wg.Wait()

	//**************************************************************************************************************************************
	// go func() {
	// 	for i := 0; i < 1000; i++ {
	// 		wg.Add(1)
	// 		go func(i int, ch chan int) {
	// 			time.Sleep(time.Second)
	// 			ch <- i
	// 			wg.Done()
	// 		}(i, ch)
	// 	}
	// 	wg.Wait()
	// 	close(ch)
	// }()

	// for v := range ch {
	// 	fmt.Printf("%d \n", v)
	// }

	//***********************************************************************************************

}
