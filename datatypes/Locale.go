package datatypes

type Locale struct {
	Entity

	FriendlyName *string `json:"friendlyName,omitempty"`
	Id           *int    `json:"id,omitempty"`
	LanguageTag  *string `json:"languageTag,omitempty"`
	Name         *string `json:"name,omitempty"`
}

type Locale_Country struct {
	Entity

	IsEuropeanUnionFlag *int                   `json:"isEuropeanUnionFlag,omitempty"`
	LongName            *string                `json:"longName,omitempty"`
	ShortName           *string                `json:"shortName,omitempty"`
	StateCount          *uint                  `json:"stateCount,omitempty"`
	States              []Locale_StateProvince `json:"states,omitempty"`
}

type Locale_StateProvince struct {
	Entity

	LongName  *string `json:"longName,omitempty"`
	ShortName *string `json:"shortName,omitempty"`
}

type Locale_Timezone struct {
	Entity

	Id        *int    `json:"id,omitempty"`
	LongName  *string `json:"longName,omitempty"`
	Name      *string `json:"name,omitempty"`
	Offset    *string `json:"offset,omitempty"`
	ShortName *string `json:"shortName,omitempty"`
}
