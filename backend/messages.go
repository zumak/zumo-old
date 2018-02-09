package backend

import (
	"encoding/json"
	"time"

	"github.com/zumak/zumo/datatypes"
)

func (b *backend) GetMessages(channelID string, limit int) ([]datatypes.Message, error) {
	messages, err := b.Store.FindMessages(channelID, limit)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
func (b *backend) AppendMessage(username, channelID, text string, detail json.RawMessage) (*datatypes.Message, error) {

	msg := &datatypes.Message{
		Sender: username,
		Text:   text,
		Detail: detail,
		Time:   time.Now(),
	}

	if _, err := b.Store.PutMessage(channelID, msg); err != nil {
		return nil, err
	}
	return msg, nil
}
