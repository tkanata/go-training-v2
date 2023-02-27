package main

import (
    "net/http"

    "github.com/tkanata/go-training-v2/controller"
    "github.com/tkanata/go-training-v2/model/repository"
)

var hr = repository.NewHandRepository()
var hc = controller.NewHandController(hr)
var ro = controller.NewRouter(hc)

func main() {
    server := http.Server{
        Addr: ":8080",
    }
    http.HandleFunc("/poker", ro.HandleCardsRequest)
    server.ListenAndServe()
}

