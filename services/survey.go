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

package services

import (
	"github.ibm.com/riethm/gopherlayer/datatypes"
	"github.ibm.com/riethm/gopherlayer/session"
	"github.ibm.com/riethm/gopherlayer/sl"
)

// The SoftLayer_Survey data type contains general information relating to a single SoftLayer survey.
type Survey struct {
	Session *session.Session
	Options sl.Options
}

func GetSurveyService(sess *session.Session) Survey {
	return Survey{Session: sess}
}

func (r Survey) Id(id int) Survey {
	r.Options.Id = &id
	return r
}

func (r Survey) Mask(mask string) Survey {
	r.Options.Mask = mask
	return r
}

func (r Survey) Filter(filter string) Survey {
	r.Options.Filter = filter
	return r
}

func (r Survey) Limit(limit int) Survey {
	r.Options.Limit = &limit
	return r
}

func (r Survey) Offset(offset int) Survey {
	r.Options.Offset = &offset
	return r
}

// Provides survey details for the given type
func (r Survey) GetActiveSurveyByType(typ *string) (resp datatypes.Survey, err error) {
	params := []interface{}{
		typ,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Survey object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Survey service. You can only retrieve the survey that your portal user has taken.
func (r Survey) GetObject() (resp datatypes.Survey, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The questions for a survey.
func (r Survey) GetQuestions() (resp []datatypes.Survey_Question, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The status of the survey
func (r Survey) GetStatus() (resp datatypes.Survey_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The type of survey
func (r Survey) GetType() (resp datatypes.Survey_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Response to a SoftLayer survey's questions.
func (r Survey) TakeSurvey(responses []datatypes.Survey_Response) (resp bool, err error) {
	params := []interface{}{
		responses,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
