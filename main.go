package spec2test

import "reflect"

// Server            *httptest.Server
// URLPath, HTTPVerb string
// ExpectedRespObj   interface{} // unmarshalled response data
// Type              ResourceType
// Header            http.Header
// NewDecoder        NewDecoderFn

// MakeExpectedRespObj returns an instance whose type is
// the same as that of `v`
func makeObj(typ reflect.Type, resourceType ResourceType) interface{} {
	return nil
}
