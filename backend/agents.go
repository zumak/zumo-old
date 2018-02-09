package backend

import (
	"sync"

	"github.com/google/uuid"
	"github.com/zumak/zumo/datatypes"
)

type Agent interface {
	OnMessage(channelID string, msg datatypes.Message)
	OnJoinChannel(channelID string)
	OnLeaveChannel(channelID string)
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
	return nil
}
func (al AgentList) Register(agent Agent) (string, error) {
	id := uuid.New().String()
	al[id] = agent

	return id, nil
}
