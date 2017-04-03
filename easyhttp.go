package main

import (
        "fmt"
        "net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    fmt.Fprintf(w, "<html><title>%s</title><body><h1>View : %s</h1></body></html>", title, title);
}

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.ListenAndServe(":8080", nil)
}