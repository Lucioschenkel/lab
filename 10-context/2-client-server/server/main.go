package main

import (
	"log"
	"net/http"
	"time"
)

func main()  {
  http.HandleFunc("/", handle)
  http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request)  {
  ctx := r.Context()
  log.Println("Request iniciada")
  defer log.Println("Request finalizada")

  select {
  case <-time.After(5 * time.Second):
    log.Println("Request processada com sucesso")
    w.Write([]byte("Request processada com sucesso"))
  case <-ctx.Done():
    log.Println("Request cancelada pelo cliente")
  }
}
