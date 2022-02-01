package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/volvinbur1/ethereum-transactions/internal/cmn"
	"github.com/volvinbur1/ethereum-transactions/internal/worker"
	"log"
	"net/http"
	"regexp"
)

var apiManager *worker.ApiManager

func SetManagerPointer(apiMgr *worker.ApiManager) {
	apiManager = apiMgr
}

// GetTransactionsHandler process call to api.
func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	jsonAnswer, errResp := analyzeRequest(r)
	w.Header().Set("Content-Type", "application/json")

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
	_, err := w.Write(jsonAnswer)
	if err != nil {
		log.Print(err)
	}
}

func analyzeRequest(r *http.Request) ([]byte, cmn.ErrorResponse) {
	actualFilter := r.URL.Query().Get(cmn.FilterParameters)
	switch actualFilter {
	case cmn.IdFilter:
		return getDesirableTransactions(actualFilter, r.URL.Query().Get(cmn.ValueParameters), "[0-9a-z]+")
	case cmn.SenderFilter:
	case cmn.RecipientFilter:
		return getDesirableTransactions(actualFilter, r.URL.Query().Get(cmn.ValueParameters), "[0-9a-zA-Z]+")
	case cmn.BlockNumberFilter:
	case cmn.TimeFilter:
		return getDesirableTransactions(actualFilter, r.URL.Query().Get(cmn.ValueParameters), "[0-9]+")
	default:
		return nil, cmn.ErrorResponse{
			Code:    cmn.UnsupportedFilterByParameter,
			Message: fmt.Sprintf("The specified filtering (`%s`) parameter is not supported", cmn.FilterParameters),
		}
	}

	return nil, cmn.ErrorResponse{}
}

func getDesirableTransactions(filter, value, validator string) ([]byte, cmn.ErrorResponse) {
	re := regexp.MustCompile(validator)
	if re.Match([]byte(value)) {
		return nil, cmn.ErrorResponse{
			Code:    cmn.IncorrectValueForSelectedFilter,
			Message: "Incorrect parameter value for filtering by " + filter,
		}
	}

	transactions, err := apiManager.DatabaseManager.FindByField(filter, value)
	if err != nil {
		return nil, cmn.ErrorResponse{}
	}

	jsonAnswer, err := json.Marshal(transactions)
	if err != nil {
		return nil, cmn.ErrorResponse{}
	}

	return jsonAnswer, cmn.ErrorResponse{}
}
