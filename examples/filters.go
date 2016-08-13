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

package main

import (
	"fmt"
	"github.ibm.com/riethm/gopherlayer/filter"
)

func main() {
	fmt.Println(
		filter.New(
			filter.Path("id", "134").Eq(),
			filter.Path("datacenter.locationName", "Dallas").Eq(),
			filter.Path("something.creationDate").Date("01/01/01"),
		).Build(),
	)

	fmt.Println(
		filter.Build(
			filter.Path("virtualGuests.domain", "example.com").Eq(),
			filter.Path("virtualGuests.id", 12345).NotEq(),
		),
	)

	filters := filter.New(
		filter.Path("virtualGuests.hostname", "KM078").StartsWith(),
		filter.Path("virtualGuests.id", 12345).NotEq(),
	)

	filters = append(filters, filter.Path("virtualGuests.domain", "example.com").Eq())

	fmt.Println(filters.Build())
}
