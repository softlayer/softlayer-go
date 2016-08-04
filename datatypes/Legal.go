package datatypes

type Legal_RegulatedWorkload struct {
	Entity

	Account        *Account                      `json:"account,omitempty"`
	AccountId      *int                          `json:"accountId,omitempty"`
	EnabledFlag    *bool                         `json:"enabledFlag,omitempty"`
	Id             *int                          `json:"id,omitempty"`
	Type           *Legal_RegulatedWorkload_Type `json:"type,omitempty"`
	WorkloadTypeId *int                          `json:"workloadTypeId,omitempty"`
}

type Legal_RegulatedWorkload_Type struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}
