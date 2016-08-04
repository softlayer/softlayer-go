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
