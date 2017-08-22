package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {

	// Create a new CB Rev Proxy instance
	cbRevProxy, err := New("http://localhost:4984")
	if err != nil {
		panic(fmt.Sprintf("Error creating cbrevproxy: %v", err))
	}

	// Register to handle all requests at /
	http.Handle("/", cbRevProxy)

	// Start HTTP listener (blocks)
	log.Fatal(http.ListenAndServe(":49840", nil))

}
