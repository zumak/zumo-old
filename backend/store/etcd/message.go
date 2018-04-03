package etcd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coreos/etcd/clientv3"
	"github.com/zumak/zumo/datatypes"
	"github.com/zumak/zumo/utils/log"
)

func (s *Store) FindMessages(channelID string, limit int) ([]datatypes.Message, error) {
	resp, err := s.Get(context.Background(),
		fmt.Sprintf("/messages/%s/", channelID),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend),
		clientv3.WithPrefix(),
		clientv3.WithLimit(int64(limit)),
	)
	if err != nil {
		return nil, err
	}

	result := []datatypes.Message{}
	for _, pair := range resp.Kvs {
		msg := &datatypes.Message{}
		err := json.Unmarshal(pair.Value, msg)
		if err != nil {
			return nil, err
		}
		result = append(result, *msg)
	}
	return result, nil
}
func (s *Store) PutMessage(channelID string, msg *datatypes.Message) (*datatypes.Message, error) {
	log.Debug("%s %+v", channelID, msg)
	str, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf("/messages/%s/%.16x", channelID, msg.Time.UnixNano())

	_, err = s.Put(context.Background(), key, string(str))
	if err != nil {
		return nil, err
	}
	return msg, nil
}
