package backend

import (
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/zumak/zumo/datatypes"
	"github.com/zumak/zumo/utils/log"
)

func (b *backend) GetChannels() ([]datatypes.Channel, error) {
	return b.Store.FindChannels()
}
func (b *backend) CreateChannel(name string, labels map[string]string) (*datatypes.Channel, error) {
	log.Debug("name: %s, labels: %v", name, labels)
	name = strings.Trim(name, " \t\r\n")
	//validate
	if len(name) < 4 {
		return nil, errors.Errorf("Too short channel name")
	}

	if labels == nil {
		labels = map[string]string{}
	}

	id := uuid.New()
	channels, err := b.Store.FindChannels()
	if err != nil {
		return nil, err
	}
	for _, channel := range channels {
		if channel.Name == name {
			return nil, errors.Errorf("channel '%s' already exist", name)
		}
	}
	channel, err := b.Store.PutChannel(&datatypes.Channel{
		ID:     id.String(),
		Name:   name,
		Labels: labels,
	})
	if err != nil {
		return nil, errors.Wrap(err, "channel create failed:")
	}

	return channel, nil
}
func (b *backend) DeleteChannel(channelID string) error {
	return b.Store.DeleteChannel(channelID)
}
