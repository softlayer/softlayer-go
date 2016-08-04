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

type Marketplace_EmailDistribution struct {
	Entity

	Email *string `json:"email,omitempty"`
	Id    *int    `json:"id,omitempty"`
}

type Marketplace_Partner struct {
	Entity

	AccountId               *int                             `json:"accountId,omitempty"`
	AttachedFiles           []Marketplace_Partner_Attachment `json:"attachedFiles,omitempty"`
	AttachmentCount         *uint                            `json:"attachmentCount,omitempty"`
	Attachments             []Marketplace_Partner_Attachment `json:"attachments,omitempty"`
	CompanyDescription      *string                          `json:"companyDescription,omitempty"`
	CompanyName             *string                          `json:"companyName,omitempty"`
	HeadlineDescription     *string                          `json:"headlineDescription,omitempty"`
	Id                      *int                             `json:"id,omitempty"`
	LinkFreeTrial           *string                          `json:"linkFreeTrial,omitempty"`
	LinkOrderPage           *string                          `json:"linkOrderPage,omitempty"`
	LinkWebsite             *string                          `json:"linkWebsite,omitempty"`
	LogoMedium              *Marketplace_Partner_Attachment  `json:"logoMedium,omitempty"`
	LogoMediumTemp          *Marketplace_Partner_Attachment  `json:"logoMediumTemp,omitempty"`
	LogoSmall               *Marketplace_Partner_Attachment  `json:"logoSmall,omitempty"`
	LogoSmallTemp           *Marketplace_Partner_Attachment  `json:"logoSmallTemp,omitempty"`
	MetaDescription         *string                          `json:"metaDescription,omitempty"`
	MetaKeywords            *string                          `json:"metaKeywords,omitempty"`
	ProductBenefits         *string                          `json:"productBenefits,omitempty"`
	ProductCategoryId       *int                             `json:"productCategoryId,omitempty"`
	ProductDescriptionLong  *string                          `json:"productDescriptionLong,omitempty"`
	ProductDescriptionShort *string                          `json:"productDescriptionShort,omitempty"`
	ProductFeatures         *string                          `json:"productFeatures,omitempty"`
	ProductName             *string                          `json:"productName,omitempty"`
	ProductTitle            *string                          `json:"productTitle,omitempty"`
	UrlIdentifier           *string                          `json:"urlIdentifier,omitempty"`
}

type Marketplace_Partner_Attachment struct {
	Entity

	AttachmentType       *Marketplace_Partner_Attachment_Type `json:"attachmentType,omitempty"`
	AttachmentTypeId     *int                                 `json:"attachmentTypeId,omitempty"`
	BaseName             *string                              `json:"baseName,omitempty"`
	DisplayName          *string                              `json:"displayName,omitempty"`
	FileName             *string                              `json:"fileName,omitempty"`
	Id                   *int                                 `json:"id,omitempty"`
	MarketplacePartnerId *int                                 `json:"marketplacePartnerId,omitempty"`
	SaveAsName           *string                              `json:"saveAsName,omitempty"`
}

type Marketplace_Partner_Attachment_Type struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
	Type    *string `json:"type,omitempty"`
}

type Marketplace_Partner_File struct {
	Entity

	Attributes *Marketplace_Partner_File_Attributes `json:"attributes,omitempty"`
	Contents   *[]byte                              `json:"contents,omitempty"`
}

type Marketplace_Partner_File_Attributes struct {
	Entity

	Bits           *int    `json:"bits,omitempty"`
	Channels       *int    `json:"channels,omitempty"`
	Height         *int    `json:"height,omitempty"`
	HtmlAttributes *string `json:"htmlAttributes,omitempty"`
	ImageType      *int    `json:"imageType,omitempty"`
	IsImage        *bool   `json:"isImage,omitempty"`
	MimeType       *string `json:"mimeType,omitempty"`
	Width          *int    `json:"width,omitempty"`
}
