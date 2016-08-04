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

type Legal_RegulatedWorkload struct {
	Session *Session
	Options
}

func (r *Session) GetLegalRegulatedWorkloadService() Legal_RegulatedWorkload {
	return Legal_RegulatedWorkload{Session: r}
}

func (r *Legal_RegulatedWorkload) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Legal_RegulatedWorkload) GetType() (resp datatypes.Legal_RegulatedWorkload_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Legal_RegulatedWorkload_Type struct {
	Session *Session
	Options
}

func (r *Session) GetLegalRegulatedWorkloadTypeService() Legal_RegulatedWorkload_Type {
	return Legal_RegulatedWorkload_Type{Session: r}
}
