package main

import (
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

	log.Printf("Starting server at port %s \n", os.Getenv("PORT"))
	log.Fatal(srv.ListenAndServe())
}
