package backend

import (
	"github.com/pkg/errors"

	"github.com/zumak/zumo/datatypes"
	"github.com/zumak/zumo/utils/log"
)

func (b *backend) GetUser(username string) (*datatypes.User, error) {
	return b.Store.GetUser(username)
}

func (b *backend) ListUsers() ([]datatypes.User, error) {
	return b.Store.FindUser()
}
func (b *backend) CreateUser(username string, labels map[string]string) (*datatypes.User, error) {
	log.Debug("username: '%s', labels: %+v", username, labels)
	user := &datatypes.User{
		Name:   username,
		Labels: labels,
	}
	if len(user.Name) < 4 {
		return nil, errors.Errorf("username is too short: %s", username)
	}

	if u, err := b.Store.GetUser(user.Name); err != nil {
		return nil, err
	} else if u != nil {
		return nil, errors.Errorf("user already exist")
	}

	user, err := b.Store.PutUser(user)
	if err != nil {
		return nil, errors.Wrap(err, "create user failed:")
	}

	log.Debug("user '%s' created! (with %+v)", username, labels)
	return user, nil
}
