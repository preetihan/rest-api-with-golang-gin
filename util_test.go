package main

import (
	"reflect"
	"testing"
)

func TestConvertIntToString(t *testing.T) {
	expectedResult := "1"

	result := convertIntToString(1)

	if reflect.DeepEqual(expectedResult, result) {
		t.Log("TestConvertIntToString PASSED")
	} else {
		t.Errorf("TestConvertIntToString FAILED, expected %v but got %v", expectedResult, result)
	}
}
