package backend

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/zumak/zumo/datatypes"
	"github.com/zumak/zumo/utils"
	"github.com/zumak/zumo/utils/log"
)

func (b *backend) CreateHook(channelID, username string) (*datatypes.Hook, error) {
	log.Debug("create hook, channelID: %s, Username: %s", channelID, username)
	// TODO check vaild
	return b.Store.PutHook(&datatypes.Hook{
		ID:        utils.RandomString(16),
		ChannelID: channelID,
		Username:  username,
	})
}
func (b *backend) DoHook(hookID, text string, detail json.RawMessage) (*datatypes.Message, error) {
	log.Debug("[backend:DoHook] HookID: %s, %s", hookID, text)

	if strings.Trim(text, " \r\n\t") == "" {
		return nil, errors.New("hook must have text")
	}

	hook, err := b.Store.GetHook(hookID)
	if err != nil {
		return nil, errors.Wrap(err, "fail to find hook")
	}

	if detail == nil {
		detail = []byte("{}")
	}

	msg := &datatypes.Message{
		Sender: hook.Username,
		Text:   text,
		Detail: detail,
		Time:   time.Now(),
	}

	if _, err := b.Store.PutMessage(hook.ChannelID, msg); err != nil {
		return nil, errors.Wrap(err, "put message failed")
	}
	return msg, nil
}
