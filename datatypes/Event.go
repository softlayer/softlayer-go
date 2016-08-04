package datatypes

import "time"

type Event_Log struct {
	Entity

	AccountId       *int           `json:"accountId,omitempty"`
	EventCreateDate *time.Time     `json:"eventCreateDate,omitempty"`
	EventName       *string        `json:"eventName,omitempty"`
	IpAddress       *string        `json:"ipAddress,omitempty"`
	Label           *string        `json:"label,omitempty"`
	MetaData        *string        `json:"metaData,omitempty"`
	ObjectId        *int           `json:"objectId,omitempty"`
	ObjectName      *string        `json:"objectName,omitempty"`
	Resource        *Entity        `json:"resource,omitempty"`
	TraceId         *string        `json:"traceId,omitempty"`
	User            *User_Customer `json:"user,omitempty"`
	UserId          *int           `json:"userId,omitempty"`
	UserType        *string        `json:"userType,omitempty"`
	Username        *string        `json:"username,omitempty"`
}
