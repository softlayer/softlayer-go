package service

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Locale struct {
	Session *Session
	Options
}

func (r *Session) GetLocaleService() Locale {
	return Locale{Session: r}
}

func (r *Locale) GetClosestToLanguageTag(languageTag *string) (resp datatypes.Locale, err error) {
	params := []interface{}{
		languageTag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Locale) GetObject() (resp datatypes.Locale, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Locale_Country struct {
	Session *Session
	Options
}

func (r *Session) GetLocaleCountryService() Locale_Country {
	return Locale_Country{Session: r}
}

func (r *Locale_Country) GetAvailableCountries() (resp []datatypes.Locale_Country, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Locale_Country) GetCountries() (resp []datatypes.Locale_Country, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Locale_Country) GetObject() (resp datatypes.Locale_Country, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Locale_Country) GetStates() (resp []datatypes.Locale_StateProvince, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Locale_StateProvince struct {
	Session *Session
	Options
}

func (r *Session) GetLocaleStateProvinceService() Locale_StateProvince {
	return Locale_StateProvince{Session: r}
}

type Locale_Timezone struct {
	Session *Session
	Options
}

func (r *Session) GetLocaleTimezoneService() Locale_Timezone {
	return Locale_Timezone{Session: r}
}

func (r *Locale_Timezone) GetAllObjects() (resp []datatypes.Locale_Timezone, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Locale_Timezone) GetObject() (resp datatypes.Locale_Timezone, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
