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

package sl

import (
	"github.ibm.com/riethm/gopherlayer.git/datatypes"
	"time"
)

// Convenience functions for returning pointers to values

// Int returns a pointer to the int value provided
func Int(v int) *int {
	return &v
}

// String returns a pointer to the string value provided
func String(v string) *string {
	return &v
}

// Bool returns a pointer to the bool value provided
func Bool(v bool) *bool {
	return &v
}

// Time converts the time.Time value provided to a datatypes.Time value,
// and returns a pointer to it
func Time(v time.Time) *datatypes.Time {
	r := datatypes.Time{Time: v}
	return &r
}
