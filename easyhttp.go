package main

import (

	"io"
	"net/http"
)


func viewHandler(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, `{"hello":"world"})
}

func main() {
 	 http.HandleFunc("/" , viewHandler)
 	 http.ListenAndServe(":8080", nil)
 }