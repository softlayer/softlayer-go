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

package datatypes

import "time"

// Void is a dummy type for identifying void return values from methods
type Void int

// Time type overrides the default json marshaler with the SoftLayer custom format
type Time struct {
	time.Time
}

func (r Time) String() (string) {
	return r.Time.Format(time.RFC3339)
}

// MarshalJSON returns the json encoding of the datatypes.Time receiver.  This
// override is necessary to ensure datetimes are formatted in the way SoftLayer
// expects - that is, using the RFC3339 format, without nanoseconds.
func (r Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + r.String() + `"`), nil
}

// MarshalText returns a text encoding of the datatypes.Time receiver.  This
// is mainly provided to complete what might be expected of a type that
// implements the Marshaler interface.
func (r Time) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}