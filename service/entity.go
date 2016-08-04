package service

type Entity struct {
	Session *Session
	Options
}

func (r *Session) GetEntityService() Entity {
	return Entity{Session: r}
}
