package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/volvinbur1/ethereum-transactions/internal/cmn"
	"github.com/volvinbur1/ethereum-transactions/internal/worker"
	"log"
	"net/http"
)

var apiManager *worker.ApiManager

func SetManagerPointer(apiMgr *worker.ApiManager) {
	apiManager = apiMgr
}

func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	jsonAnswer, errResp := analyzeRequest(r)
	if errResp.Code != cmn.Ok {
		errorJson, err := json.Marshal(errResp)
		if err != nil {
			log.Print(err)
			http.Error(w, "unknown error", http.StatusInternalServerError)
		}
		http.Error(w, string(errorJson), http.StatusBadRequest)
		return
	}

	if jsonAnswer == nil || len(jsonAnswer) == 0 {
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(jsonAnswer)
	if err != nil {
		log.Print(err)
	}
}

func analyzeRequest(r *http.Request) ([]byte, cmn.ErrorResponse) {
	actualFilter := r.URL.Query().Get(cmn.FilterParameters)
	switch actualFilter {
	case cmn.IdFilter:
		return nil, cmn.ErrorResponse{}
	case cmn.SenderFilter:
		return nil, cmn.ErrorResponse{}
	case cmn.RecipientFilter:
		return nil, cmn.ErrorResponse{}
	case cmn.BlockNumberFilter:
		return nil, cmn.ErrorResponse{}
	case cmn.TimeFilter:
		return nil, cmn.ErrorResponse{}
	default:
		return nil, cmn.ErrorResponse{
			Code:    cmn.UnsupportedFilterByParameter,
			Message: fmt.Sprintf("The specified filtering (`%s`) parameter is not supported.", cmn.FilterParameters),
		}
	}
}

func transactionsById() {

}
