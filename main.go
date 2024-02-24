package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const port = ":8795"

func main() {
	handler := http.HandlerFunc(getTime)
	http.Handle("/time", handler)
	http.ListenAndServe(port, nil)
}

func getTime(w http.ResponseWriter, r *http.Request) {

}
