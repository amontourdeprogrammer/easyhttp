package main

import(
    "fmt"
    "io/ioutil"
    "net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
    fernando := r.URL.Path[len("/"):]
    content, err := ioutil.ReadFile("web/" + fernando)
    if err != nil {
        fmt.Fprintf(w, "Error : %s", err);
    } else {
        fmt.Fprintf(w, "%s", content);
    }
}



func main() {    
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
    
}
