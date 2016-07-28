package softlayer

type McAfee_Epolicy_Orchestrator_Version36_Agent_Details interface {
	Entity

	GetCurrentPolicy() (McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details, error)
}

type McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details interface {
	Entity

	GetCurrentPolicy() (McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details, error)
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event interface {
	Entity

	GetVirusActionTaken() (McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description, error)
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent interface {
	Entity

	GetSignature() (McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature, error)
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_BlockedApplicationEvent interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent interface {
	Entity

	GetSignature() (McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature, error)
}

type McAfee_Epolicy_Orchestrator_Version36_Policy_Object interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Product_Properties interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version45_Agent_Details interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details interface {
	Entity

	GetAgentDetails() (McAfee_Epolicy_Orchestrator_Version45_Agent_Details, error)
	GetPolicies() (McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details, error)
}

type McAfee_Epolicy_Orchestrator_Version45_Event interface {
	Entity

	GetAgentDetails() (McAfee_Epolicy_Orchestrator_Version45_Agent_Details, error)
	GetVirusActionTaken() (McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description, error)
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Version7 interface {
	McAfee_Epolicy_Orchestrator_Version45_Event

	GetSignature() (McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7, error)
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Version8 interface {
	McAfee_Epolicy_Orchestrator_Version45_Event

	GetSignature() (McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8, error)
}

type McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7 interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8 interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version45_Policy_Object interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version45_Product_Properties interface {
	Entity
}
