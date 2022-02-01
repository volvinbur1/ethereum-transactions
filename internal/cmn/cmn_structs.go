package cmn

import "time"

type Transaction struct {
	Id          string    `json:"id" bson:"id"`
	Sender      string    `json:"sender" bson:"sender"`
	Recipient   string    `json:"recipient" bson:"recipient"`
	BlockNumber int       `json:"blockNumber" bson:"blockNumber"`
	Timestamp   time.Time `json:"timestamp" bson:"timestamp"`
	Value       float64   `json:"value" bson:"value"`
	Gas         float64   `json:"gas" bson:"gas"`
}
