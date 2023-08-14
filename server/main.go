package main

import (
	"fmt"
	"log"
	"os"

	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      routes(),
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	fmt.Printf("Starting server at port %s \n", os.Getenv("PORT"))
	log.Fatal(srv.ListenAndServe())
}
