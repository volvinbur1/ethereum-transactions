package endpoints

import (
	"github.com/volvinbur1/ethereum-transactions/internal/worker"
	"net/http"
)

var apiManager *worker.ApiManager

func SetManagerPointer(apiMgr *worker.ApiManager) {
	apiManager = apiMgr
}

func GetTransactionById(w http.ResponseWriter, r *http.Request) {

}

func GetTransactionBySender(w http.ResponseWriter, r *http.Request) {

}

func GetTransactionByRecipient(w http.ResponseWriter, r *http.Request) {

}

func GetTransactionByTime(w http.ResponseWriter, r *http.Request) {

}
