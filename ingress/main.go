package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("API Gateway received request")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello from API Gateway!"))
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}