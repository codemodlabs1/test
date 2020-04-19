package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	read, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	fmt.Println("------------")
	fmt.Println(string(read))

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/webhook", handler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
