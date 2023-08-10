package main

import (
	// "fmt"

	"fmt"
	"log"

	"net/http"
	"time"
)

const port = ":8080"

func main() {
	

	srv := &http.Server{
		Addr:         port,
		Handler:      routes(),
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	fmt.Printf("Starting server at port %s \n", port)

	err := srv.ListenAndServe()
	log.Fatal(err)

