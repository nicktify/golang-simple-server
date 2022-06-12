package main

import (
  "fmt"
  "net/http"
  "log"
  "encoding/json"
)

func hello(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "hello")
}

type ResponseBody struct {
    SomeNumber int64
    SomeString string
}

func jsonResponse(w http.ResponseWriter, req *http.Request) {
  object := ResponseBody{
    12,
    "Hello there",
  }
  w.Header().Set("Content-Type", "application/json")
  jsonResp, err := json.Marshal(object)
  if err != nil {
    log.Fatalf("Error happened in JSON marshal. Err: %s", err)
  }
  w.Write(jsonResp)
}

func main() {
  http.HandleFunc("/hello", hello)
  http.HandleFunc("/json", jsonResponse)
  http.ListenAndServe(":8080", nil)
}

