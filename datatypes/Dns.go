/**
 * Copyright 2016 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package datatypes

type Dns_Domain struct {
	Entity

	Account             *Account                    `json:"account,omitempty"`
	Id                  *int                        `json:"id,omitempty"`
	ManagedResourceFlag *bool                       `json:"managedResourceFlag,omitempty"`
	Name                *string                     `json:"name,omitempty"`
	ResourceRecordCount *uint                       `json:"resourceRecordCount,omitempty"`
	ResourceRecords     []Dns_Domain_ResourceRecord `json:"resourceRecords,omitempty"`
	Secondary           *Dns_Secondary              `json:"secondary,omitempty"`
	Serial              *int                        `json:"serial,omitempty"`
	UpdateDate          *Time                       `json:"updateDate,omitempty"`
}

type Dns_Domain_Forward struct {
	Dns_Domain
}

type Dns_Domain_Registration struct {
	Entity

	Account                        *Account                                                `json:"account,omitempty"`
	CreateDate                     *Time                                                   `json:"createDate,omitempty"`
	DomainRegistrationStatus       *Dns_Domain_Registration_Status                         `json:"domainRegistrationStatus,omitempty"`
	DomainRegistrationStatusId     *int                                                    `json:"domainRegistrationStatusId,omitempty"`
	ExpireDate                     *Time                                                   `json:"expireDate,omitempty"`
	Id                             *int                                                    `json:"id,omitempty"`
	LockedFlag                     *int                                                    `json:"lockedFlag,omitempty"`
	ModifyDate                     *Time                                                   `json:"modifyDate,omitempty"`
	Name                           *string                                                 `json:"name,omitempty"`
	RegistrantVerificationStatus   *Dns_Domain_Registration_Registrant_Verification_Status `json:"registrantVerificationStatus,omitempty"`
	RegistrantVerificationStatusId *int                                                    `json:"registrantVerificationStatusId,omitempty"`
	ServiceProvider                *Service_Provider                                       `json:"serviceProvider,omitempty"`
	ServiceProviderId              *int                                                    `json:"serviceProviderId,omitempty"`
}

type Dns_Domain_Registration_Registrant_Verification_Status struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Dns_Domain_Registration_Status struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Dns_Domain_ResourceRecord struct {
	Entity

	Data              *string     `json:"data,omitempty"`
	Domain            *Dns_Domain `json:"domain,omitempty"`
	DomainId          *int        `json:"domainId,omitempty"`
	Expire            *int        `json:"expire,omitempty"`
	Host              *string     `json:"host,omitempty"`
	Id                *int        `json:"id,omitempty"`
	Minimum           *int        `json:"minimum,omitempty"`
	MxPriority        *int        `json:"mxPriority,omitempty"`
	Refresh           *int        `json:"refresh,omitempty"`
	ResponsiblePerson *string     `json:"responsiblePerson,omitempty"`
	Retry             *int        `json:"retry,omitempty"`
	Ttl               *int        `json:"ttl,omitempty"`
	Type              *string     `json:"type,omitempty"`
}

type Dns_Domain_ResourceRecord_AType struct {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_AaaaType struct {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_CnameType struct {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_MxType struct {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_NsType struct {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_PtrType struct {
	Dns_Domain_ResourceRecord

	IsGatewayAddress *bool `json:"isGatewayAddress,omitempty"`
}

type Dns_Domain_ResourceRecord_SoaType struct {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_SpfType struct {
	Dns_Domain_ResourceRecord_TxtType
}

type Dns_Domain_ResourceRecord_SrvType struct {
	Dns_Domain_ResourceRecord

	Port     *int    `json:"port,omitempty"`
	Priority *int    `json:"priority,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	Service  *string `json:"service,omitempty"`
	Weight   *int    `json:"weight,omitempty"`
}

type Dns_Domain_ResourceRecord_TxtType struct {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_Reverse struct {
	Dns_Domain

	NetworkAddress *string `json:"networkAddress,omitempty"`
}

type Dns_Domain_Reverse_Version4 struct {
	Dns_Domain_Reverse
}

type Dns_Domain_Reverse_Version6 struct {
	Dns_Domain_Reverse
}

type Dns_Message struct {
	Entity

	CreateDate     *Time                      `json:"createDate,omitempty"`
	Domain         *Dns_Domain                `json:"domain,omitempty"`
	Id             *int                       `json:"id,omitempty"`
	Message        *string                    `json:"message,omitempty"`
	Priority       *string                    `json:"priority,omitempty"`
	ResourceRecord *Dns_Domain_ResourceRecord `json:"resourceRecord,omitempty"`
	Secondary      *Dns_Secondary             `json:"secondary,omitempty"`
}

type Dns_Secondary struct {
	Entity

	Account           *Account      `json:"account,omitempty"`
	CreateDate        *Time         `json:"createDate,omitempty"`
	Domain            *Dns_Domain   `json:"domain,omitempty"`
	ErrorMessageCount *uint         `json:"errorMessageCount,omitempty"`
	ErrorMessages     []Dns_Message `json:"errorMessages,omitempty"`
	Id                *int          `json:"id,omitempty"`
	LastUpdate        *Time         `json:"lastUpdate,omitempty"`
	MasterIpAddress   *string       `json:"masterIpAddress,omitempty"`
	Status            *Dns_Status   `json:"status,omitempty"`
	StatusId          *int          `json:"statusId,omitempty"`
	StatusText        *string       `json:"statusText,omitempty"`
	TransferFrequency *int          `json:"transferFrequency,omitempty"`
	ZoneName          *string       `json:"zoneName,omitempty"`
}

type Dns_Status struct {
	Entity

	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
