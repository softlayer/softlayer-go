package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Security_Certificate interface {
	Entity

	CreateObject(templateObject datatypes.Security_Certificate) (datatypes.Security_Certificate, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Security_Certificate) (bool, error)
	FindByCommonName(commonName string) (datatypes.Security_Certificate, error)
	GetObject() (datatypes.Security_Certificate, error)
	GetPemFormat() (string, error)

	GetAssociatedServiceCount() (int, error)
	GetLoadBalancerVirtualIpAddresses() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, error)
}

type Security_Certificate_Entry interface {
	Entity
}

type Security_Certificate_Request interface {
	Entity

	CancelSslOrder() (bool, error)
	GetAdministratorEmailDomains(commonName string) (string, error)
	GetAdministratorEmailPrefixes() (string, error)
	GetObject() (datatypes.Security_Certificate_Request, error)
	GetPreviousOrderData() (datatypes.Container_Product_Order_Security_Certificate, error)
	GetSslCertificateRequests(accountId int) (datatypes.Security_Certificate_Request, error)
	ResendEmail(emailType string) (bool, error)
	ValidateCsr(csr string, validityMonths int, itemId int, serverType string) (bool, error)

	GetAccount() (datatypes.Account, error)
	GetCertificateAuthorityName() (string, error)
	GetOrder() (datatypes.Billing_Order, error)
	GetOrderItem() (datatypes.Billing_Order_Item, error)
	GetStatus() (datatypes.Security_Certificate_Request_Status, error)
}

type Security_Certificate_Request_ServerType interface {
	Entity

	GetAllObjects() (datatypes.Security_Certificate_Request, error)
	GetObject() (datatypes.Security_Certificate_Request_ServerType, error)
}

type Security_Certificate_Request_Status interface {
	Entity

	GetObject() (datatypes.Security_Certificate_Request_Status, error)
	GetSslRequestStatuses() (datatypes.Security_Certificate_Request_Status, error)
}

type Security_Directory_Service_Host_Xref_Hardware interface {
	Entity

	GetHost() (datatypes.Hardware, error)
}

type Security_SecureTransportCipher interface {
	Entity
}

type Security_SecureTransportProtocol interface {
	Entity
}

type Security_Ssh_Key interface {
	Entity

	CreateObject(templateObject datatypes.Security_Ssh_Key) (datatypes.Security_Ssh_Key, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Security_Ssh_Key) (bool, error)
	GetObject() (datatypes.Security_Ssh_Key, error)

	GetAccount() (datatypes.Account, error)
	GetBlockDeviceTemplateGroups() (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	GetSoftwarePasswords() (datatypes.Software_Component_Password, error)
}
