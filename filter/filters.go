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

// See reference at https://sldn.softlayer.com/article/object-filters.
// Examples in the README.md file and in the examples directory.
package filter

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Filter struct {
	Path string
	Op   string
	Opts map[string]interface{}
	Val  interface{}
}

type Filters []Filter

// Returns an array of Filters that you can later call .Build() on.
func New(args ...Filter) Filters {
	return args
}

// This is like calling New().Build().
// Returns a JSON string that can be used as the object filter.
func Build(args ...Filter) string {
	filters := Filters{}

	for _, arg := range args {
		filters = append(filters, arg)
	}

	return filters.Build()
}

// This creates a new Filter. The first argument (path) is required.
// The second argument (val) is optional.
func Path(path string, val ...interface{}) Filter {
	if len(val) > 0 {
		return Filter{Path: path, Val: val[0]}
	}

	return Filter{Path: path}
}

// Builds the filter string in JSON format
func (fs Filters) Build() string {
	// Loops around filters,
	// splitting path on '.' and looping around path pieces.
	// Idea is to create a map/tree like map[string]interface{}.
	// Every component in the path is a node to create in the tree.
	// Once we get to the leaf, we set the operation.
	// map[string]interface{}{"operation": op+" "+value}
	// If Op is "", then just map[string]interface{}{"operation": value}.
	// Afterwards, the Opts are traversed; []map[string]interface{}{}
	// For every entry in Opts, we create one map, and append it to an array of maps.
	// At the end, json.Marshal the whole thing.
	result := map[string]interface{}{}
	for _, filter := range fs {
		if filter.Path == "" {
			continue
		}

		cursor := result
		nodes := strings.Split(filter.Path, ".")
		for len(nodes) > 1 {
			branch := nodes[0]
			if cursor[branch] == nil {
				cursor[branch] = map[string]interface{}{}
			}
			cursor = cursor[branch].(map[string]interface{})
			nodes = nodes[1:len(nodes)]
		}

		leaf := nodes[0]
		if filter.Val != nil {
			operation := filter.Val
			if filter.Op != "" {
				format := "%s"
				if reflect.TypeOf(filter.Val).String() != "string" {
					format = "%d"
				}
				operation = filter.Op + " " + fmt.Sprintf(format, filter.Val)
			}

			cursor[leaf] = map[string]interface{}{
				"operation": operation,
			}
		}

		if filter.Opts == nil {
			continue
		}

		options := []map[string]interface{}{}
		for name, value := range filter.Opts {
			options = append(options, map[string]interface{}{
				"name":  name,
				"value": value,
			})
		}

		cursor[leaf] = map[string]interface{}{
			"operation": filter.Op,
			"options":   options,
		}
	}

	jsonStr, _ := json.Marshal(result)
	return string(jsonStr)
}

// Add options to the filter. Can be chained for multiple options.
func (f Filter) Opt(name string, value interface{}) Filter {
	if f.Opts == nil {
		f.Opts = map[string]interface{}{}
	}

	f.Opts[name] = value
	return f
}

// Set this filter to test if property is equal to the value
func (f Filter) Eq() Filter {
	f.Op = ""
	return f
}

// Set this filter to test if property is not equal to the value
func (f Filter) NotEq() Filter {
	f.Op = "!="
	return f
}

// Set this filter to test if property is like the value
func (f Filter) Like() Filter {
	f.Op = "~"
	return f
}

// Set this filter to test if property is unlike value
func (f Filter) NotLike() Filter {
	f.Op = "!~"
	return f
}

// Set this filter to test if property is less than value
func (f Filter) LessThan() Filter {
	f.Op = "<"
	return f
}

// Set this filter to test if property is less than or equal to the value
func (f Filter) LessThanOrEqual() Filter {
	f.Op = "<="
	return f
}

// Set this filter to test if property is greater than value
func (f Filter) GreaterThan() Filter {
	f.Op = ">"
	return f
}

// Set this filter to test if property is greater than or equal to value
func (f Filter) GreaterThanOrEqual() Filter {
	f.Op = ">="
	return f
}

// Set this filter to test if property is null
func (f Filter) IsNull() Filter {
	f.Op = "is null"
	f.Val = nil
	return f
}

// Set this filter to test if property is not null
func (f Filter) NotNull() Filter {
	f.Op = "not null"
	f.Val = nil
	return f
}

// Set this filter to test if property contains the value
func (f Filter) Contains() Filter {
	f.Op = "*="
	return f
}

// Set this filter to test if property does not contain the value
func (f Filter) NotContains() Filter {
	f.Op = "!*="
	return f
}

// Set this filter to test if property starts with the value
func (f Filter) StartsWith() Filter {
	f.Op = "^="
	return f
}

// Set this filter to test if property does not start with the value
func (f Filter) NotStartsWith() Filter {
	f.Op = "!^="
	return f
}

// Set this filter to test if property ends with the value
func (f Filter) EndsWith() Filter {
	f.Op = "$="
	return f
}

// Set this filter to test if property does not end with the value
func (f Filter) NotEndsWith() Filter {
	f.Op = "!$="
	return f
}

// Set this filter to test if property is one of the values in args.
func (f Filter) In(args ...interface{}) Filter {
	f.Op = "in"
	values := []interface{}{}
	for _, arg := range args {
		values = append(values, arg)
	}

	return f.Opt("data", values)
}

// Set this filter to test if property has a date older than the value in days.
func (f Filter) DaysPast() Filter {
	f.Op = ">= currentDate -"
	return f
}

// Set this filter to test if property has the exact date as the value.
func (f Filter) Date(date string) Filter {
	f.Op = "isDate"
	f.Val = nil
	return f.Opt("date", date)
}

// Set this filter to test if property has a date before the value.
func (f Filter) DateBefore(date string) Filter {
	f.Op = "lessThanDate"
	f.Val = nil
	return f.Opt("date", date)
}

// Set this filter to test if property has a date after the value.
func (f Filter) DateAfter(date string) Filter {
	f.Op = "greaterThanDate"
	f.Val = nil
	return f.Opt("date", date)
}

// Set this filter to test if property has a date between the values.
func (f Filter) DateBetween(start string, end string) Filter {
	f.Op = "betweenDate"
	f.Val = nil
	return f.Opt("startDate", start).Opt("endDate", end)
}
