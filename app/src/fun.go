package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
  http.HandleFunc("/", FunServer)
  http.HandleFunc("/healthz", HealthZ)
  http.ListenAndServe(":80", nil)
}

func FunServer(w http.ResponseWriter, r *http.Request) {
  var hostname, _ = os.Hostname()
  fmt.Fprintf(w, "Host: %s Secret: %s", hostname, os.Getenv("SECRET"))
}

func HealthZ(w http.ResponseWriter, r *http.Request) {
  // just for demo
  fmt.Fprintf(w, "OK")
}
