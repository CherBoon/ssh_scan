package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":13337", ERouter())
}
