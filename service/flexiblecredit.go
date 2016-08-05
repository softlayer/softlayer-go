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

type FlexibleCredit_Program struct {
	Session *Session
	Options
}

func (r *Session) GetFlexibleCreditProgramService() FlexibleCredit_Program {
	return FlexibleCredit_Program{Session: r}
}

func (r *FlexibleCredit_Program) GetAffiliatesAvailableForSelfEnrollmentByVerificationType(verificationTypeKeyName *string) (resp []datatypes.FlexibleCredit_Affiliate, err error) {
	params := []interface{}{
		verificationTypeKeyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *FlexibleCredit_Program) GetCompanyTypes() (resp []datatypes.FlexibleCredit_Company_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *FlexibleCredit_Program) GetObject() (resp datatypes.FlexibleCredit_Program, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *FlexibleCredit_Program) SelfEnrollNewAccount(accountTemplate *datatypes.Account) (resp datatypes.Account, err error) {
	params := []interface{}{
		accountTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
