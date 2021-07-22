package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := "3000"
	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()
	log.Println("v1 API")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	})
	apiRouter.HandleFunc("/", get).Methods(http.MethodGet)
	apiRouter.HandleFunc("/", post).Methods(http.MethodPost)
	apiRouter.HandleFunc("/", delete).Methods(http.MethodDelete)
	apiRouter.HandleFunc("/", patch).Methods(http.MethodPatch)
	apiRouter.HandleFunc("/health", patch).Methods(http.MethodPatch)

	s := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Fatalln(s.ListenAndServe())

}
