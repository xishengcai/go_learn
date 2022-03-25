package main

import (
	"fmt"
	"github.com/amahi/spdy"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	//io.WriteString(w,"")

	fmt.Fprintf(w, "Hello V1")
}

func main() {
	http.HandleFunc("/", handler)

	//use spdy's Listen and serve
	err := spdy.ListenAndServe("localhost:4040", nil)
	if err != nil {
		//error handling here
	}
}
