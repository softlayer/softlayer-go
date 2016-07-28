package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Locale interface {
	Entity

	GetClosestToLanguageTag(languageTag string) (datatypes.Locale, error)
	GetObject() (datatypes.Locale, error)
}

type Locale_Country interface {
	Entity

	GetAvailableCountries() (datatypes.Locale_Country, error)
	GetCountries() (datatypes.Locale_Country, error)
	GetObject() (datatypes.Locale_Country, error)

	GetStates() (datatypes.Locale_StateProvince, error)
}

type Locale_StateProvince interface {
	Entity
}

type Locale_Timezone interface {
	Entity

	GetAllObjects() (datatypes.Locale_Timezone, error)
	GetObject() (datatypes.Locale_Timezone, error)
}
