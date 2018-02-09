package backend

func (b *backend) OpenSession(username string, agent Agent) error {
	id, err := b.agents.Register(username, agent)
	if err != nil {
		return err
	}
	defer b.agents.Unregister(username, id)

	return nil
}
