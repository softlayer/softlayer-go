package service

type Utility_Bandwidth_Graph struct {
	Session *Session
	Options
}

func (r *Session) GetUtilityBandwidthGraphService() Utility_Bandwidth_Graph {
	return Utility_Bandwidth_Graph{Session: r}
}

type Utility_Network struct {
	Session *Session
	Options
}

func (r *Session) GetUtilityNetworkService() Utility_Network {
	return Utility_Network{Session: r}
}

func (r *Utility_Network) NsLookup(address *string, typ *string) (resp string, err error) {
	params := []interface{}{
		address,
		typ,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Utility_Network) Whois(address *string) (resp string, err error) {
	params := []interface{}{
		address,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Utility_ObjectFilter struct {
	Session *Session
	Options
}

func (r *Session) GetUtilityObjectFilterService() Utility_ObjectFilter {
	return Utility_ObjectFilter{Session: r}
}

type Utility_ObjectFilter_Operation struct {
	Session *Session
	Options
}

func (r *Session) GetUtilityObjectFilterOperationService() Utility_ObjectFilter_Operation {
	return Utility_ObjectFilter_Operation{Session: r}
}

type Utility_ObjectFilter_Operation_Option struct {
	Session *Session
	Options
}

func (r *Session) GetUtilityObjectFilterOperationOptionService() Utility_ObjectFilter_Operation_Option {
	return Utility_ObjectFilter_Operation_Option{Session: r}
}
