package backend

import "io"

func (b *backend) OpenSession(username string, agent Agent) error {
	id, err := b.agents.Register(username, agent)
	if err != nil {
		return err
	}
	defer b.agents.Unregister(username, id)

	for {
		// TODO if get other message?
		channelID, msg, err := agent.Read()
		if err != nil {
			if err == io.EOF {
				// just end
				return nil
			}
			return err
		}
		//
		b.AppendMessage(username, channelID, msg.Text, msg.Detail)
	}
	return nil
}
