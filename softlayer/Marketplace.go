package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Marketplace_EmailDistribution interface {
	Entity
}

type Marketplace_Partner interface {
	Entity

	GetAllObjects() (datatypes.Marketplace_Partner, error)
	GetAllPublishedPartners(searchTerm string) (datatypes.Marketplace_Partner, error)
	GetFeaturedPartners(non bool) (datatypes.Marketplace_Partner, error)
	GetFile(name string) (datatypes.Marketplace_Partner_File, error)
	GetObject() (datatypes.Marketplace_Partner, error)
	GetPartnerByUrlIdentifier(urlIdentifier string) (datatypes.Marketplace_Partner, error)

	GetAttachments() (datatypes.Marketplace_Partner_Attachment, error)
	GetLogoMedium() (datatypes.Marketplace_Partner_Attachment, error)
	GetLogoMediumTemp() (datatypes.Marketplace_Partner_Attachment, error)
	GetLogoSmall() (datatypes.Marketplace_Partner_Attachment, error)
	GetLogoSmallTemp() (datatypes.Marketplace_Partner_Attachment, error)
}

type Marketplace_Partner_Attachment interface {
	Entity

	GetAttachmentType() (datatypes.Marketplace_Partner_Attachment_Type, error)
}

type Marketplace_Partner_Attachment_Type interface {
	Entity
}

type Marketplace_Partner_File interface {
	Entity
}

type Marketplace_Partner_File_Attributes interface {
	Entity
}
