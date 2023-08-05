package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func helloHandler(rw http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(rw, "%s", "hello")

}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()

	router.HandleFunc("/", helloHandler)
	fmt.Printf("server started at port %s", port)
	http.ListenAndServe(":"+port, router)

}
