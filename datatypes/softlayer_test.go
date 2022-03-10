package datatypes

import (
	"fmt"
	"testing"
)

func TestSearchUnmarshal(t *testing.T) {
	jsonString := []byte(`{
		"resourceType":"SoftLayer_Virtual_Guest",
		"resource": {
			"hostname": "test",
			"domain": "atest.com",
			"id": 1234,
			"accountId": 111111
		},
		"relevanceScore": "12.3"

	}`)
	fmt.Printf("TESTING\n")
	var result Container_Search_Result
	err := result.UnmarshalJSON(jsonString)
	if err != nil {
		t.Errorf("There was an error! %v\n", err)
	}
	if *result.ResourceType != "SoftLayer_Virtual_Guest" {
		t.Errorf("Failed to get resource type\n%v != SoftLayer_Virtual_Guest\n", result.ResourceType)
	}
	fmt.Printf("%+v\n", result.Resource)

	test := result.Resource
	if *test.(*Virtual_Guest).Hostname != "test" {
		t.Errorf("Failed to get a resource!\n%v != 'test'\n", result.Resource)
	}

}
