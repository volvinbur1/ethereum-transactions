package main

import (
	"github.com/gorilla/mux"
	"github.com/volvinbur1/ethereum-transactions/internal/endpoints"
	"github.com/volvinbur1/ethereum-transactions/internal/worker"
	"log"
	"net/http"
	"os"
)

func createRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/transactions/byId/{id:[0-9a-z]+}", endpoints.GetTransactionById).Methods(http.MethodGet)
	router.HandleFunc("/api/transactions/bySender/{sender:[0-9a-zA-Z]+}", endpoints.GetTransactionBySender).Methods(http.MethodGet)
	router.HandleFunc("/api/transactions/byRecipient/{recipient:[0-9a-zA-Z]+}", endpoints.GetTransactionByRecipient).Methods(http.MethodGet)
	router.HandleFunc("/api/transactions/byTime/{time:[0-9]+}", endpoints.GetTransactionByTime).Methods(http.MethodGet)

	return router
}

func main() {
	mgr := worker.ApiManager{}
	endpoints.SetManagerPointer(&mgr)

	router := createRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}
