package datatypes

import (
	"encoding/json"
	"time"
)

// Message is
type Message struct {
	Text   string
	Sender string
	Time   time.Time
	Detail json.RawMessage
}
