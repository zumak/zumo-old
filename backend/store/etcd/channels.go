package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/coreos/etcd/clientv3"
	"github.com/zumak/zumo/datatypes"
)

func (s *Store) FindChannels() ([]datatypes.Channel, error) {
	resp, err := s.Get(
		context.Background(),
		"/channels/",
		clientv3.WithPrefix(),
		//clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend),
		//clientv3.WithLimit(limit),
	)
	if err != nil {
		return nil, err
	}

	result := []datatypes.Channel{}
	for _, pair := range resp.Kvs {
		channel := &datatypes.Channel{}
		err := json.Unmarshal(pair.Value, channel)
		if err != nil {
			return nil, err
		}
		result = append(result, *channel)
	}
	return result, nil
}
func (s *Store) GetChannel(channelID string) (*datatypes.Channel, error) {
	resp, err := s.Get(
		context.Background(),
		fmt.Sprintf("/channels/%s", channelID),
	)
	if err != nil {
		return nil, err
	}
	if len(resp.Kvs) == 0 {
		return nil, nil
	}

	if len(resp.Kvs) > 1 {
		return nil, errors.New("is not a unique")
	}

	channel := &datatypes.Channel{}
	err = json.Unmarshal(resp.Kvs[0].Value, channel)
	if err != nil {
		return nil, err
	}

	return channel, nil
}
func (s *Store) PutChannel(channel *datatypes.Channel) (*datatypes.Channel, error) {
	str, err := json.Marshal(channel)
	if err != nil {
		return nil, err
	}

	_, err = s.Put(context.Background(), "/channels/"+channel.ID, string(str))
	if err != nil {
		return nil, err
	}
	return channel, nil
}
func (s *Store) DeleteChannel(ID string) error {
	_, err := s.KV.Delete(
		context.Background(),
		fmt.Sprintf("/channels/%s/", ID),
	)
	if err != nil {
		return err
	}

	return nil
}
