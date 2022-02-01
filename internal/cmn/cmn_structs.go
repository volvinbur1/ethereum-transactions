package cmn

import "time"

type Transaction struct {
	Id               string    `json:"id"`
	SenderAddress    string    `json:"senderAddress"`
	RecipientAddress string    `json:"recipientAddress"`
	BlockNumber      int       `json:"blockNumber"`
	Timestamp        time.Time `json:"timestamp"`
	Value            float64   `json:"value"`
	Gas              float64   `json:"gas"`
}
