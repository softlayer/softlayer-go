package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Dns_Domain interface {
	Entity

	CreateARecord(host string, data string, ttl int) (datatypes.Dns_Domain_ResourceRecord_AType, error)
	CreateAaaaRecord(host string, data string, ttl int) (datatypes.Dns_Domain_ResourceRecord_AaaaType, error)
	CreateCnameRecord(host string, data string, ttl int) (datatypes.Dns_Domain_ResourceRecord_CnameType, error)
	CreateMxRecord(host string, data string, ttl int, mxPriority int) (datatypes.Dns_Domain_ResourceRecord_MxType, error)
	CreateNsRecord(host string, data string, ttl int) (datatypes.Dns_Domain_ResourceRecord_NsType, error)
	CreateObject(templateObject datatypes.Dns_Domain) (datatypes.Dns_Domain, error)
	CreateObjects(templateObjects datatypes.Dns_Domain) (datatypes.Dns_Domain, error)
	CreatePtrRecord(ipAddress string, ptrRecord string, ttl int) (datatypes.Dns_Domain_ResourceRecord, error)
	CreateSpfRecord(host string, data string, ttl int) (datatypes.Dns_Domain_ResourceRecord_SpfType, error)
	CreateTxtRecord(host string, data string, ttl int) (datatypes.Dns_Domain_ResourceRecord_TxtType, error)
	DeleteObject() (bool, error)
	GetByDomainName(name string) (datatypes.Dns_Domain, error)
	GetObject() (datatypes.Dns_Domain, error)
	GetZoneFileContents() (string, error)

	GetAccount() (datatypes.Account, error)
	GetManagedResourceFlag() (bool, error)
	GetResourceRecords() (datatypes.Dns_Domain_ResourceRecord, error)
	GetSecondary() (datatypes.Dns_Secondary, error)
}

type Dns_Domain_Forward interface {
	Dns_Domain
}

type Dns_Domain_Registration interface {
	Entity

	AddNameserversToDomain(nameservers string) (bool, error)
	DeleteRegisteredNameserver(nameserver string) (bool, error)
	GetAuthenticationCode() (string, error)
	GetDomainInformation() (datatypes.Container_Dns_Domain_Registration_Information, error)
	GetDomainNameservers() (datatypes.Container_Dns_Domain_Registration_Nameserver, error)
	GetExtendedAttributes(domainName string) (datatypes.Container_Dns_Domain_Registration_ExtendedAttribute, error)
	GetObject() (datatypes.Dns_Domain_Registration, error)
	GetRegisteredNameserver() (datatypes.Container_Dns_Domain_Registration_Nameserver, error)
	GetRegistrantVerificationStatusDetail() (datatypes.Container_Dns_Domain_Registration_Registrant_Verification_StatusDetail, error)
	GetTransferInformation(domainName string) (datatypes.Container_Dns_Domain_Registration_Transfer_Information, error)
	LockDomain() (bool, error)
	LookupDomain(domainName string) (datatypes.Container_Dns_Domain_Registration_Lookup, error)
	ModifyContact(contact datatypes.Container_Dns_Domain_Registration_Contact) (bool, error)
	ModifyRegisteredNameserver(oldNameserver string, newNameserver string, ipAddress string) (bool, error)
	RegisterNameserver(nameserver string, ipAddress string) (bool, error)
	RemoveNameserversFromDomain(nameservers string) (bool, error)
	SendAuthenticationCode() (bool, error)
	SendRegistrantVerificationEmail() (bool, error)
	SendTransferApprovalEmail() (bool, error)
	SetAuthenticationCode(authenticationCode string) (bool, error)
	UnlockDomain() (bool, error)

	GetAccount() (datatypes.Account, error)
	GetDomainRegistrationStatus() (datatypes.Dns_Domain_Registration_Status, error)
	GetRegistrantVerificationStatus() (datatypes.Dns_Domain_Registration_Registrant_Verification_Status, error)
	GetServiceProvider() (datatypes.Service_Provider, error)
}

type Dns_Domain_Registration_Registrant_Verification_Status interface {
	Entity

	GetAllObjects() (datatypes.Dns_Domain_Registration_Registrant_Verification_Status, error)
	GetObject() (datatypes.Dns_Domain_Registration_Registrant_Verification_Status, error)
}

type Dns_Domain_Registration_Status interface {
	Entity

	GetAllObjects() (datatypes.Dns_Domain_Registration_Status, error)
	GetObject() (datatypes.Dns_Domain_Registration_Status, error)
}

type Dns_Domain_ResourceRecord interface {
	Entity

	CreateObject(templateObject datatypes.Dns_Domain_ResourceRecord) (datatypes.Dns_Domain_ResourceRecord, error)
	CreateObjects(templateObjects datatypes.Dns_Domain_ResourceRecord) (datatypes.Dns_Domain_ResourceRecord, error)
	DeleteObject() (bool, error)
	DeleteObjects(templateObjects datatypes.Dns_Domain_ResourceRecord) (bool, error)
	EditObject(templateObject datatypes.Dns_Domain_ResourceRecord) (bool, error)
	EditObjects(templateObjects datatypes.Dns_Domain_ResourceRecord) (bool, error)
	GetObject() (datatypes.Dns_Domain_ResourceRecord, error)

	GetDomain() (datatypes.Dns_Domain, error)
}

type Dns_Domain_ResourceRecord_AType interface {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_AaaaType interface {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_CnameType interface {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_MxType interface {
	Dns_Domain_ResourceRecord

	CreateObject(templateObject datatypes.Dns_Domain_ResourceRecord_MxType) (datatypes.Dns_Domain_ResourceRecord_MxType, error)
	CreateObjects(templateObjects datatypes.Dns_Domain_ResourceRecord) (datatypes.Dns_Domain_ResourceRecord, error)
	DeleteObject() (bool, error)
	DeleteObjects(templateObjects datatypes.Dns_Domain_ResourceRecord_MxType) (bool, error)
	EditObject(templateObject datatypes.Dns_Domain_ResourceRecord_MxType) (bool, error)
	EditObjects(templateObjects datatypes.Dns_Domain_ResourceRecord_MxType) (bool, error)
	GetObject() (datatypes.Dns_Domain_ResourceRecord_MxType, error)
}

type Dns_Domain_ResourceRecord_NsType interface {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_PtrType interface {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_SoaType interface {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_SpfType interface {
	Dns_Domain_ResourceRecord_TxtType
}

type Dns_Domain_ResourceRecord_SrvType interface {
	Dns_Domain_ResourceRecord

	CreateObject(templateObject datatypes.Dns_Domain_ResourceRecord_SrvType) (datatypes.Dns_Domain_ResourceRecord_SrvType, error)
	CreateObjects(templateObjects datatypes.Dns_Domain_ResourceRecord) (datatypes.Dns_Domain_ResourceRecord, error)
	DeleteObject() (bool, error)
	DeleteObjects(templateObjects datatypes.Dns_Domain_ResourceRecord_SrvType) (bool, error)
	EditObject(templateObject datatypes.Dns_Domain_ResourceRecord_SrvType) (bool, error)
	EditObjects(templateObjects datatypes.Dns_Domain_ResourceRecord_SrvType) (bool, error)
	GetObject() (datatypes.Dns_Domain_ResourceRecord_SrvType, error)
}

type Dns_Domain_ResourceRecord_TxtType interface {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_Reverse interface {
	Dns_Domain
}

type Dns_Domain_Reverse_Version4 interface {
	Dns_Domain_Reverse
}

type Dns_Domain_Reverse_Version6 interface {
	Dns_Domain_Reverse
}

type Dns_Message interface {
	Entity

	GetDomain() (datatypes.Dns_Domain, error)
	GetResourceRecord() (datatypes.Dns_Domain_ResourceRecord, error)
	GetSecondary() (datatypes.Dns_Secondary, error)
}

type Dns_Secondary interface {
	Entity

	ConvertToPrimary() (bool, error)
	CreateObject(templateObject datatypes.Dns_Secondary) (datatypes.Dns_Secondary, error)
	CreateObjects(templateObjects datatypes.Dns_Secondary) (datatypes.Dns_Secondary, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Dns_Secondary) (bool, error)
	GetByDomainName(name string) (datatypes.Dns_Secondary, error)
	GetObject() (datatypes.Dns_Secondary, error)
	TransferNow() (bool, error)

	GetAccount() (datatypes.Account, error)
	GetDomain() (datatypes.Dns_Domain, error)
	GetErrorMessages() (datatypes.Dns_Message, error)
	GetStatus() (datatypes.Dns_Status, error)
}

type Dns_Status interface {
	Entity
}
