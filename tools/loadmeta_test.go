package main

import (
	"testing"
)


func TestRemovePrefix(t *testing.T) {
	result := RemovePrefix("SoftLayer_Test")
	if result != "Test" {
		t.Errorf("removePrefix(SoftLayer_Test) resulted in %v\n", result)
	}

	result = RemovePrefix("Something_Else")
	if result != "Something_Else" {
		t.Errorf("removePrefix(Something_Else) resulted in %v\n", result)
	}
}


func TestConvertType(t *testing.T) {
	typeMap := map[string]string{
		"unsignedLong": "uint",
		"boolean": "bool",
		"dateTime": "Time",
		"decimal": "Float64",
		"base64Binary": "[]byte",
		"json": "string",
		"SoftLayer_Hardware": "Hardware",
		"SoftLayer_Entity": "Entity",
		"McAfee_Something": "McAfee_Something",
		"test": "test",
	}
	var result string
	for input, expected := range typeMap {
		result = ConvertType(input, "datatypes")
		if result != expected {
			t.Errorf("ConvertType(%v) != %v, actually %v", input, expected, result)
		}
	}

	result = ConvertType("SoftLayer_Entity", "datatypes", "SoftLayer_Test", "resource")
	if result != "interface{}" {
		t.Errorf("ConvertType(SoftLayer_Entity, resource) != %v, actually %v", "interface{}", result)
	}

}

func TestRemoveReservedWords(t *testing.T) {
	result := RemoveReservedWords("type")
	if result != "typ" {
		t.Errorf("RemoveReservedWords(type) != typ")
	}
	result = RemoveReservedWords("other")
	if result != "other" {
		t.Errorf("RemoveReservedWords(other) != other")
	}
}

func TestDesnak(t *testing.T) {
	result := Desnake("SoftLayer_Test_This_Thing")
	if result != "SoftLayerTestThisThing" {
		t.Errorf("Desnake(SoftLayer_Test_This_Thing) != SoftLayerTestThisThing")
	}
}