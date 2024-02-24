package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	handler := http.HandlerFunc(getTime)
	http.Handle("/time", handler)
	http.ListenAndServe(":8795", nil)
}

func getTime(w http.ResponseWriter, r *http.Request) {

}
