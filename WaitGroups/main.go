package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/hiteshpatel/waitgroups/smartPrint"
)

var signals = make([]string, 2)

var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	fmt.Println("Practice wait groups")
	urls := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
	}

	for _, v := range urls {
		go getStatusCode(v)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Printf("%v", signals)
}

func getStatusCode(endpoint string) error {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		return err
	}
	st := fmt.Sprintf("%d status code for %s", res.StatusCode, endpoint)
	// fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
	smartPrint.Smartprint(st)
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()

	return nil
}
