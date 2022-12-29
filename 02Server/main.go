package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Requesting...")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(databytes))

	response.Body.Close()
}
