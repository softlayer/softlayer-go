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

package service

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Marketplace_EmailDistribution struct {
	Session *Session
	Options
}

func (r *Session) GetMarketplaceEmailDistributionService() Marketplace_EmailDistribution {
	return Marketplace_EmailDistribution{Session: r}
}

type Marketplace_Partner struct {
	Session *Session
	Options
}

func (r *Session) GetMarketplacePartnerService() Marketplace_Partner {
	return Marketplace_Partner{Session: r}
}

func (r *Marketplace_Partner) GetAllObjects() (resp []datatypes.Marketplace_Partner, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Marketplace_Partner) GetAllPublishedPartners(searchTerm *string) (resp []datatypes.Marketplace_Partner, err error) {
	params := []interface{}{
		searchTerm,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Marketplace_Partner) GetFeaturedPartners(non *bool) (resp []datatypes.Marketplace_Partner, err error) {
	params := []interface{}{
		non,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Marketplace_Partner) GetFile(name *string) (resp datatypes.Marketplace_Partner_File, err error) {
	params := []interface{}{
		name,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Marketplace_Partner) GetObject() (resp datatypes.Marketplace_Partner, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Marketplace_Partner) GetPartnerByUrlIdentifier(urlIdentifier *string) (resp datatypes.Marketplace_Partner, err error) {
	params := []interface{}{
		urlIdentifier,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Marketplace_Partner) GetAttachments() (resp []datatypes.Marketplace_Partner_Attachment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Marketplace_Partner) GetLogoMedium() (resp datatypes.Marketplace_Partner_Attachment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Marketplace_Partner) GetLogoMediumTemp() (resp datatypes.Marketplace_Partner_Attachment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Marketplace_Partner) GetLogoSmall() (resp datatypes.Marketplace_Partner_Attachment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Marketplace_Partner) GetLogoSmallTemp() (resp datatypes.Marketplace_Partner_Attachment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Marketplace_Partner_Attachment struct {
	Session *Session
	Options
}

func (r *Session) GetMarketplacePartnerAttachmentService() Marketplace_Partner_Attachment {
	return Marketplace_Partner_Attachment{Session: r}
}

func (r *Marketplace_Partner_Attachment) GetAttachmentType() (resp datatypes.Marketplace_Partner_Attachment_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Marketplace_Partner_Attachment_Type struct {
	Session *Session
	Options
}

func (r *Session) GetMarketplacePartnerAttachmentTypeService() Marketplace_Partner_Attachment_Type {
	return Marketplace_Partner_Attachment_Type{Session: r}
}

type Marketplace_Partner_File struct {
	Session *Session
	Options
}

func (r *Session) GetMarketplacePartnerFileService() Marketplace_Partner_File {
	return Marketplace_Partner_File{Session: r}
}

type Marketplace_Partner_File_Attributes struct {
	Session *Session
	Options
}

func (r *Session) GetMarketplacePartnerFileAttributesService() Marketplace_Partner_File_Attributes {
	return Marketplace_Partner_File_Attributes{Session: r}
}
