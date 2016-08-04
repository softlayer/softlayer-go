package datatypes

import "time"

type Survey struct {
	Entity

	Active        *int              `json:"active,omitempty"`
	CreateDate    *time.Time        `json:"createDate,omitempty"`
	Id            *int              `json:"id,omitempty"`
	Name          *string           `json:"name,omitempty"`
	QuestionCount *uint             `json:"questionCount,omitempty"`
	Questions     []Survey_Question `json:"questions,omitempty"`
	Status        *Survey_Status    `json:"status,omitempty"`
	StatusId      *int              `json:"statusId,omitempty"`
	Type          *Survey_Type      `json:"type,omitempty"`
	TypeId        *int              `json:"typeId,omitempty"`
}

type Survey_Answer struct {
	Entity

	Answer           *string          `json:"answer,omitempty"`
	AnswerOrder      *int             `json:"answerOrder,omitempty"`
	Id               *int             `json:"id,omitempty"`
	SurveyQuestion   *Survey_Question `json:"surveyQuestion,omitempty"`
	SurveyQuestionId *int             `json:"surveyQuestionId,omitempty"`
}

type Survey_Question struct {
	Entity

	AnswerCount   *uint           `json:"answerCount,omitempty"`
	Answers       []Survey_Answer `json:"answers,omitempty"`
	Id            *int            `json:"id,omitempty"`
	IsRequired    *int            `json:"isRequired,omitempty"`
	MultiAnswer   *int            `json:"multiAnswer,omitempty"`
	Question      *string         `json:"question,omitempty"`
	QuestionOrder *int            `json:"questionOrder,omitempty"`
	Survey        *Survey         `json:"survey,omitempty"`
	SurveyId      *int            `json:"surveyId,omitempty"`
}

type Survey_Response struct {
	Entity

	OtherAnswer    *string        `json:"otherAnswer,omitempty"`
	SurveyAnswer   *Survey_Answer `json:"surveyAnswer,omitempty"`
	SurveyAnswerId *int           `json:"surveyAnswerId,omitempty"`
}

type Survey_Status struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Survey_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
}
