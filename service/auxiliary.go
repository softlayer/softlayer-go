package service

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Auxiliary_Marketing_Event struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryMarketingEventService() Auxiliary_Marketing_Event {
	return Auxiliary_Marketing_Event{Session: r}
}

func (r *Auxiliary_Marketing_Event) GetMarketingEvents() (resp []datatypes.Auxiliary_Marketing_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Auxiliary_Marketing_Event) GetObject() (resp datatypes.Auxiliary_Marketing_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Auxiliary_Network_Status struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryNetworkStatusService() Auxiliary_Network_Status {
	return Auxiliary_Network_Status{Session: r}
}

func (r *Auxiliary_Network_Status) GetNetworkStatus(target *string) (resp []datatypes.Container_Auxiliary_Network_Status_Reading, err error) {
	params := []interface{}{
		target,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Auxiliary_Notification_Emergency struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryNotificationEmergencyService() Auxiliary_Notification_Emergency {
	return Auxiliary_Notification_Emergency{Session: r}
}

func (r *Auxiliary_Notification_Emergency) GetAllObjects() (resp []datatypes.Auxiliary_Notification_Emergency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Auxiliary_Notification_Emergency) GetCurrentNotifications() (resp []datatypes.Auxiliary_Notification_Emergency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Auxiliary_Notification_Emergency) GetObject() (resp datatypes.Auxiliary_Notification_Emergency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Auxiliary_Notification_Emergency) GetSignature() (resp datatypes.Auxiliary_Notification_Emergency_Signature, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Auxiliary_Notification_Emergency) GetStatus() (resp datatypes.Auxiliary_Notification_Emergency_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Auxiliary_Notification_Emergency_Signature struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryNotificationEmergencySignatureService() Auxiliary_Notification_Emergency_Signature {
	return Auxiliary_Notification_Emergency_Signature{Session: r}
}

type Auxiliary_Notification_Emergency_Status struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryNotificationEmergencyStatusService() Auxiliary_Notification_Emergency_Status {
	return Auxiliary_Notification_Emergency_Status{Session: r}
}

type Auxiliary_Press_Release struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryPressReleaseService() Auxiliary_Press_Release {
	return Auxiliary_Press_Release{Session: r}
}

func (r *Auxiliary_Press_Release) GetAllObjects() (resp []datatypes.Auxiliary_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Auxiliary_Press_Release) GetObject() (resp datatypes.Auxiliary_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Auxiliary_Press_Release) GetRenderedPressRelease() (resp []datatypes.Auxiliary_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Auxiliary_Press_Release) GetRenderedPressReleases(resultLimit *string, year *string) (resp []datatypes.Auxiliary_Press_Release, err error) {
	params := []interface{}{
		resultLimit,
		year,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Auxiliary_Press_Release) GetWebsiteHighlightPressReleases() (resp []datatypes.Auxiliary_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Auxiliary_Press_Release) GetAbout() (resp []datatypes.Auxiliary_Press_Release_About_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Auxiliary_Press_Release) GetContacts() (resp []datatypes.Auxiliary_Press_Release_Contact_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Auxiliary_Press_Release) GetMediaPartners() (resp []datatypes.Auxiliary_Press_Release_Media_Partner_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Auxiliary_Press_Release) GetPressReleaseContent() (resp datatypes.Auxiliary_Press_Release_Content, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Auxiliary_Press_Release_About struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryPressReleaseAboutService() Auxiliary_Press_Release_About {
	return Auxiliary_Press_Release_About{Session: r}
}

func (r *Auxiliary_Press_Release_About) GetObject() (resp datatypes.Auxiliary_Press_Release_About, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Auxiliary_Press_Release_About_Press_Release struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryPressReleaseAboutPressReleaseService() Auxiliary_Press_Release_About_Press_Release {
	return Auxiliary_Press_Release_About_Press_Release{Session: r}
}

func (r *Auxiliary_Press_Release_About_Press_Release) GetObject() (resp datatypes.Auxiliary_Press_Release_About_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Auxiliary_Press_Release_About_Press_Release) GetAboutParagraphs() (resp []datatypes.Auxiliary_Press_Release_About, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Auxiliary_Press_Release_About_Press_Release) GetPressReleases() (resp []datatypes.Auxiliary_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Auxiliary_Press_Release_Contact struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryPressReleaseContactService() Auxiliary_Press_Release_Contact {
	return Auxiliary_Press_Release_Contact{Session: r}
}

func (r *Auxiliary_Press_Release_Contact) GetObject() (resp datatypes.Auxiliary_Press_Release_Contact, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Auxiliary_Press_Release_Contact_Press_Release struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryPressReleaseContactPressReleaseService() Auxiliary_Press_Release_Contact_Press_Release {
	return Auxiliary_Press_Release_Contact_Press_Release{Session: r}
}

func (r *Auxiliary_Press_Release_Contact_Press_Release) GetObject() (resp datatypes.Auxiliary_Press_Release_Contact_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Auxiliary_Press_Release_Contact_Press_Release) GetContacts() (resp []datatypes.Auxiliary_Press_Release_Contact, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Auxiliary_Press_Release_Contact_Press_Release) GetPressReleases() (resp []datatypes.Auxiliary_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Auxiliary_Press_Release_Content struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryPressReleaseContentService() Auxiliary_Press_Release_Content {
	return Auxiliary_Press_Release_Content{Session: r}
}

func (r *Auxiliary_Press_Release_Content) GetObject() (resp datatypes.Auxiliary_Press_Release_Content, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Auxiliary_Press_Release_Media_Partner struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryPressReleaseMediaPartnerService() Auxiliary_Press_Release_Media_Partner {
	return Auxiliary_Press_Release_Media_Partner{Session: r}
}

func (r *Auxiliary_Press_Release_Media_Partner) GetObject() (resp datatypes.Auxiliary_Press_Release_Media_Partner, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Auxiliary_Press_Release_Media_Partner_Press_Release struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryPressReleaseMediaPartnerPressReleaseService() Auxiliary_Press_Release_Media_Partner_Press_Release {
	return Auxiliary_Press_Release_Media_Partner_Press_Release{Session: r}
}

func (r *Auxiliary_Press_Release_Media_Partner_Press_Release) GetObject() (resp datatypes.Auxiliary_Press_Release_Media_Partner_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Auxiliary_Press_Release_Media_Partner_Press_Release) GetMediaPartners() (resp []datatypes.Auxiliary_Press_Release_Media_Partner, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Auxiliary_Press_Release_Media_Partner_Press_Release) GetPressReleases() (resp []datatypes.Auxiliary_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Auxiliary_Shipping_Courier struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryShippingCourierService() Auxiliary_Shipping_Courier {
	return Auxiliary_Shipping_Courier{Session: r}
}

type Auxiliary_Shipping_Courier_Type struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryShippingCourierTypeService() Auxiliary_Shipping_Courier_Type {
	return Auxiliary_Shipping_Courier_Type{Session: r}
}

func (r *Auxiliary_Shipping_Courier_Type) GetObject() (resp datatypes.Auxiliary_Shipping_Courier_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Auxiliary_Shipping_Courier_Type) GetTypeByKeyName(keyName *string) (resp datatypes.Auxiliary_Shipping_Courier_Type, err error) {
	params := []interface{}{
		keyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Auxiliary_Shipping_Courier_Type) GetCourier() (resp []datatypes.Auxiliary_Shipping_Courier, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
