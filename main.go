package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./phaser-tutorial")))
	http.ListenAndServe(":3000", nil)
}
