package backend

import (
	"encoding/json"

	"github.com/zumak/zumo/backend/store"
	"github.com/zumak/zumo/datatypes"
	"github.com/zumak/zumo/utils/log"
)

// Config is
type Config struct {
	Store store.Config
}

// Backend is
type Backend interface {
	GetChannels() ([]datatypes.Channel, error)
	CreateChannel(name string, labels map[string]string) (*datatypes.Channel, error)
	DeleteChannel(channelID string) error

	GetMessages(channelID string, limit int) ([]datatypes.Message, error)
	AppendMessage(username, channelID, text string, detail json.RawMessage) (*datatypes.Message, error)

	GetUser(username string) (*datatypes.User, error)
	ListUsers() ([]datatypes.User, error)
	CreateUser(username string, labels map[string]string) (*datatypes.User, error)
	//DeleteUser(username string) error

	CreateToken(username, unhashedKey string) (*datatypes.Token, error)
	Token(tokenString string) (*datatypes.Token, error)

	CreateHook(channelID, username string) (*datatypes.Hook, error)
	DoHook(hookID, text string, detail json.RawMessage) (*datatypes.Message, error)

	JoinnedChannel(username string) ([]datatypes.Channel, error)

	Join(channeID, username string) error
	Leave(channelID, username string) error

	RestoreData(namespace, key string, v interface{}) error
	SaveData(namespace, key string, v interface{}) error

	// block until end, and automatically remove it when return
	OpenSession(username string, agent Agent) error
}

// Agents manage UserAgent for session like websocket, http2, grpc
type Agents interface {
	// Register(username string, agent Agent) (string, error)
	// Unregister(id) error
}

// New is
func New(conf *Config) (Backend, error) {
	//b := &backend{}
	s, reciver, err := store.New(conf.Store.Driver, conf.Store.Endpoint, conf.Store.Options)
	if err != nil {
		return nil, err
	}

	go runDispatcher(reciver)

	am := &AgentManager{
		agents: map[string]AgentList{},
	}

	return &backend{s, am}, nil
}

type backend struct {
	store.Store
	agents *AgentManager
}

func runDispatcher(events *store.EventReciever) {
	for {
		select {
		case msg := <-events.PutMessage:
			log.Info("message recieved: %+v", msg)
			// TODO
			// 1. find sessions(UserAgent) that member of channel
			// 2. send message to them
		}
	}
}
