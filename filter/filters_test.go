/**
 * Copyright 2019 IBM Corp.
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

package filter

import (
	"testing"
)

func TestFilter(t *testing.T) {
	var actual, expected string

	expected = `{"virtualGuests":{"hostname":{"operation":"example.com"}}}`
	actual = Path("virtualGuests.hostname").Eq("example.com").Build()
	if actual != expected {
		t.Errorf("\nExpected: %s\nActual:   %s", expected, actual)
	}

	expected = `{"datacenter":{"locationName":{"operation":"Dallas"}},"id":{"operation":"134"},"something":{"creationDate":{"operation":"isDate","options":[{"name":"date","value":["01/01/01"]}]}}}`
	actual = New(
		Path("id").Eq("134"),
		Path("datacenter.locationName").Eq("Dallas"),
		Path("something.creationDate").Date("01/01/01"),
	).Build()
	if actual != expected {
		t.Errorf("\nExpected: %s\nActual:   %s", expected, actual)
	}

	expected = `{"virtualGuests":{"domain":{"operation":"example.com"},"id":{"operation":"!= 12345"}}}`
	actual = Build(
		Path("virtualGuests.domain").Eq("example.com"),
		Path("virtualGuests.id").NotEq(12345),
	)
	if actual != expected {
		t.Errorf("\nExpected: %s\nActual:   %s", expected, actual)
	}

	filters := New(
		Path("virtualGuests.hostname").StartsWith("KM078"),
		Path("virtualGuests.id").NotEq(12345),
	)

	filters = append(filters, Path("virtualGuests.domain").Eq("example.com"))

	actual = filters.Build()
	expected = `{"virtualGuests":{"domain":{"operation":"example.com"},"hostname":{"operation":"^= KM078"},"id":{"operation":"!= 12345"}}}`
	if actual != expected {
		t.Errorf("\nExpected: %s\nActual:   %s", expected, actual)
	}
}

func TestFilterOpts(t *testing.T) {
	var actual, expected string

	expected = `{"createDate":{"operation":"orderBy","options":[{"name":"sort","value":["DESC"]}]},"name":{"operation":"25GB - Ubuntu / Ubuntu / 18.04-64 Minimal for VSI"}}`

	filters := New(Path("name").Eq("25GB - Ubuntu / Ubuntu / 18.04-64 Minimal for VSI"))
	filters = append(
		filters,
		Path("createDate").Eq("orderBy").Opt("sort", [1]string{"DESC"}),
	)

	actual = filters.Build()

	if actual != expected {
		t.Errorf("\nExpected: %s\nActual:   %s", expected, actual)
	}
}
