package generator


type Type struct {
	Name         string              `json:"name"`
	Base         string              `json:"base"`
	TypeDoc      string              `json:"typeDoc"`
	Properties   map[string]Property `json:"properties"`
	ServiceDoc   string              `json:"serviceDoc"`
	Methods      map[string]Method   `json:"methods"`
	NoService    bool                `json:"noservice"`
	Overview     string              `json:"docOverview"`
	Deprecated   bool                `json:"deprecated"`
	ServiceGroup string
}

type Property struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	TypeArray  bool   `json:"typeArray"`
	Form       string `json:"form"`
	Doc        string `json:"doc"`
	Overview   string `json:"docOverview"`
	Deprecated bool   `json:"deprecated"`
}

type Method struct {
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	TypeArray  bool        `json:"typeArray"`
	Doc        string      `json:"doc"`
	Static     bool        `json:"static"`
	NoAuth     bool        `json:"noauth"`
	Limitable  bool        `json:"limitable"`
	Filterable bool        `json:"filterable"`
	Maskable   bool        `json:"maskable"`
	Parameters []Parameter `json:"parameters"`
	Overview   string      `json:"docOverview"`
	Deprecated bool        `json:"deprecated"`
}

type Parameter struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	TypeArray    bool        `json:"typeArray"`
	Doc          string      `json:"doc"`
	DefaultValue interface{} `json:"defaultValue"`
	ParentMethod string
}