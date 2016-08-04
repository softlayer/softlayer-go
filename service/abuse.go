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

type Abuse_Lockdown_Resource struct {
	Session *Session
	Options
}

func (r *Session) GetAbuseLockdownResourceService() Abuse_Lockdown_Resource {
	return Abuse_Lockdown_Resource{Session: r}
}

func (r *Abuse_Lockdown_Resource) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Abuse_Lockdown_Resource) GetInvoiceItem() (resp datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
