package backend

import (
	"github.com/pkg/errors"
	"github.com/zumak/zumo/datatypes"
)

func (b *backend) Join(channelID, username string) error {
	// b.channels[channelID].Member.append(username)
	channel, err := b.Store.GetChannel(channelID)
	if err != nil {
		return err
	}
	// TODO  duplicate check
	channel.Member[username] = struct{}{}

	_, err = b.Store.PutChannel(channel) // maybe need hint?
	if err != nil {
		return err
	}

	return nil
}
func (b *backend) Leave(channelID, username string) error {
	channel, err := b.Store.GetChannel(channelID)
	if err != nil {
		return err
	}
	delete(channel.Member, username)

	// check member if not exist return error
	_, err = b.Store.PutChannel(channel)
	if err != nil {
		return err
	}

	return nil
}
func (b *backend) JoinnedChannel(username string) ([]datatypes.Channel, error) {
	channels, err := b.Store.FindChannels()
	if err != nil {
		return nil, errors.Wrap(err, "cannot get channel list:")
	}

	result := []datatypes.Channel{}
	for _, channel := range channels {
		if _, ok := channel.Member[username]; ok {
			result = append(result, channel)
		}
	}

	return result, nil
}
