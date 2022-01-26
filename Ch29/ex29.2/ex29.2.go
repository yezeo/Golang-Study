package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "World"
	}
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)
}

func main() {
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":3000", nil)
}
