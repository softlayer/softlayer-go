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

/**
 * AUTOMATICALLY GENERATED CODE - DO NOT MODIFY
 */

package datatypes

// no documentation yet
type Marketplace_EmailDistribution struct {
	Entity

	// no documentation yet
	Email *string `json:"email,omitempty"`

	// no documentation yet
	Id *int `json:"id,omitempty"`
}

// no documentation yet
type Marketplace_Partner struct {
	Entity

	// no documentation yet
	AccountId *int `json:"accountId,omitempty"`

	// no documentation yet
	AttachedFiles []Marketplace_Partner_Attachment `json:"attachedFiles,omitempty"`

	// A count of
	AttachmentCount *uint `json:"attachmentCount,omitempty"`

	// no documentation yet
	Attachments []Marketplace_Partner_Attachment `json:"attachments,omitempty"`

	// no documentation yet
	CompanyDescription *string `json:"companyDescription,omitempty"`

	// no documentation yet
	CompanyName *string `json:"companyName,omitempty"`

	// no documentation yet
	HeadlineDescription *string `json:"headlineDescription,omitempty"`

	// no documentation yet
	Id *int `json:"id,omitempty"`

	// no documentation yet
	LinkFreeTrial *string `json:"linkFreeTrial,omitempty"`

	// no documentation yet
	LinkOrderPage *string `json:"linkOrderPage,omitempty"`

	// no documentation yet
	LinkWebsite *string `json:"linkWebsite,omitempty"`

	// no documentation yet
	LogoMedium *Marketplace_Partner_Attachment `json:"logoMedium,omitempty"`

	// no documentation yet
	LogoMediumTemp *Marketplace_Partner_Attachment `json:"logoMediumTemp,omitempty"`

	// no documentation yet
	LogoSmall *Marketplace_Partner_Attachment `json:"logoSmall,omitempty"`

	// no documentation yet
	LogoSmallTemp *Marketplace_Partner_Attachment `json:"logoSmallTemp,omitempty"`

	// no documentation yet
	MetaDescription *string `json:"metaDescription,omitempty"`

	// no documentation yet
	MetaKeywords *string `json:"metaKeywords,omitempty"`

	// no documentation yet
	ProductBenefits *string `json:"productBenefits,omitempty"`

	// no documentation yet
	ProductCategoryId *int `json:"productCategoryId,omitempty"`

	// no documentation yet
	ProductDescriptionLong *string `json:"productDescriptionLong,omitempty"`

	// no documentation yet
	ProductDescriptionShort *string `json:"productDescriptionShort,omitempty"`

	// no documentation yet
	ProductFeatures *string `json:"productFeatures,omitempty"`

	// no documentation yet
	ProductName *string `json:"productName,omitempty"`

	// no documentation yet
	ProductTitle *string `json:"productTitle,omitempty"`

	// no documentation yet
	UrlIdentifier *string `json:"urlIdentifier,omitempty"`
}

// no documentation yet
type Marketplace_Partner_Attachment struct {
	Entity

	// no documentation yet
	AttachmentType *Marketplace_Partner_Attachment_Type `json:"attachmentType,omitempty"`

	// no documentation yet
	AttachmentTypeId *int `json:"attachmentTypeId,omitempty"`

	// no documentation yet
	BaseName *string `json:"baseName,omitempty"`

	// no documentation yet
	DisplayName *string `json:"displayName,omitempty"`

	// no documentation yet
	FileName *string `json:"fileName,omitempty"`

	// no documentation yet
	Id *int `json:"id,omitempty"`

	// no documentation yet
	MarketplacePartnerId *int `json:"marketplacePartnerId,omitempty"`

	// no documentation yet
	SaveAsName *string `json:"saveAsName,omitempty"`
}

// no documentation yet
type Marketplace_Partner_Attachment_Type struct {
	Entity

	// no documentation yet
	Id *int `json:"id,omitempty"`

	// no documentation yet
	KeyName *string `json:"keyName,omitempty"`

	// no documentation yet
	Type *string `json:"type,omitempty"`
}

// no documentation yet
type Marketplace_Partner_File struct {
	Entity

	// no documentation yet
	Attributes *Marketplace_Partner_File_Attributes `json:"attributes,omitempty"`

	// no documentation yet
	Contents *[]byte `json:"contents,omitempty"`
}

// no documentation yet
type Marketplace_Partner_File_Attributes struct {
	Entity

	// no documentation yet
	Bits *int `json:"bits,omitempty"`

	// no documentation yet
	Channels *int `json:"channels,omitempty"`

	// no documentation yet
	Height *int `json:"height,omitempty"`

	// no documentation yet
	HtmlAttributes *string `json:"htmlAttributes,omitempty"`

	// no documentation yet
	ImageType *int `json:"imageType,omitempty"`

	// no documentation yet
	IsImage *bool `json:"isImage,omitempty"`

	// no documentation yet
	MimeType *string `json:"mimeType,omitempty"`

	// no documentation yet
	Width *int `json:"width,omitempty"`
}
