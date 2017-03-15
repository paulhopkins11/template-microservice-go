package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Success! The Framework Training GO template microservices is up and running!")
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}