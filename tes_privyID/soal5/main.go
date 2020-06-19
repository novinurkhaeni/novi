package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	var env string
    env = os.Getenv("DEVELOPMENT")
    if os.Getenv("DEVELOPMENT") == env {
        fmt.Println(os.Getenv("DEVELOPMENT"))
        fmt.Fprintf(w, "Exciting New Feature")
    } else {
        fmt.Println(os.Getenv("DEVELOPMENT"))
        fmt.Fprintf(w, "existing boring feature")
    }
}

func handleRequests() {
	os.Setenv("DEVELOPMENT_PORT", ":8081")
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(os.Getenv("DEVELOPMENT_PORT"), nil))
}

func main() {
    handleRequests()
}