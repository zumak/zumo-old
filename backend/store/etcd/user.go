package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/coreos/etcd/clientv3"
	"github.com/zumak/zumo/datatypes"
)

func (s *Store) FindUser() ([]datatypes.User, error) {
	resp, err := s.KV.Get(
		context.Background(),
		"/users/",
		clientv3.WithPrefix(),
	)
	if err != nil {
		return nil, err
	}
	result := []datatypes.User{}

	for _, pair := range resp.Kvs {
		user := &datatypes.User{}
		err := json.Unmarshal(pair.Value, user)
		if err != nil {
			return nil, err
		}
		result = append(result, *user)
	}

	return result, nil
}
func (s *Store) GetUser(username string) (*datatypes.User, error) {
	resp, err := s.KV.Get(
		context.Background(),
		fmt.Sprintf("/users/%s", username),
	)
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, nil
	}
	if len(resp.Kvs) > 1 {
		return nil, errors.New("not unique user")
	}

	user := &datatypes.User{}
	err = json.Unmarshal(resp.Kvs[0].Value, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *Store) PutUser(user *datatypes.User) (*datatypes.User, error) {
	str, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	_, err = s.KV.Put(
		context.Background(),
		fmt.Sprintf("/users/%s", user.Name),
		string(str),
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
