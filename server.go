package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/configMap", ConfigMap)
	http.ListenAndServe(":8087", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello, I`m %s. I`m %s.", name, age)
	
	// w.Write([]byte("<h1>Hello Everyone</h1>"))
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("/go/myfamily/myfamily.txt")

	if err != nil {
		log.Fatalf("Eroro readoing file: ", err)
	}

	fmt.Fprintf(w, "My Family:  %s.", string(data))
}