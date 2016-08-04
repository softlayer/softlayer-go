package service

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Survey struct {
	Session *Session
	Options
}

func (r *Session) GetSurveyService() Survey {
	return Survey{Session: r}
}

func (r *Survey) GetActiveSurveyByType(typ *string) (resp datatypes.Survey, err error) {
	params := []interface{}{
		typ,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Survey) GetObject() (resp datatypes.Survey, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Survey) TakeSurvey(responses []datatypes.Survey_Response) (resp bool, err error) {
	params := []interface{}{
		responses,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Survey) GetQuestions() (resp []datatypes.Survey_Question, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Survey) GetStatus() (resp datatypes.Survey_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Survey) GetType() (resp datatypes.Survey_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Survey_Answer struct {
	Session *Session
	Options
}

func (r *Session) GetSurveyAnswerService() Survey_Answer {
	return Survey_Answer{Session: r}
}

func (r *Survey_Answer) GetSurveyQuestion() (resp datatypes.Survey_Question, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Survey_Question struct {
	Session *Session
	Options
}

func (r *Session) GetSurveyQuestionService() Survey_Question {
	return Survey_Question{Session: r}
}

func (r *Survey_Question) GetAnswers() (resp []datatypes.Survey_Answer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Survey_Question) GetSurvey() (resp datatypes.Survey, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Survey_Response struct {
	Session *Session
	Options
}

func (r *Session) GetSurveyResponseService() Survey_Response {
	return Survey_Response{Session: r}
}

func (r *Survey_Response) GetSurveyAnswer() (resp datatypes.Survey_Answer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Survey_Status struct {
	Session *Session
	Options
}

func (r *Session) GetSurveyStatusService() Survey_Status {
	return Survey_Status{Session: r}
}

type Survey_Type struct {
	Session *Session
	Options
}

func (r *Session) GetSurveyTypeService() Survey_Type {
	return Survey_Type{Session: r}
}
