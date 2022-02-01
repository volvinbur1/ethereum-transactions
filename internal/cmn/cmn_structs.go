package cmn

import "time"

type Transaction struct {
	Id          string    `json:"id"`
	Sender      string    `json:"sender"`
	Recipient   string    `json:"recipient"`
	BlockNumber int       `json:"blockNumber"`
	Timestamp   time.Time `json:"timestamp"`
	Value       float64   `json:"value"`
	Gas         float64   `json:"gas"`
}
