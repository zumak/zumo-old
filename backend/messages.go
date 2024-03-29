package backend

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"github.com/zumak/zumo/datatypes"
	"github.com/zumak/zumo/utils/log"
)

func (b *backend) GetMessages(channelID string, limit int) ([]datatypes.Message, error) {
	messages, err := b.Store.FindMessages(channelID, limit)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
func (b *backend) AppendMessage(username, channelID, text string, detail json.RawMessage) (*datatypes.Message, error) {
	log.Debug("%s:%s %s", username, channelID, text)
	channel, err := b.Store.GetChannel(channelID)
	if err != nil || channel == nil {
		return nil, errors.New("channel not found")
	}
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
