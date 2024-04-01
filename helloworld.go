package main

import (
	"fmt"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
