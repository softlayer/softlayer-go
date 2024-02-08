package generator

import (
	"fmt"
	"strings"
	"text/template"
)

const license = `/**
 * Copyright 2016-2024 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with 
 * the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed 
 * on an "AS IS" BASIS,WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 */
 `
const codegenWarning = `// AUTOMATICALLY GENERATED CODE - DO NOT MODIFY`

// Define custom template functions
var fMap = template.FuncMap{
	"convertType":     ConvertType,         // Converts SoftLayer types to Go types
	"removePrefix":    RemovePrefix,        // Remove 'SoftLayer_' prefix. if it exists
	"removeReserved":  RemoveReservedWords, // Substitute language-reserved identifiers
	"titleCase":       strings.Title,       // TitleCase the argument
	"desnake":         Desnake,             // Remove '_' from Snake_Case
	"goDoc":           GoDoc,               // Format a go doc string
	"tags":            Tags,                // Remove omitempty tags if required
	"phraseMethodArg": phraseMethodArg,     // Get proper phrase for method argument
	"getTypePrefix":   getTypePrefix,       // resturns [], * or nothing.
	"deprecatedDoc":   deprecatedDoc,       // Marks things as deprecated if needed.
	"nilParam":        NilParam,            // For unit testing fakes
}

var datatype = fmt.Sprintf(`%s

%s

package datatypes

{{range .}}{{.TypeDoc|goDoc}}
type {{.Name|removePrefix}} struct {
	{{.Base|removePrefix}}

{{$base := .Name}}{{range .Properties}}{{.Doc|goDoc}}{{deprecatedDoc .Deprecated}}`+
	`{{$thisType := convertType .Type "datatypes" $base .Name}}
	{{.Name|titleCase}} {{getTypePrefix .TypeArray $thisType}}{{$thisType}}`+
	"`json:\"{{.Name|tags}}\" xmlrpc:\"{{.Name|tags}}\"`"+`

{{end}}
}

{{end}}
`, license, codegenWarning)

var services = fmt.Sprintf(`%s

%s

package services

import (
	"fmt"
	"strings"
)

{{range .}}{{$base := .Name|removePrefix}}{{.TypeDoc|goDoc}}{{deprecatedDoc .Deprecated}}
	type {{$base}} struct {
		Session session.SLSession
		Options sl.Options
	}

	// Get{{$base | desnake}}Service returns an instance of the {{$base}} SoftLayer service
	func Get{{$base | desnake}}Service(sess session.SLSession) {{$base}} {
		return {{$base}}{Session: sess}
	}

	func (r {{$base}}) Id(id int) {{$base}} {
		r.Options.Id = &id
		return r
	}

	func (r {{$base}}) Mask(mask string) {{$base}} {
		if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
			mask = fmt.Sprintf("mask[%%s]", mask)
		}

		r.Options.Mask = mask
		return r
	}

	func (r {{$base}}) Filter(filter string) {{$base}} {
		r.Options.Filter = filter
		return r
	}

	func (r {{$base}}) Limit(limit int) {{$base}} {
		r.Options.Limit = &limit
		return r
	}

	func (r {{$base}}) Offset(offset int) {{$base}} {
		r.Options.Offset = &offset
		return r
	}

	{{$rawBase := .Name}}{{range .Methods}}{{$methodName := .Name}}{{.Doc|goDoc}}{{deprecatedDoc .Deprecated}}
	func (r {{$base}}) {{.Name|titleCase}}({{range .Parameters}}{{phraseMethodArg $methodName .Name .TypeArray .Type}}{{end}}) ({{if .Type|ne "void"}}resp {{if .TypeArray}}[]{{end}}{{convertType .Type "services"}}, {{end}}err error) {
		{{if .Type|eq "void"}}var resp datatypes.Void
		{{end}}{{if or (eq .Name "placeOrder") (eq .Name "verifyOrder")}}err = datatypes.SetComplexType(orderData)
		if err != nil {
			return
		}
		{{end}}{{if len .Parameters | lt 0}}params := []interface{}{
			{{range .Parameters}}{{.Name|removeReserved}},
			{{end}}
		}
		{{end}}err = r.Session.DoRequest("{{$rawBase}}", "{{.Name}}", {{if len .Parameters | lt 0}}params{{else}}nil{{end}}, &r.Options, &resp)
	return
	}
	{{if .TypeArray}}
	func (r {{$base}}) {{.Name|titleCase}}Iter({{range .Parameters}}{{phraseMethodArg $methodName .Name .TypeArray .Type}}{{end}}) ({{if .Type|ne "void"}}resp {{if .TypeArray}}[]{{end}}{{convertType .Type "services"}}, {{end}}err error) {
		{{if len .Parameters | lt 0}}params := []interface{}{
			{{range .Parameters}}{{.Name|removeReserved}},
			{{end}}
		}
		{{end}}limit := r.Options.ValidateLimit()
		err = r.Session.DoRequest("{{$rawBase}}", "{{.Name}}", {{if len .Parameters | lt 0}}params{{else}}nil{{end}}, &r.Options, &resp)
		if err != nil {
			return
		}
		apicalls := r.Options.GetRemainingAPICalls()
		var wg sync.WaitGroup
		for x := 1; x <= apicalls; x++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				offset := i * limit
				this_resp := []{{convertType .Type "services"}}{}
				options := r.Options
				options.Offset = &offset
				err = r.Session.DoRequest("{{$rawBase}}", "{{.Name}}", {{if len .Parameters | lt 0}}params{{else}}nil{{end}}, &options, &this_resp)
				resp = append(resp, this_resp...)
			}(x)
		}
		wg.Wait()
	return
	}
	{{end }}
	{{end }}

{{end}}
`, license, codegenWarning)

var testTemplate = `// AUTO GENERATED by tools/loadmeta.go
package services_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session/sessionfakes"
)

var _ = Describe("{{(index . 0 ).ServiceGroup}} Tests", func() {
	var slsession *sessionfakes.FakeSLSession
	BeforeEach(func() {
		slsession = &sessionfakes.FakeSLSession{}
	})
{{range .}}
	Context("Testing {{.Name}} service", func() {
		var sl_service services.{{.Name | removePrefix}}
		BeforeEach(func() {
			sl_service = services.Get{{.Name | removePrefix | desnake}}Service(slsession)
		})
{{- $serviceName := .Name }}
		Context("{{$serviceName}} Set Options", func() {
			It("Set Options properly", func() {
				t_id := 1234
				t_filter := "{'testFilter':{'test'}}"
				t_limit := 100
				t_offset := 5
				sl_service = sl_service.Id(t_id).Filter(t_filter).Offset(t_offset).Limit(t_limit)
				Expect(sl_service.Options.Id).To(HaveValue(Equal(t_id)))
				Expect(sl_service.Options.Filter).To(HaveValue(Equal(t_filter)))
				Expect(sl_service.Options.Limit).To(HaveValue(Equal(t_limit)))
				Expect(sl_service.Options.Offset).To(HaveValue(Equal(t_offset)))
			})
		})
		Context("{{$serviceName}} Set Mask", func() {
			It("Set Options properly", func() {
				t_mask1 := "mask[test,test2]"
				sl_service = sl_service.Mask(t_mask1)
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
				// Mask("test,test2") should set the mask to be "mask[test,test2]" aka t_mask1
				sl_service = sl_service.Mask("test,test2")
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
			})
		})
{{- range .Methods }}
		Context("{{$serviceName}}::{{.Name}}", func() {
			It("API Call Test", func() {
{{- if .Type|ne "void"}}
				_, err := sl_service.{{.Name | titleCase}}({{.Parameters | nilParam}})
{{- else }}
				err := sl_service.{{.Name | titleCase}}({{.Parameters | nilParam}})
{{- end }}
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
{{- end }}
	})
{{ end }}
})
`
