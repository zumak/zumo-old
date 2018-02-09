package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"path"
)

// genaral data, for pod or bots.
func (s *Store) RestoreData(namespace, key string, data interface{}) error {
	//s.KV.
	resp, err := s.KV.Get(
		context.Background(),
		path.Join("data", namespace, key),
	)
	if err != nil {
		return err
	}
	if len(resp.Kvs) != 1 {
		return errors.New("not unique")
	}
	return json.Unmarshal(resp.Kvs[0].Value, data)

}
func (s *Store) SaveData(namespace, key string, data interface{}) error {
	str, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = s.KV.Put(
		context.Background(),
		path.Join("data", namespace, key),
		string(str),
	)
	if err != nil {
		return err
	}
	return nil
}
