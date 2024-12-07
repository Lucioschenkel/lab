package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", Run)
	http.ListenAndServe(":8080", nil)
}

func Run(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("error") == "yese" {
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
