package softlayer

type Utility_Bandwidth_Graph interface {
	Entity
}

type Utility_Network interface {
	Entity

	NsLookup(address string, typ string) (string, error)
	Whois(address string) (string, error)
}

type Utility_ObjectFilter interface {
	Entity
}

type Utility_ObjectFilter_Operation interface {
	Entity
}

type Utility_ObjectFilter_Operation_Option interface {
	Entity
}
