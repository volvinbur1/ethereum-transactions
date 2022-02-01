package main

import (
	"github.com/gorilla/mux"
	"github.com/volvinbur1/ethereum-transactions/internal/endpoints"
	"github.com/volvinbur1/ethereum-transactions/internal/worker"
	"log"
	"net/http"
	"os"
)

func main() {
	mgr := worker.ApiManager{}
	endpoints.SetManagerPointer(&mgr)

	router := mux.NewRouter()
	router.HandleFunc("/api/getTransactions", endpoints.GetTransactionsHandler).Methods(http.MethodGet)

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}
