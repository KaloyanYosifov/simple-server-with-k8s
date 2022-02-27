package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Hello world, the web server

	homepageHandler := func(writer http.ResponseWriter, req *http.Request) {
		hostname, err := os.Hostname()

		if err != nil {
			http.Error(writer, "Bad request", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(writer, "Hello, world! %s\n", hostname)
	}

	doSomethingRequest := func(writer http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			http.Error(writer, "Unsupported verb!", http.StatusNotFound)
			return
		}

		jsonData, err := ioutil.ReadAll(req.Body)

		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(writer, "can't read body", http.StatusBadRequest)
			return
		}

		fmt.Printf("Data: %v \n", string(jsonData))
		var person Person
		unmarshalErr := json.Unmarshal(jsonData, &person)

		if unmarshalErr != nil {
			fmt.Printf("Something went wrong: %v\n", err)
			return
		}

		fmt.Fprintf(writer, "Welcome: %v", person.Name)
		// io.WriteString(writer, "Hello, world!\n")
	}

	http.HandleFunc("/", homepageHandler)
	http.HandleFunc("/do-something", doSomethingRequest)
	log.Println("Listing for requests at http://localhost:7011/")
	log.Fatal(http.ListenAndServe(":7011", nil))
}
