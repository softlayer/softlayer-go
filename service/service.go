package service

type Service_Provider struct {
	Session *Session
	Options
}

func (r *Session) GetServiceProviderService() Service_Provider {
	return Service_Provider{Session: r}
}
