package worker

import (
	"github.com/volvinbur1/ethereum-transactions/internal/db"
	"sync"
	"sync/atomic"
	"time"
)

type ApiManager struct {
	DatabaseManager *db.Manager

	isWorking atomic.Value
	wg        sync.WaitGroup
}

func (m *ApiManager) Start() {
	m.isWorking.Store(true)
	m.DatabaseManager = db.New()

	m.wg.Add(1)
	go m.etherscanMonitoring()
}

func (m *ApiManager) Stop() {
	m.isWorking.Store(false)
	defer m.DatabaseManager.Disconnect()
}

func (m *ApiManager) etherscanMonitoring() {
	defer m.wg.Done()

	for m.isWorking.Load().(bool) {
		// TODO add request to etherscan
		time.Sleep(time.Second)
	}
}
