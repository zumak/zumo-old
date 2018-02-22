package store

import (
	"github.com/zumak/zumo/datatypes"
)

// Store is cache of kv store, and watch kv store for update
type Store interface {
	FindChannels() ([]datatypes.Channel, error)
	GetChannel(channelID string) (*datatypes.Channel, error) // Must return nil if not exist
	PutChannel(channel *datatypes.Channel) (*datatypes.Channel, error)
	DeleteChannel(ID string) error

	FindMessages(channelID string, limit int) ([]datatypes.Message, error)
	PutMessage(channelID string, msg *datatypes.Message) (*datatypes.Message, error)

	FindUser() ([]datatypes.User, error)
	GetUser(username string) (*datatypes.User, error) // Must return nil if not exist
	PutUser(user *datatypes.User) (*datatypes.User, error)

	//GetToken(username, hashedKey string) (*datatypes.Token, error)
	FindToken(username string) ([]datatypes.Token, error)
	PutToken(token *datatypes.Token) (*datatypes.Token, error)

	GetHook(hookID string) (*datatypes.Hook, error)
	PutHook(*datatypes.Hook) (*datatypes.Hook, error)

	// genaral data, for pod or bots.
	RestoreData(namespace, key string, v interface{}) error
	SaveData(namespace, key string, v interface{}) error
}

type InitFunc func(emitter *EventEmitter, endpoint string, option map[string]string) (Store, error)

var drivers map[string]InitFunc

func init() {
	drivers = map[string]InitFunc{}
}
func Register(name string, driver InitFunc) {
	drivers[name] = driver
}
func New(driver string, endpoint string, opt map[string]string) (Store, *EventReciever, error) {
	emmiter, receiver := newEventSystem()

	s, err := drivers[driver](emmiter, endpoint, opt)
	if err != nil {
		return nil, nil, err
	}
	return s, receiver, nil
}

type Config struct {
	Driver   string
	Endpoint string
	Options  map[string]string
}
