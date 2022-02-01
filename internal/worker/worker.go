package worker

import (
	"sync"
	"time"
)

type ApiManager struct {
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
