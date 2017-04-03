package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
    hello := r.URL.Path[len("/"):]
    content, err := ioutil.ReadFile("web/" + hello)
    if err != nil {
        fmt.Fprintf(w, "Error : %s", err);
    } else {
        fmt.Fprintf(w, "%s", content);
    }
}

func main() {
    http.HandleFunc("/", viewHandler)
    http.ListenAndServe(":8080", nil)
}

