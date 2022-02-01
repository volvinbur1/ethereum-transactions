package cmn

import "time"

// Transaction describes parameters of each transaction in ethereum blockchain
type Transaction struct {
	Id          string    `json:"id" bson:"id"`
	Sender      string    `json:"sender" bson:"sender"`
	Recipient   string    `json:"recipient" bson:"recipient"`
	BlockNumber int       `json:"blockNumber" bson:"blockNumber"`
	Timestamp   time.Time `json:"timestamp" bson:"timestamp"`
	Value       float64   `json:"value" bson:"value"`
	Gas         float64   `json:"gas" bson:"gas"`
}

// getTransaction endpoints url parameters
const (
	FilterParameters = "filterBy"
	ValueParameters  = "value"
)

// filter possible values enum
const (
	IdFilter          = "id"
	SenderFilter      = "sender"
	RecipientFilter   = "recipient"
	BlockNumberFilter = "blockNumber"
	TimeFilter        = "timestamp"
)
