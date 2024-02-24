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
  w.WriteHeader(http.StatusCreated)
  w.Header().Set("Content-Type", "application/json")

  resp := make(map[string]string)
  resp["time"] = time.Now().Format(time.RFC3339)
  jsonResp, err := json.Marshal(resp)

  w.Write(jsonResp)
}
