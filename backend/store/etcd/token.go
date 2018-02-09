package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/coreos/etcd/clientv3"
	"github.com/zumak/zumo/datatypes"
)

func (s *Store) GetToken(username, hashedKey string) (*datatypes.Token, error) {
	resp, err := s.KV.Get(
		context.Background(),
		fmt.Sprintf("/tokens/%s:%s", username, hashedKey),
	)
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) != 1 {
		return nil, errors.New("not unique")
	}

	token := &datatypes.Token{}
	err = json.Unmarshal(resp.Kvs[0].Value, token)
	if err != nil {
		return nil, err
	}

	return token, nil
}
func (s *Store) FindToken(username string) ([]datatypes.Token, error) {
	resp, err := s.Get(
		context.Background(),
		fmt.Sprintf("/tokens/%s:", username),
		clientv3.WithPrefix(),
	)
	if err != nil {
		return nil, err
	}

	result := []datatypes.Token{}
	for _, pair := range resp.Kvs {
		token := &datatypes.Token{}
		err := json.Unmarshal(pair.Value, token)
		if err != nil {
			return nil, err
		}
		result = append(result, *token)
	}
	return result, nil

}
func (s *Store) PutToken(token *datatypes.Token) (*datatypes.Token, error) {
	str, err := json.Marshal(token)
	if err != nil {
		return nil, err
	}

	_, err = s.KV.Put(
		context.Background(),
		fmt.Sprintf("/tokens/%s:%s", token.Username, token.HashedKey),
		string(str),
	)
	if err != nil {
		return nil, err
	}

	return token, nil
}
