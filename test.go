package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		log.Println("hello Server")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s", d)
	})
	http.ListenAndServe(":8080", nil)
}