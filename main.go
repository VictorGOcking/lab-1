package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"fmt"
)

const PORT = ":8795"

func getTime(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusCreated)
  w.Header().Set("Content-Type", "application/json")

  resp := make(map[string]string)
  resp["time"] = time.Now().Format(time.RFC3339)
  jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

  w.Write(jsonResp)
}

func main() {
	handler := http.HandlerFunc(getTime)
	http.Handle("/time", handler)

	fmt.Println("Server is listening on the port", PORT)
	http.ListenAndServe(PORT, nil)
}