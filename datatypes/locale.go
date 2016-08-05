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

type Locale struct {
	Entity

	FriendlyName *string `json:"friendlyName,omitempty"`
	Id           *int    `json:"id,omitempty"`
	LanguageTag  *string `json:"languageTag,omitempty"`
	Name         *string `json:"name,omitempty"`
}

type Locale_Country struct {
	Entity

	IsEuropeanUnionFlag *int                   `json:"isEuropeanUnionFlag,omitempty"`
	LongName            *string                `json:"longName,omitempty"`
	ShortName           *string                `json:"shortName,omitempty"`
	StateCount          *uint                  `json:"stateCount,omitempty"`
	States              []Locale_StateProvince `json:"states,omitempty"`
}

type Locale_StateProvince struct {
	Entity

	LongName  *string `json:"longName,omitempty"`
	ShortName *string `json:"shortName,omitempty"`
}

type Locale_Timezone struct {
	Entity

	Id        *int    `json:"id,omitempty"`
	LongName  *string `json:"longName,omitempty"`
	Name      *string `json:"name,omitempty"`
	Offset    *string `json:"offset,omitempty"`
	ShortName *string `json:"shortName,omitempty"`
}
