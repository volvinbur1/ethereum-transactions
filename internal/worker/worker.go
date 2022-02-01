package worker

import (
	"github.com/volvinbur1/ethereum-transactions/internal/db"
	"sync"
	"time"
)

type ApiManager struct {
	DatabaseManager *db.Manager

	isWorking bool
	wg        sync.WaitGroup
}

func (m *ApiManager) Start() {
	go m.etherscanMonitoring()
}

func (m *ApiManager) etherscanMonitoring() {
	for {
		// TODO add request to etherscan
		time.Sleep(time.Second)
	}
}
