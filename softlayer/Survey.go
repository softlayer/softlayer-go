package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Survey interface {
	Entity

	GetActiveSurveyByType(typ string) (datatypes.Survey, error)
	GetObject() (datatypes.Survey, error)
	TakeSurvey(responses datatypes.Survey_Response) (bool, error)

	GetQuestions() (datatypes.Survey_Question, error)
	GetStatus() (datatypes.Survey_Status, error)
	GetType() (datatypes.Survey_Type, error)
}

type Survey_Answer interface {
	Entity

	GetSurveyQuestion() (datatypes.Survey_Question, error)
}

type Survey_Question interface {
	Entity

	GetAnswers() (datatypes.Survey_Answer, error)
	GetSurvey() (datatypes.Survey, error)
}

type Survey_Response interface {
	Entity

	GetSurveyAnswer() (datatypes.Survey_Answer, error)
}

type Survey_Status interface {
	Entity
}

type Survey_Type interface {
	Entity
}
