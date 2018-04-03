package backend

import (
	"sync"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/zumak/zumo/datatypes"
)

type Agent interface {
	OnMessage(channelID string, msg datatypes.Message)
	OnJoinChannel(channelID string)
	OnLeaveChannel(channelID string)
	Read() (channelID string, msg datatypes.Message, err error)
}

type AgentManager struct {
	sync.RWMutex
	agents map[string]AgentList
}

type AgentList map[string]Agent

func (am *AgentManager) Register(username string, agent Agent) (string, error) {
	am.Lock()
	defer am.Unlock()
	if am.agents[username] == nil {
		am.agents[username] = AgentList{}
	}
	return am.agents[username].Register(agent)
}
func (am *AgentManager) Unregister(username, id string) error {
	am.Lock()
	defer am.Lock()
	if am.agents[username] == nil {
		return errors.New("not found agent")
	}
	return am.agents[username].Unregister(id)
}
func (am *AgentManager) Get(username string) AgentList {
	am.RLock()
	defer am.RUnlock()
	return am.agents[username]
}
func (al AgentList) Register(agent Agent) (string, error) {
	id := uuid.New().String()
	al[id] = agent

	return id, nil
}
func (al AgentList) Unregister(id string) error {
	delete(al, id)
	return nil
}

// TODO lock
func (al AgentList) ForEach(cb func(agent Agent) error) error {
	if al == nil {
		return nil
	}
	for _, a := range al {
		if err := cb(a); err != nil {
			return err
		}
	}
	return nil
}
